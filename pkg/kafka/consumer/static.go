package consumer

import (
	"github.com/lowl11/boost2"
	"github.com/lowl11/boost2/internal/kafka/consumer"
	"github.com/lowl11/boost2/internal/kafka/consumer_group"
	"github.com/lowl11/boost2/pkg/kafka/configurator"
)

func New(topicName string, config *configurator.Configurator) boost2.IConsumer {
	return consumer.New(topicName, config)
}

func NewGroup(topicName, groupName string, config *configurator.Configurator) boost2.IConsumer {
	return consumer_group.New(topicName, groupName, config)
}
