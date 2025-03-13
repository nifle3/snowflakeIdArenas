package logger

import (
	"fmt"
	"log/slog"
	"os"
	"strings"
)

const (
	dev  = "dev"
	prod = "prod"
)

func MustSetup(envType string) {
	var newLogger *slog.Logger
	envType = strings.ToLower(envType)

	switch envType {
	case dev:
		newLogger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		}))
	case prod:
		newLogger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelError,
			AddSource: true,
		}))
	default:
		panic(fmt.Sprintf("Invalid env type must be 'dev' 'prod' give %s\n", envType))
	}

	slog.SetDefault(newLogger)
}
