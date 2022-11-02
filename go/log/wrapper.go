package log

import (
	"go.uber.org/zap"
)

const STACKTRACE = "stacktrace"

type StackTraceLogger struct {
	suggaredLogger *zap.SugaredLogger
}

func NewStackTraceLogger(suggaredLogger *zap.SugaredLogger) *StackTraceLogger {
	return &StackTraceLogger{suggaredLogger: suggaredLogger}
}

// Debug uses fmt.Sprint to construct and log a message.
func (s *StackTraceLogger) Debug(args ...interface{}) {
	s.suggaredLogger.WithOptions(zap.AddCallerSkip(1)).Debug(args...)
}

// Info uses fmt.Sprint to construct and log a message.
func (s *StackTraceLogger) Info(args ...interface{}) {
	s.suggaredLogger.WithOptions(zap.AddCallerSkip(1)).Info(args...)
}

// Warn uses fmt.Sprint to construct and log a message.
func (s *StackTraceLogger) Warn(args ...interface{}) {
	s.suggaredLogger.WithOptions(zap.AddCallerSkip(1)).Warn(args...)
}

// Error uses fmt.Sprint to construct and log a message.
func (s *StackTraceLogger) Error(args ...interface{}) {
	if useSentry() {
		args = append(args, zap.StackSkip(STACKTRACE, 1))
	}
	s.suggaredLogger.WithOptions(zap.AddCallerSkip(1)).Error(args...)
}

// DPanic uses fmt.Sprint to construct and log a message. In development, the
// logger then panics. (See DPanicLevel for details.)
func (s *StackTraceLogger) DPanic(args ...interface{}) {
	if useSentry() {
		args = append(args, zap.StackSkip(STACKTRACE, 1))
	}
	s.suggaredLogger.WithOptions(zap.AddCallerSkip(1)).DPanic(args...)
}

// Panic uses fmt.Sprint to construct and log a message, then panics.
func (s *StackTraceLogger) Panic(args ...interface{}) {
	if useSentry() {
		args = append(args, zap.StackSkip(STACKTRACE, 1))
	}
	s.suggaredLogger.WithOptions(zap.AddCallerSkip(1)).Panic(args...)
}

// Fatal uses fmt.Sprint to construct and log a message, then calls os.Exit.
func (s *StackTraceLogger) Fatal(args ...interface{}) {
	if useSentry() {
		args = append(args, zap.StackSkip(STACKTRACE, 1))
	}
	s.suggaredLogger.WithOptions(zap.AddCallerSkip(1)).Fatal(args...)
}

// Debugf uses fmt.Sprintf to log a templated message.
func (s *StackTraceLogger) Debugf(template string, args ...interface{}) {
	s.suggaredLogger.WithOptions(zap.AddCallerSkip(1)).Debugf(template, args...)
}

// Infof uses fmt.Sprintf to log a templated message.
func (s *StackTraceLogger) Infof(template string, args ...interface{}) {
	s.suggaredLogger.WithOptions(zap.AddCallerSkip(1)).Infof(template, args...)
}

// Warnf uses fmt.Sprintf to log a templated message.
func (s *StackTraceLogger) Warnf(template string, args ...interface{}) {
	s.suggaredLogger.WithOptions(zap.AddCallerSkip(1)).Warnf(template, args...)
}

// Errorf uses fmt.Sprintf to log a templated message.
func (s *StackTraceLogger) Errorf(template string, args ...interface{}) {
	if useSentry() {
		args = append(args, zap.StackSkip(STACKTRACE, 1))
	}
	s.suggaredLogger.WithOptions(zap.AddCallerSkip(1)).Errorf(template, args...)
}

// DPanicf uses fmt.Sprintf to log a templated message. In development, the
// logger then panics. (See DPanicLevel for details.)
func (s *StackTraceLogger) DPanicf(template string, args ...interface{}) {
	if useSentry() {
		args = append(args, zap.StackSkip(STACKTRACE, 1))
	}
	s.suggaredLogger.WithOptions(zap.AddCallerSkip(1)).DPanicf(template, args...)
}

// Panicf uses fmt.Sprintf to log a templated message, then panics.
func (s *StackTraceLogger) Panicf(template string, args ...interface{}) {
	if useSentry() {
		args = append(args, zap.StackSkip(STACKTRACE, 1))
	}
	s.suggaredLogger.WithOptions(zap.AddCallerSkip(1)).Panicf(template, args...)
}

// Fatalf uses fmt.Sprintf to log a templated message, then calls os.Exit.
func (s *StackTraceLogger) Fatalf(template string, args ...interface{}) {
	if useSentry() {
		args = append(args, zap.StackSkip(STACKTRACE, 1))
	}
	s.suggaredLogger.WithOptions(zap.AddCallerSkip(1)).Fatalf(template, args...)
}
