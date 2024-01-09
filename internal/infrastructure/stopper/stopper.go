package stopper

import (
	"os"
	"os/signal"
)

type Service struct {
	stoppers []func()
	signals  chan os.Signal
}

var instance *Service

func Get() *Service {
	if instance != nil {
		return instance
	}

	signals := make(chan os.Signal)
	signal.Notify(signals, os.Interrupt, os.Kill)
	instance = &Service{
		stoppers: make([]func(), 0),
		signals:  signals,
	}
	return instance
}
