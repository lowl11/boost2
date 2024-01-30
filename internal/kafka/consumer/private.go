package consumer

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/lowl11/boost2/data/types"
	"github.com/lowl11/boost2/log"
	"sync"
)

func (consumer *Consumer) handleConsumers(ctx context.Context, partitions []int32, handlerFunc types.KafkaConsumerHandler) {
	goroutines := make([]*sync.WaitGroup, len(partitions))
	consumer.stoppers = make([]chan bool, 0, len(partitions))

	for i := 0; i < len(partitions); i++ {
		goroutines[i] = &sync.WaitGroup{}
		goroutines[i].Add(1)
		consumer.stoppers = append(consumer.stoppers, make(chan bool, 1))
	}

	for i := 0; i < len(goroutines); i++ {
		go consumer.handleConsumerFunc(ctx, goroutines[i], int32(i), handlerFunc)
	}
}

func (consumer *Consumer) handleConsumerFunc(ctx context.Context, wg *sync.WaitGroup, partitionNum int32, handlerFunc types.KafkaConsumerHandler) {
	defer wg.Done()

	partConsumer, err := consumer.client.ConsumePartition(consumer.topicName, partitionNum, consumer.config.Config().Consumer.Offsets.Initial)
	if err != nil {
		log.Error("Consuming partition error: ", err)
	}

	for {
		select {
		case message := <-partConsumer.Messages():
			callHandlerFunc(handlerFunc, message, consumer.errorHandler)
		case kafkaError := <-partConsumer.Errors():
			log.Error("Kafka consumer error: ", kafkaError.Error(), ". Partition: ", kafkaError.Partition)
		case <-consumer.stoppers[partitionNum]:
			log.Info("Stopping consumer by stopper with partition #", partitionNum+1)
			if err = partConsumer.Close(); err != nil {
				log.Error("Close partition consumer error: ", err)
			}
			return
		case <-ctx.Done():
			log.Info("Stopping consumer by context with partition #", partitionNum+1)
			return
		}
	}
}

func callHandlerFunc(handlerFunc types.KafkaConsumerHandler, message *sarama.ConsumerMessage, errorHandler types.ErrorHandler) {
	defer func() {
		err := recover()
		if err == nil {
			return
		}

		log.Error("Kafka consumer handler catch PANIC: ", err)
	}()

	if err := handlerFunc(message); err != nil {
		log.Error("Kafka handler function error: ", err, ". Partition: ", message.Partition)
		if errorHandler != nil {
			errorHandler(err)
		}
	}
}
