package telemetry

import (
	"github.com/gojaguar/jaguar/config"
	"go.uber.org/zap"
)

func SetupLogger(cfg config.Config) (*zap.Logger, error) {
	var logger *zap.Logger

	var err error

	switch cfg.Environment {
	case "production":
		logger, err = zap.NewProduction()
	case "stagging":
		logger, err = zap.NewDevelopment()
	default:
		logger = zap.NewNop()
	}

	if err != nil {
		return nil, err
	}

	logger = logger.Named(cfg.Name)

	return logger, nil
}
