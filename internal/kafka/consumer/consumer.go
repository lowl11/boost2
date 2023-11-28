package consumer

import (
	"github.com/IBM/sarama"
	"github.com/lowl11/boost2/data/types"
	"github.com/lowl11/boost2/internal/infrastructure/stopper"
	"github.com/lowl11/boost2/pkg/kafka/configurator"
)

type Consumer struct {
	config       *configurator.Configurator
	topicName    string
	client       sarama.Consumer
	stoppers     []chan bool
	errorHandler types.ErrorHandler
}

func New(topicName string, config *configurator.Configurator) *Consumer {
	consumer := &Consumer{
		config:    config,
		topicName: topicName,
	}
	stopper.Get().Add(consumer.Stop)
	return consumer
}
