package server

import (
	"context"
	"io"
	moviesv1 "movie_buf/api/movies/v1"
	"movie_buf/internal/conf"
	"net"
	"net/http"
	"time"

	healthv1 "google.golang.org/grpc/health/grpc_health_v1"

	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/trace"
	"go.uber.org/multierr"
	"go.uber.org/zap"
	"google.golang.org/grpc/health"
	"gorm.io/gorm"
)

type Services struct {
	MoviesWritter moviesv1.MoviesWriterServiceServer
	MoviesReader  moviesv1.MoviesReaderServiceServer
	Health        *health.Server
}

type shutDowner interface {
	Shutdown(ctx context.Context) error
}

type grpcServer interface {
	Serve(net.Listener) error
}

type Application struct {
	server         grpcServer
	listener       net.Listener
	logger         *zap.Logger
	db             *gorm.DB
	services       Services
	tracerProvider trace.TracerProvider
	meterProvider  metric.MeterProvider
	shutdown       []shutDowner
	closer         []io.Closer
	cfg            conf.ServerConfig
	metricsServer  *http.Server
}

func (app Application) Run(ctx context.Context) error {
	go app.checkHealth(ctx)
	go app.serveMetrics()

	app.logger.Info("Running application")
	return app.server.Serve(app.listener)
}

// Shutdown releases any held resources by dependencies of this Application.
func (app Application) Shutdown(ctx context.Context) error {
	var err error
	for _, downer := range app.shutdown {
		if downer == nil {
			continue
		}
		err = multierr.Append(err, downer.Shutdown(ctx))
	}
	for _, closer := range app.closer {
		if closer == nil {
			continue
		}
		err = multierr.Append(err, closer.Close())
	}
	return err
}

func (app Application) checkHealth(ctx context.Context) {
	app.logger.Info("Running health service")
	for {
		if ctx.Err() != nil {
			return
		}
		app.services.Health.SetServingStatus("app.db", app.checkDatabaseHealth())
		time.Sleep(10 * time.Second)
	}
}

func (app Application) checkDatabaseHealth() healthv1.HealthCheckResponse_ServingStatus {
	state := healthv1.HealthCheckResponse_SERVING
	db, err := app.db.DB()
	if err != nil {
		state = healthv1.HealthCheckResponse_NOT_SERVING
	}
	if err = db.Ping(); err != nil {
		state = healthv1.HealthCheckResponse_NOT_SERVING
	}
	return state
}

func (app Application) serveMetrics() {
	if err := app.metricsServer.ListenAndServe(); err != nil {
		app.logger.Error("failed to listen and server to metrics server", zap.Error(err))
	}
}
