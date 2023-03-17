package log

import (
	"fmt"
	"github.com/TheZeroSlave/zapsentry"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func useSentry() bool {
	return os.Getenv("SENTRY_DSN") != ""
}

func addSentryToLogger(log *zap.Logger) *zap.Logger {
	if !useSentry() {
		return log
	}

	cfg := zapsentry.Configuration{
		Level:             zapcore.ErrorLevel, // when to send message to sentry
		DisableStacktrace: false,              // disable stacktrace
		Tags: map[string]string{
			"component": "system",
		},
	}
	core, err := zapsentry.NewCore(cfg, zapsentry.NewSentryClientFromDSN(os.Getenv("SENTRY_DSN")))

	// in case of err it will return noop core. so we can safely attach it
	if err != nil {
		log.Warn("failed to init zap", zap.Error(err))
	}
	return zapsentry.AttachCoreToLogger(core, log)
}

var Sugar *StackTraceLogger
var logger *zap.Logger

func init() {
	if logger == nil {
		logger, err := zap.NewProduction()
		if os.Getenv("DEBUG") == "true" {
			logger, err = zap.NewDevelopment()
			logger.Sugar().Debugf("debug mode enabled")
		}
		if err != nil {
			fmt.Println(err)
		}
		logger = addSentryToLogger(logger)
		Sugar = NewStackTraceLogger(logger.Sugar())
	}
}

func SyncLogger() {
	if logger != nil {
		err := logger.Sync()
		if err != nil {
			fmt.Println(err)
		}
	}
}
