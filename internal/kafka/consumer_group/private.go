package consumer_group

import (
	"context"
	"github.com/lowl11/boost2/data/types"
	"github.com/lowl11/boost2/internal/kafka/consumer_group/handler"
	"github.com/lowl11/boost2/log"
)

func (consumerGroup *ConsumerGroup) handleConsumers(handlerFunc types.KafkaConsumerHandler) error {
	h := handler.New(handlerFunc, consumerGroup.errorHandler, consumerGroup.stopper)

	for {
		if err := startConsume(consumerGroup, h); err != nil {
			return err
		}
	}
}

func startConsume(group *ConsumerGroup, h *handler.Handler) error {
	ctx, cancel := context.WithTimeout(context.Background(), group.connectionTimeout)
	defer cancel()

	select {
	case err := <-group.client.Errors():
		log.Error("Consumer group catch error: ", err)
		return err
	default:
		if err := group.client.Consume(ctx, []string{group.topicName}, h); err != nil {
			return err
		}
	}

	return nil
}
