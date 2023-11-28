package consumer_group

import (
	"github.com/lowl11/boost2/data/interfaces"
	"github.com/lowl11/boost2/data/types"
)

func (consumerGroup *ConsumerGroup) SetErrorHandler(errorHandler types.ErrorHandler) interfaces.IConsumer {
	//consumer.errorHandler = errorHandler
	return consumerGroup
}

func (consumerGroup *ConsumerGroup) StartConsume(handlerFunc types.KafkaConsumerHandler) error {
	return nil
}

func (consumerGroup *ConsumerGroup) StartConsumeAsync(handlerFunc types.KafkaConsumerHandler) {
	//
}
