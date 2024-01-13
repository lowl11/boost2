package consumer_group

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/lowl11/boost2/data/types"
	"github.com/lowl11/boost2/internal/infrastructure/destroyer"
	"github.com/lowl11/boost2/log"
)

func (consumerGroup *ConsumerGroup) SetErrorHandler(errorHandler types.ErrorHandler) *ConsumerGroup {
	consumerGroup.errorHandler = errorHandler
	return consumerGroup
}

func (consumerGroup *ConsumerGroup) StartConsume(ctx context.Context, handler sarama.ConsumerGroupHandler) error {
	client, err := sarama.NewConsumerGroup(
		consumerGroup.config.Hosts(),
		consumerGroup.groupName,
		consumerGroup.config.Config(),
	)
	if err != nil {
		return err
	}

	consumerGroup.client = client

	destroyer.Get().Add(func() {
		if err = client.Close(); err != nil {
			log.Error("Close Kafka Consumer connection error: ", err)
		}
	})

	return consumerGroup.handleConsumers(ctx, handler)
}

func (consumerGroup *ConsumerGroup) StartConsumeAsync(ctx context.Context, handler sarama.ConsumerGroupHandler) {
	go func() {
		if err := consumerGroup.StartConsume(ctx, handler); err != nil {
			log.Fatal("Start consume error: ", err)
		}
	}()
}
