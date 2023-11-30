package stopper

import (
	"os"
	"os/signal"
)

func (service *Service) Add(stoppers ...func()) *Service {
	service.stoppers = append(service.stoppers, stoppers...)
	return service
}

func (service *Service) GetSignals() chan os.Signal {
	return service.signals
}

func (service *Service) Catch() {
	go func() {
		signal.Notify(service.signals, os.Interrupt)
		<-service.signals

		for _, stopper := range service.stoppers {
			stopper()
		}
	}()
}
