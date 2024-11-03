package logger

import (
	"go.uber.org/zap"
)

type Logger interface {
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
	Sync() error
}

type zapLogger struct {
	log *zap.Logger
}

func New() Logger {
	logger, _ := zap.NewProduction()
	return &zapLogger{log: logger}
}

func (l *zapLogger) Info(msg string, fields ...zap.Field) {
	l.log.Info(msg, fields...)
}

func (l *zapLogger) Warn(msg string, fields ...zap.Field) {
	l.log.Warn(msg, fields...)
}

func (l *zapLogger) Error(msg string, fields ...zap.Field) {
	l.log.Error(msg, fields...)
}

func (l *zapLogger) Fatal(msg string, fields ...zap.Field) {
	l.log.Fatal(msg, fields...)
}

func (l *zapLogger) Sync() error {
	return l.log.Sync()
}
