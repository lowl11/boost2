package stopper

import (
	"os"
	"os/signal"
)

func (service *Service) Add(stoppers ...func()) *Service {
	service.stoppers = append(service.stoppers, stoppers...)
	return service
}

func (service *Service) Catch() {
	go func() {
		signals := make(chan os.Signal, 1)
		signal.Notify(signals, os.Interrupt)
		<-signals

		for _, stopper := range service.stoppers {
			stopper()
		}
	}()
}
