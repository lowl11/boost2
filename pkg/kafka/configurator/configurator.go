package configurator

import "github.com/IBM/sarama"

type Configurator struct {
	config *sarama.Config

	hosts     []string
	groupName string
}

func New(hosts ...string) *Configurator {
	service := &Configurator{
		config: sarama.NewConfig(),
		hosts:  hosts,
	}
	service.loadBasic()
	return service
}
