package async_producer

import (
	"github.com/IBM/sarama"
	"github.com/lowl11/boost2/pkg/kafka/configurator"
)

type Producer struct {
	client sarama.AsyncProducer
}

func New(config *configurator.Configurator) (*Producer, error) {
	client, err := sarama.NewAsyncProducer(config.Hosts(), config.Config())
	if err != nil {
		return nil, err
	}

	return &Producer{
		client: client,
	}, nil
}
