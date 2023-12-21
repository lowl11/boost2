package consumer_group

import (
	"context"
	"github.com/lowl11/boost2/data/types"
	"github.com/lowl11/boost2/internal/kafka/consumer_group/handler"
	"time"
)

func (consumerGroup *ConsumerGroup) handleConsumers(handlerFunc types.KafkaConsumerHandler) error {
	if consumerGroup.stoppers == nil {
		consumerGroup.stoppers = make([]chan bool, 0, 1)
	}

	consumerGroup.stoppers = append(consumerGroup.stoppers, make(chan bool, 1))

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	for {
		select {
		case err := <-consumerGroup.client.Errors():
			return err
		default:
			if err := consumerGroup.client.Consume(
				ctx,
				[]string{consumerGroup.topicName},
				handler.New(handlerFunc, consumerGroup.errorHandler, consumerGroup.stoppers[0]),
			); err != nil {
				return err
			}
		}
	}
}
