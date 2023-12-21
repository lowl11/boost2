package consumer_group

import (
	"context"
	"github.com/lowl11/boost2/data/types"
	"github.com/lowl11/boost2/internal/kafka/consumer_group/handler"
	"github.com/lowl11/boost2/log"
)

func (consumerGroup *ConsumerGroup) handleConsumers(ctx context.Context, handlerFunc types.KafkaConsumerHandler) error {
	h := handler.New(handlerFunc, consumerGroup.errorHandler, consumerGroup.stopper)

	for {
		select {
		case err := <-consumerGroup.client.Errors():
			log.Error("Consumer group catch error: ", err)
			return err
		case <-ctx.Done():
			return nil
		default:
			if err := consumerGroup.client.Consume(ctx, []string{consumerGroup.topicName}, h); err != nil {
				return err
			}
		}
	}
}
