package logger

import (
	"go.uber.org/zap"
)

type ILogger interface {
	Init()
	Stop()
	// C logger client
	C() *zap.Logger
	Debug(scope string, fields ...zap.Field)
	Info(scope string, fields ...zap.Field)
	Warn(scope string, fields ...zap.Field)
	Error(scope string, fields ...zap.Field)
}
