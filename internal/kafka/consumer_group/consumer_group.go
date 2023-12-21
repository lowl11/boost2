package consumer_group

import (
	"github.com/IBM/sarama"
	"github.com/lowl11/boost2/data/types"
	"github.com/lowl11/boost2/internal/infrastructure/stopper"
	"github.com/lowl11/boost2/pkg/kafka/configurator"
	"time"
)

type ConsumerGroup struct {
	config            *configurator.Configurator
	topicName         string
	groupName         string
	client            sarama.ConsumerGroup
	innerClient       sarama.Client
	stoppers          []chan bool
	errorHandler      types.ErrorHandler
	connectionTimeout time.Duration
}

func New(topicName string, config *configurator.Configurator) *ConsumerGroup {
	consumerGroup := &ConsumerGroup{
		config:            config,
		topicName:         topicName,
		groupName:         config.Group(),
		connectionTimeout: config.GetConnectionTimeout(),
	}
	stopper.Get().Add(consumerGroup.Stop)
	return consumerGroup
}
