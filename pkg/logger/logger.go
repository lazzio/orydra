package logger

import (
	"log/slog"
	"os"
	"runtime"
)

var (
	// Logger is the global logger
	Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
)

// GetFunctionName returns the name of the function that called the logger
// @return string
func GetFunctionName() string {
	pc, _, _, ok := runtime.Caller(2)
	if !ok {
		return "unknown"
	}
	return runtime.FuncForPC(pc).Name()
}
