package logger

import (
	"log/slog"
	"os"
)

var Logger *slog.Logger

func Init() {
	Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
}

func GetLogger() *slog.Logger {
	if Logger == nil {
		Init()
	}

	return Logger
}
