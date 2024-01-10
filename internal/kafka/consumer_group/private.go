package consumer_group

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/lowl11/boost2/log"
)

func (consumerGroup *ConsumerGroup) handleConsumers(ctx context.Context, handler sarama.ConsumerGroupHandler) error {
	go func() {
		if err := consumerGroup.client.Consume(ctx, []string{consumerGroup.topicName}, handler); err != nil {
			log.Fatal("Start consuming error: ", err)
		}
	}()

	for {
		select {
		case err := <-consumerGroup.client.Errors():
			if err != nil {
				return err
			}
		case <-consumerGroup.stopper:
			return nil
		case <-ctx.Done():
			return nil
		}
	}
}
