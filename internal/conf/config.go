package conf

import (
	"fmt"

	"github.com/caarlos0/env/v9"
	"github.com/gojaguar/jaguar/config"
)

type Metrics struct {
	Enabled bool   `env:"ENABLED" envDefault:"true"`
	Host    string `env:"HOST" envDefault:"localhost"`
	Port    uint16 `env:"PORT" envDefault:"4317"`
}

func (m Metrics) Address() string {
	return fmt.Sprintf("%s:%d", m.Host, m.Port)
}

type Tracing struct {
	Enabled bool   `env:"ENABLED" envDefault:"true"`
	Host    string `env:"HOST" envDefault:"localhost"`
	Port    uint16 `env:"PORT" envDefault:"4317"`
}

func (t Tracing) Address() string {
	return fmt.Sprintf("%s:%d", t.Host, t.Port)
}

type ServerConfig struct {
	config.Config
	DB      config.Database `envPrefix:"DATABASE_"`
	Metrics Metrics         `envPrefix:"METRICS_"`
	Tracing Tracing         `envPrefix:"TRACING_"`
}

func (c ServerConfig) Listener() (network string, address string) {
	return "tcp", fmt.Sprintf(":%d", c.Port)
}

func ReadServerConfig() (ServerConfig, error) {
	var cfg ServerConfig
	err := env.Parse(&cfg)
	if err != nil {
		return ServerConfig{}, err
	}
	return cfg, nil
}

type GatewayConfig struct {
	config.Config
	ServerAddress string  `env:"SERVER_ADDRESS"`
	Metrics       Metrics `envPrefix:"METRICS_"`
	Tracing       Tracing `envPrefix:"TRACING_"`
}

func ReadGatewayConfig() (GatewayConfig, error) {
	var cfg GatewayConfig
	err := env.Parse(&cfg)
	if err != nil {
		return GatewayConfig{}, err
	}
	return cfg, nil
}
