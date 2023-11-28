package async_producer

import "github.com/lowl11/boost2/pkg/kafka/configurator"

type Producer struct {
	//
}

func New(config *configurator.Configurator) (*Producer, error) {
	//sarama.NewSyncProducer()
	return &Producer{}, nil
}
