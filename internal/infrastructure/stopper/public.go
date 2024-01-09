package stopper

import (
	"os"
)

func (service *Service) Add(stoppers ...func()) *Service {
	service.stoppers = append(service.stoppers, stoppers...)
	return service
}

func (service *Service) GetSignals() chan os.Signal {
	return service.signals
}

func (service *Service) Run() {
	for _, stopper := range service.stoppers {
		stopper()
	}
}
