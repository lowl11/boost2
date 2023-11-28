package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	sugar *zap.SugaredLogger
}

var instance *Logger

func Get() *Logger {
	if instance != nil {
		return instance
	}

	atom := zap.NewAtomicLevelAt(zapcore.DebugLevel)
	config := zap.NewProductionConfig()
	config.Level = atom
	zapLogger, _ := config.Build()
	// TODO: sync logger

	logger := &Logger{
		sugar: zapLogger.Sugar(),
	}
	return logger
}
