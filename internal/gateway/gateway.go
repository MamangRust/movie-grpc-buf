package gateway

import (
	"movie_buf/internal/conf"
	"movie_buf/internal/telemetry"
	"net/http"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
)

type Gateway struct {
	Telemeter telemetry.Telemetry
	mux       *runtime.ServeMux
	handler   http.Handler
	Config    conf.GatewayConfig
}
