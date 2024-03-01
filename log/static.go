package log

import (
	"github.com/lowl11/boost2/internal/infrastructure/logger"
	"github.com/lowl11/boost2/pkg/kafka/configurator"
	"go.uber.org/zap/zapcore"
)

var (
	DebugLevel   = zapcore.DebugLevel
	InfoLevel    = zapcore.InfoLevel
	WarmLevel    = zapcore.WarnLevel
	ErrorLevel   = zapcore.ErrorLevel
	DPanicLevel  = zapcore.DPanicLevel
	PanicLevel   = zapcore.PanicLevel
	FatalLevel   = zapcore.FatalLevel
	InvalidLevel = zapcore.InvalidLevel
)

type Config struct {
	SendToKafka bool

	Cfg      *configurator.Configurator
	LogTopic string
	LogLevel zapcore.Level
}

func InitLogger(cfg *Config) error {
	return logger.Configured(logger.Config{
		SendToKafka: cfg.SendToKafka,
		Cfg:         cfg.Cfg,
		LogTopic:    cfg.LogTopic,
		LogLevel:    cfg.LogLevel,
	})
}

func Debug(args ...any) {
	logger.Get().Debug(args...)
}

func Debugf(template string, args ...any) {
	logger.Get().Debugf(template, args...)
}

func Info(args ...any) {
	logger.Get().Info(args...)
}

func Infof(template string, args ...any) {
	logger.Get().Infof(template, args...)
}

func Warn(args ...any) {
	logger.Get().Warn(args...)
}

func Warnf(template string, args ...any) {
	logger.Get().Warnf(template, args...)
}

func Error(args ...any) {
	logger.Get().Error(args...)
}

func Errorf(template string, args ...any) {
	logger.Get().Errorf(template, args...)
}

func Fatal(args ...any) {
	logger.Get().Fatal(args...)
}

func Fatalf(template string, args ...any) {
	logger.Get().Fatalf(template, args...)
}
