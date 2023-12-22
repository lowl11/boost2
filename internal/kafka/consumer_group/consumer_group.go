package consumer_group

import (
	"github.com/IBM/sarama"
	"github.com/lowl11/boost2/data/types"
	"github.com/lowl11/boost2/internal/infrastructure/stopper"
	"github.com/lowl11/boost2/pkg/kafka/configurator"
)

type ConsumerGroup struct {
	config       *configurator.Configurator
	topicName    string
	groupName    string
	client       sarama.ConsumerGroup
	stopper      chan bool
	errorHandler types.ErrorHandler
	alwaysCommit bool
}

func New(topicName, groupName string, config *configurator.Configurator) *ConsumerGroup {
	consumerGroup := &ConsumerGroup{
		config:       config,
		topicName:    topicName,
		groupName:    groupName,
		stopper:      make(chan bool),
		alwaysCommit: config.IsAlwaysCommit(),
	}
	stopper.Get().Add(consumerGroup.Stop)
	return consumerGroup
}
