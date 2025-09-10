package gateway

import "go.uber.org/zap"

type LoggerGateway interface {
	Info(msg string, fields ...zap.Field)
	Debug(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
}
