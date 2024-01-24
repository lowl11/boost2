package consumer

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/lowl11/boost2/data/types"
	"github.com/lowl11/boost2/internal/infrastructure/destroyer"
	"github.com/lowl11/boost2/log"
)

func (consumer *Consumer) Stop() {
	for _, stopper := range consumer.stoppers {
		stopper <- true
	}
}

func (consumer *Consumer) SetErrorHandler(errorHandler types.ErrorHandler) *Consumer {
	consumer.errorHandler = errorHandler
	return consumer
}

func (consumer *Consumer) StartConsume(ctx context.Context, handlerFunc types.KafkaConsumerHandler) error {
	client, err := sarama.NewConsumer(consumer.config.Hosts(), consumer.config.Config())
	if err != nil {
		return err
	}

	consumer.client = client

	destroyer.Get().Add(func() {
		if err = client.Close(); err != nil {
			log.Error("Close Kafka Consumer connection error: ", err)
		}
	})

	partitions, err := client.Partitions(consumer.topicName)
	if err != nil {
		return err
	}

	log.Debugf("Consume %d partitions", len(partitions))

	consumer.handleConsumers(ctx, partitions, handlerFunc)

	return nil
}

func (consumer *Consumer) StartConsumeAsync(ctx context.Context, handlerFunc types.KafkaConsumerHandler) {
	go func() {
		if err := consumer.StartConsume(ctx, handlerFunc); err != nil {
			log.Error("Start Kafka consumer error: ", err)
		}
	}()
}
