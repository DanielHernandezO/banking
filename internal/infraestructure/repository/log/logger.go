package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type loggerCollector struct {
	logger *zap.Logger
}

func NewLoggerCollector() *loggerCollector {
	return &loggerCollector{
		logger: buildLogger(),
	}
}

func buildLogger() *zap.Logger {
	config := zap.NewProductionConfig()
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig = encoderConfig
	logger, err := config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	return logger
}

func (l *loggerCollector) Info(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

func (l *loggerCollector) Debug(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}

func (l *loggerCollector) Error(msg string, fields ...zap.Field) {
	l.logger.Info(msg, fields...)
}
