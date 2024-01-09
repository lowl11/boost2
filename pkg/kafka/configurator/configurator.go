package configurator

import (
	"github.com/IBM/sarama"
)

type Configurator struct {
	config *sarama.Config
	hosts  []string
}

func New(hosts ...string) *Configurator {
	return &Configurator{
		config: sarama.NewConfig(),
		hosts:  hosts,
	}
}
