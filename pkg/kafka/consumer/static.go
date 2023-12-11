package consumer

import (
	"github.com/lowl11/boost2"
	"github.com/lowl11/boost2/internal/kafka/consumer"
	"github.com/lowl11/boost2/internal/kafka/consumer_group"
	"github.com/lowl11/boost2/log"
	"github.com/lowl11/boost2/pkg/kafka/configurator"
)

func New(topicName string, config *configurator.Configurator) boost2.IConsumer {
	return consumer.New(topicName, config)
}

func NewGroup(config *configurator.Configurator, topicNames ...string) boost2.IConsumer {
	if len(topicNames) == 0 {
		log.Fatal("No given topic names for Kafka Consumer Group")
	}

	return consumer_group.New(config, topicNames...)
}
