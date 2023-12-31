package consumer_group

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/lowl11/boost2/data/interfaces"
	"github.com/lowl11/boost2/data/types"
	"github.com/lowl11/boost2/internal/infrastructure/destroyer"
	"github.com/lowl11/boost2/log"
)

func (consumerGroup *ConsumerGroup) Stop() {
	consumerGroup.stopper <- true
}

func (consumerGroup *ConsumerGroup) SetErrorHandler(errorHandler types.ErrorHandler) interfaces.IConsumer {
	consumerGroup.errorHandler = errorHandler
	return consumerGroup
}

func (consumerGroup *ConsumerGroup) StartConsume(ctx context.Context, handlerFunc types.KafkaConsumerHandler) error {
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

	return consumerGroup.handleConsumers(ctx, handlerFunc)
}

func (consumerGroup *ConsumerGroup) StartConsumeAsync(ctx context.Context, handlerFunc types.KafkaConsumerHandler) {
	go func() {
		if err := consumerGroup.StartConsume(ctx, handlerFunc); err != nil {
			log.Fatal("Start consume error: ", err)
		}
	}()
}
