package interceptors

import (
	"movie_buf/internal/telemetry"
	"time"

	grpc_logging "github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/retry"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
)

func NewClientUnaryInterceptors(telemeter telemetry.Telemetry) grpc.DialOption {
	return grpc.WithChainUnaryInterceptor(
		grpc_logging.UnaryClientInterceptor(interceptorLogger(telemeter.Logger)),
		retry.UnaryClientInterceptor(
			retry.WithCodes(codes.ResourceExhausted, codes.Unavailable),
			retry.WithMax(10),
			retry.WithBackoff(retry.BackoffExponential(50*time.Millisecond)),
		),
	)
}

func NewClientStreamInterceptors(telemeter telemetry.Telemetry) grpc.DialOption {
	return grpc.WithChainStreamInterceptor(
		grpc_logging.StreamClientInterceptor(interceptorLogger(telemeter.Logger)),
		retry.StreamClientInterceptor(
			retry.WithCodes(codes.ResourceExhausted, codes.Unavailable),
			retry.WithMax(10),
			retry.WithBackoff(retry.BackoffExponential(50*time.Millisecond)),
		),
	)
}
