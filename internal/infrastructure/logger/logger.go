package logger

import (
	"os"

	"github.com/IBM/sarama"
	"github.com/lowl11/boost2/internal/infrastructure/stopper"
	"github.com/lowl11/boost2/pkg/kafka/configurator"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	sugar      *zap.SugaredLogger
	logChannel chan func()
}

var instance *Logger

type Config struct {
	SendToKafka bool

	Cfg      *configurator.Configurator
	LogTopic string
	LogLevel zapcore.Level

	ServiceName string
}

func Configured(cfg Config) error {
	// Create Zap encoder
	encoderCfg := zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.EpochTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	encoder := zapcore.NewJSONEncoder(encoderCfg)

	// Create Zap core
	core := zapcore.NewCore(encoder, zapcore.AddSync(os.Stdout), zapcore.DebugLevel)

	if cfg.SendToKafka {
		producer, err := sarama.NewSyncProducer(cfg.Cfg.Hosts(), cfg.Cfg.Config())
		if err != nil {
			return err
		}

		kafkaSink := NewKafkaSink(producer, cfg.LogTopic, cfg.ServiceName)
		// Send logs to Kafka by adding KafkaSink to the core
		core = zapcore.NewTee(core, zapcore.NewCore(encoder, zapcore.AddSync(kafkaSink), cfg.LogLevel))
	}

	logger := zap.New(core)

	zap.ReplaceGlobals(logger)
	stopper.Get().Add(func() {
		_ = logger.Sync()
	})

	instance = &Logger{
		sugar:      logger.Sugar(),
		logChannel: make(chan func()),
	}
	instance.listen()

	return nil
}

func Get() *Logger {
	if instance != nil {
		return instance
	}

	config := zap.NewProductionConfig()
	addEncoder(&config)
	addLevel(&config)
	zapLogger, _ := config.Build()
	stopper.Get().Add(func() {
		_ = zapLogger.Sync()
	})

	instance = &Logger{
		sugar:      zapLogger.Sugar(),
		logChannel: make(chan func()),
	}
	instance.listen()
	return instance
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
