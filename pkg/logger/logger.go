package logger

import (
	"log/slog"
	"os"
)

var (
	// Logger is the global logger
	Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
)
