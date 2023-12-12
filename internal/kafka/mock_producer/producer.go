package mock_producer

import "github.com/lowl11/boost2/pkg/kafka/configurator"

type Producer struct {
	config *configurator.Configurator
}

func New(config *configurator.Configurator) (*Producer, error) {
	return &Producer{
		config: config,
	}, nil
}
