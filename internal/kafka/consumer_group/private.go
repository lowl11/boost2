package consumer_group

import (
	"context"
	"github.com/lowl11/boost2/data/types"
	"github.com/lowl11/boost2/internal/kafka/consumer_group/handler"
)

func (consumerGroup *ConsumerGroup) handleConsumers(handlerFunc types.KafkaConsumerHandler) error {
	err := consumerGroup.client.Consume(
		context.Background(),
		consumerGroup.topicNames,
		handler.New(handlerFunc, consumerGroup.errorHandler, consumerGroup.stoppers[0]),
	)
	if err != nil {
		return err
	}

	return nil
}
