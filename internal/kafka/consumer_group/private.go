package consumer_group

import (
	"context"
	"github.com/lowl11/boost2/data/types"
	"github.com/lowl11/boost2/internal/kafka/consumer_group/handler"
)

func (consumerGroup *ConsumerGroup) handleConsumers(handlerFunc types.KafkaConsumerHandler) error {
	if consumerGroup.stoppers == nil {
		consumerGroup.stoppers = make([]chan bool, 0, 1)
	}

	consumerGroup.stoppers = append(consumerGroup.stoppers, make(chan bool, 1))

	for {
		select {
		case err := <-consumerGroup.client.Errors():
			return err
		default:
			if err := consumerGroup.client.Consume(
				context.Background(),
				[]string{consumerGroup.topicName},
				handler.New(handlerFunc, consumerGroup.errorHandler, consumerGroup.stoppers[0]),
			); err != nil {
				return err
			}
		}
	}
}
