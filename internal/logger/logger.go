package logger

import (
	"context"
	"log/slog"
	"os"
)

type contextKey string

const RequestIDKey contextKey = "RequestID"

var Logger *slog.Logger

func Init(ctx context.Context) {
	baseHandler := slog.NewJSONHandler(os.Stdout, nil)
	Logger = slog.New(baseHandler).With(slog.String("service", "granite-identity"))
	Logger = Logger.With(slog.String("request_id", ctx.Value(RequestIDKey).(string)))
}

func GetLogger() *slog.Logger {
	return Logger
}
