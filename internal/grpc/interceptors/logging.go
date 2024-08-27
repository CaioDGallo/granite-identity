package interceptors

import (
	"context"

	"github.com/CaioDGallo/granite-identity/internal/logger"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func LoggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	var requestID string
	if ok {
		if values := md.Get("X-Request-ID"); len(values) > 0 {
			requestID = values[0]
		}
	}

	if requestID == "" {
		requestID = uuid.New().String()
	}

	ctx = context.WithValue(ctx, logger.RequestIDKey, requestID)

	logger.Init(ctx)

	resp, err := handler(ctx, req)

	return resp, err
}
