package logger

import (
	"log/slog"
	"time"

	"github.com/go-chi/httplog/v2"
	"github.com/spf13/viper"
)

var (
	HttpLogger = httplog.NewLogger(viper.GetString("APPNAME"), httplog.Options{
		LogLevel:         slog.LevelInfo,
		JSON:             true,
		Concise:          false,
		RequestHeaders:   false,
		ResponseHeaders:  false,
		MessageFieldName: "msg",
		LevelFieldName:   "level",
		TimeFieldName:    "time",
		TimeFieldFormat:  time.RFC3339,
		QuietDownPeriod:  10 * time.Second,
	})
)
