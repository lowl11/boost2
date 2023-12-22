package consumer_group

import (
	"context"
	"github.com/lowl11/boost2/data/types"
	"github.com/lowl11/boost2/internal/kafka/consumer_group/handler"
	"github.com/lowl11/boost2/log"
)

func (consumerGroup *ConsumerGroup) handleConsumers(ctx context.Context, handlerFunc types.KafkaConsumerHandler) error {
	h := handler.New(handlerFunc, consumerGroup.errorHandler, consumerGroup.stopper)

	go func() {
		if err := consumerGroup.client.Consume(ctx, []string{consumerGroup.topicName}, h); err != nil {
			log.Fatal("Start consuming error: ", err)
		}
	}()

	for {
		select {
		case err := <-consumerGroup.client.Errors():
			if err != nil {
				return err
			}
		case <-ctx.Done():
			return nil
		}
	}
}
