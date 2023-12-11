package logger

import (
	"github.com/lowl11/boost2/internal/infrastructure/stopper"
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

	config := zap.NewProductionConfig()
	addEncoder(&config)
	addLevel(&config)
	zapLogger, _ := config.Build()
	stopper.Get().Add(func() {
		if err := zapLogger.Sync(); err != nil {
			instance.Error("Sync logger error: ", err)
		}
	})

	logger := &Logger{
		sugar: zapLogger.Sugar(),
	}
	return logger
}

func addEncoder(config *zap.Config) {
	config.Encoding = "console"
	config.EncoderConfig = zapcore.EncoderConfig{
		MessageKey:  "message",
		LevelKey:    "level",
		TimeKey:     "time",
		EncodeLevel: zapcore.CapitalColorLevelEncoder,
		EncodeTime:  zapcore.RFC3339TimeEncoder,
	}
}

func addLevel(config *zap.Config) {
	atom := zap.NewAtomicLevelAt(zapcore.DebugLevel)
	config.Level = atom
}
