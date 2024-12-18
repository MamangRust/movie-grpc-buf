package server

import (
	"database/sql"
	"fmt"
	"io"
	moviesv1 "movie_buf/api/movies/v1"
	"movie_buf/internal/conf"
	"movie_buf/internal/domain"
	"movie_buf/internal/interceptors"
	"movie_buf/internal/service"
	"movie_buf/internal/telemetry"
	"net"
	"net/http"

	"github.com/bufbuild/protovalidate-go"
	"github.com/gojaguar/jaguar/database"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/zap"
	"google.golang.org/grpc/health"
	healthv1 "google.golang.org/grpc/health/grpc_health_v1"
	"gorm.io/gorm"

	"google.golang.org/grpc"
)

func Setup(cfg conf.ServerConfig) (Application, error) {

	telemeter, err := telemetry.SetupTelemetry(cfg.Config, cfg.Tracing, cfg.Metrics)
	if err != nil {
		return Application{}, err
	}

	telemeter.Logger.Debug("Initializing server", zap.String("server.name", cfg.Name), zap.String("server.environment", cfg.Environment))

	db, dbConn, err := setupDB(cfg, telemeter.Logger, telemeter.TracerProvider)
	if err != nil {
		return Application{}, err
	}

	err = domain.MigrateModels(db)
	if err != nil {
		return Application{}, err
	}

	l, err := setupListener(cfg, telemeter.Logger)
	if err != nil {
		return Application{}, err
	}

	validator, err := protovalidate.New()
	if err != nil {
		return Application{}, err
	}

	srv := grpc.NewServer(interceptors.NewServerInterceptors(telemeter, validator)...)
	svc := setupServices(db, telemeter.Logger, telemeter.TracerProvider, telemeter.MeterProvider)
	registerServices(srv, svc)

	metricsServer := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.Port+1),
		Handler: promhttp.Handler(),
	}

	return Application{
		server:         srv,
		listener:       l,
		logger:         telemeter.Logger,
		tracerProvider: telemeter.TracerProvider,
		meterProvider:  telemeter.MeterProvider,
		db:             db,
		services:       svc,
		metricsServer:  metricsServer,
		shutdown: []shutDowner{
			telemeter.TraceExporter,
			telemeter.MeterExporter,
		},
		closer: []io.Closer{
			dbConn,
			metricsServer,
		},
		cfg: cfg,
	}, nil
}

func registerServices(srv *grpc.Server, svc Services) {
	moviesv1.RegisterMoviesWriterServiceServer(srv, svc.MoviesWritter)
	moviesv1.RegisterMoviesReaderServiceServer(srv, svc.MoviesReader)
	healthv1.RegisterHealthServer(srv, svc.Health)
}

// setupServices initializes the Application Services.
func setupServices(db *gorm.DB, logger *zap.Logger, tracerProvider trace.TracerProvider, meterProvider metric.MeterProvider) Services {
	logger.Debug("Initializing services")
	moviesWriterService := service.NewMoviesWriter(db, logger, meterProvider.Meter("movies"))
	moviesReaderService := service.NewMoviesReader(db, logger, meterProvider.Meter("movies"))
	healthService := health.NewServer()
	return Services{
		MoviesWritter: moviesWriterService,
		MoviesReader:  moviesReaderService,
		Health:        healthService,
	}
}

// setupListener initializes a new tcp listener used by a gRPC server.
func setupListener(cfg conf.ServerConfig, logger *zap.Logger) (net.Listener, error) {
	protocol, address := cfg.Listener()
	logger.Debug("Initializing listener", zap.String("listener.protocol", protocol), zap.String("listener.address", address))
	l, err := net.Listen(protocol, address)
	if err != nil {
		logger.Error("Failed to initialize listener", zap.Error(err))
		return nil, err
	}
	return l, nil
}

// setupDB initializes a new connection with a DB server.
func setupDB(cfg conf.ServerConfig, logger *zap.Logger, provider trace.TracerProvider) (*gorm.DB, *sql.DB, error) {
	logger.Debug("Initializing DB connection", zap.String("db.engine", cfg.DB.Engine), zap.String("db.dsn", cfg.DB.DSN()))
	db, err := database.SetupConnectionSQL(cfg.DB)
	if err != nil {
		logger.Error("Failed to initialize DB connection", zap.Error(err))
		return nil, nil, err
	}
	err = db.Use(otelgorm.NewPlugin(
		otelgorm.WithDBName(cfg.DB.Name),
		otelgorm.WithAttributes(attribute.String("db.engine", cfg.DB.Engine)),
		otelgorm.WithTracerProvider(provider),
	))
	if err != nil {
		logger.Error("Failed to set up DB plugin", zap.Error(err))
		return nil, nil, err
	}
	dbConn, err := db.DB()
	if err != nil {
		logger.Error("Failed get db connection", zap.Error(err))
		return nil, nil, err
	}
	return db, dbConn, nil
}
