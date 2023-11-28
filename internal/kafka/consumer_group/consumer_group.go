package consumer_group

import "github.com/lowl11/boost2/pkg/kafka/configurator"

type ConsumerGroup struct {
	//
}

func New(topicName string, config *configurator.Configurator) *ConsumerGroup {
	return &ConsumerGroup{}
}
