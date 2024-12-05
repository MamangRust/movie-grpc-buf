package main

import (
	"context"
	"log"
	"movie_buf/internal/conf"
	"movie_buf/internal/gateway"
)

func main() {
	ctx := context.Background()

	cfg, err := conf.ReadGatewayConfig()
	if err != nil {
		log.Fatalln("Failed to read client config:", err)
	}

	gw, err := gateway.Setup(ctx, cfg)
	if err != nil {
		log.Fatalln("Failed to initialize gateway:", err)
	}
	if err := gateway.Run(gw); err != nil {
		log.Fatalln("Failed to run gateway:", err)
	}
}
