package consumer

import (
	"github.com/lowl11/boost2/internal/kafka/consumer"
	"github.com/lowl11/boost2/internal/kafka/consumer_group"
	"github.com/lowl11/boost2/pkg/kafka/configurator"
)

func New(topicName string, config *configurator.Configurator) *consumer.Consumer {
	return consumer.New(topicName, config)
}

func NewGroup(topicName, groupName string, config *configurator.Configurator) *consumer_group.ConsumerGroup {
	return consumer_group.New(topicName, groupName, config)
}
