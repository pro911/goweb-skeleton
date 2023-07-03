package logger

import (
	"go.uber.org/zap"
)

func ZapDebug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...)
}

func ZapInfo(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

func ZapWarn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

func ZapError(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

func ZapDPanic(msg string, fields ...zap.Field) {
	Logger.DPanic(msg, fields...)
}

func ZapPanic(msg string, fields ...zap.Field) {
	Logger.Panic(msg, fields...)
}

func ZapFatal(msg string, fields ...zap.Field) {
	Logger.Fatal(msg, fields...)
}
