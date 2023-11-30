package stopper

import "os"

type Service struct {
	stoppers []func()
	signals  chan os.Signal
}

var instance *Service

func Get() *Service {
	if instance != nil {
		return instance
	}

	instance = &Service{
		stoppers: make([]func(), 0),
		signals:  make(chan os.Signal, 1),
	}
	return instance
}
