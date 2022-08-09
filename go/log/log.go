package log

import (
	"fmt"
	"github.com/TheZeroSlave/zapsentry"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func addSentryToLogger(log *zap.Logger) *zap.Logger {
	sentryDsn := os.Getenv("SENTRY_DSN")
	if sentryDsn == "" {
		return log
	}

	cfg := zapsentry.Configuration{
		Level: zapcore.ErrorLevel, // when to send message to sentry
		Tags: map[string]string{
			"component": "system",
		},
	}
	core, err := zapsentry.NewCore(cfg, zapsentry.NewSentryClientFromDSN(sentryDsn))

	// to use breadcrumbs feature - create new scope explicitly
	log = log.With(zapsentry.NewScope())

	// in case of err it will return noop core. so we can safely attach it
	if err != nil {
		log.Warn("failed to init zap", zap.Error(err))
	}
	return zapsentry.AttachCoreToLogger(core, log)
}

var Sugar *zap.SugaredLogger
var logger *zap.Logger

func init() {
	if logger == nil {
		logger, err := zap.NewProduction()
		if os.Getenv("DEBUG") == "true" {
			logger, err = zap.NewDevelopment()
		}
		if err != nil {
			fmt.Println(err)
		}
		logger = addSentryToLogger(logger)
		Sugar = logger.Sugar()
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
