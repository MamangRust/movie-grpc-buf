package telemetry

import (
	"context"
	"fmt"
	"movie_buf/internal/conf"

	"github.com/gojaguar/jaguar/config"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/metric"
	"go.opentelemetry.io/otel/metric/noop"
	metric_sdk "go.opentelemetry.io/otel/sdk/metric"
)

func SetupMetrics(cfg config.Config, metrics conf.Metrics) (metric.MeterProvider, metric_sdk.Exporter, error) {
	if !metrics.Enabled {
		return noop.NewMeterProvider(), nil, nil
	}

	var meterProvider metric.MeterProvider
	var meterExporer metric_sdk.Exporter

	switch cfg.Environment {
	case "production", "stagging":
		var err error

		meterProvider, meterExporer, err = newMetrics(cfg, metrics)

		if err != nil {
			return nil, nil, err
		}
	default:
		meterProvider = noop.NewMeterProvider()
	}

	return meterProvider, meterExporer, nil
}

func newMetrics(cfg config.Config, metrics conf.Metrics) (metric.MeterProvider, metric_sdk.Exporter, error) {

	ctx := context.Background()

	res, err := newResource(ctx, cfg)

	if err != nil {
		return nil, nil, err
	}

	conn, err := newConn(ctx, metrics.Address())

	if err != nil {
		return nil, nil, err
	}

	meterExporter, err := otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithGRPCConn(conn))

	if err != nil {
		return nil, nil, fmt.Errorf("failed to create metrics exproter: %w", err)
	}

	meterProvider := metric_sdk.NewMeterProvider(
		metric_sdk.WithReader(metric_sdk.NewPeriodicReader(meterExporter)),
		metric_sdk.WithResource(res),
	)

	return meterProvider, meterExporter, nil
}
