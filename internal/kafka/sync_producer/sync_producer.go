package sync_producer

import (
	"github.com/IBM/sarama"
	"github.com/lowl11/boost2/pkg/kafka/configurator"
)

type Producer struct {
	client    sarama.SyncProducer
	isBatch   bool
	batchSize int
}

func New(config *configurator.Configurator) (*Producer, error) {
	client, err := sarama.NewSyncProducer(config.Hosts(), config.Config())
	if err != nil {
		return nil, err
	}

	return &Producer{
		client:    client,
		isBatch:   config.IsBatch(),
		batchSize: config.BatchSize(),
	}, nil
}
