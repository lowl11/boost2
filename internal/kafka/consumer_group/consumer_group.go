package consumer_group

import (
	"github.com/IBM/sarama"
	"github.com/lowl11/boost2/data/types"
	"github.com/lowl11/boost2/internal/infrastructure/stopper"
	"github.com/lowl11/boost2/pkg/kafka/configurator"
)

type ConsumerGroup struct {
	config       *configurator.Configurator
	topicNames   []string
	groupName    string
	client       sarama.ConsumerGroup
	innerClient  sarama.Client
	stoppers     []chan bool
	errorHandler types.ErrorHandler
}

func New(config *configurator.Configurator, topicNames ...string) *ConsumerGroup {
	consumerGroup := &ConsumerGroup{
		config:     config,
		topicNames: topicNames,
		groupName:  config.Group(),
	}
	stopper.Get().Add(consumerGroup.Stop)
	return consumerGroup
}
