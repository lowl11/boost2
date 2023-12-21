package consumer_group

import (
	"context"
	"github.com/lowl11/boost2/data/types"
	"github.com/lowl11/boost2/internal/kafka/consumer_group/handler"
	"github.com/lowl11/boost2/log"
)

func (consumerGroup *ConsumerGroup) handleConsumers(handlerFunc types.KafkaConsumerHandler) error {
	if consumerGroup.stoppers == nil {
		consumerGroup.stoppers = make([]chan bool, 0, 1)
	}

	consumerGroup.stoppers = append(consumerGroup.stoppers, make(chan bool, 1))

	for {
		if err := startConsume(consumerGroup, handlerFunc); err != nil {
			return err
		}
	}
}

func startConsume(group *ConsumerGroup, handlerFunc types.KafkaConsumerHandler) error {
	ctx, cancel := context.WithTimeout(context.Background(), group.connectionTimeout)
	defer cancel()

	select {
	case err := <-group.client.Errors():
		log.Error("Consumer group catch error: ", err)
		return nil
	default:
		if err := group.client.Consume(
			ctx,
			[]string{group.topicName},
			handler.New(handlerFunc, group.errorHandler, group.stoppers[0]),
		); err != nil {
			return err
		}
	}

	return nil
}
