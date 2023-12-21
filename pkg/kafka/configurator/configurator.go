package configurator

import (
	"github.com/IBM/sarama"
	"time"
)

type Configurator struct {
	config *sarama.Config

	hosts     []string
	groupName string
	isBatch   bool
	batchSize int

	connectionTimeout *time.Duration
}

func New(hosts ...string) *Configurator {
	service := &Configurator{
		config: sarama.NewConfig(),
		hosts:  hosts,
	}
	service.loadBasic()
	return service
}
