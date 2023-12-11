package consumer_group

import (
	"context"
	"github.com/lowl11/boost2/data/types"
	"github.com/lowl11/boost2/internal/kafka/consumer_group/handler"
	"github.com/lowl11/boost2/log"
)

func (consumerGroup *ConsumerGroup) handleConsumers(handlerFunc types.KafkaConsumerHandler) error {
	for _, topic := range consumerGroup.topicNames {
		partitions, err := consumerGroup.innerClient.Partitions(topic)
		if err != nil {
			return err
		}

		if consumerGroup.stoppers == nil {
			consumerGroup.stoppers = make([]chan bool, 0, len(partitions))
		}

		consumerGroup.stoppers = append(consumerGroup.stoppers, make(chan bool, 1))

		for i := 0; i < len(partitions); i++ {
			go func(topic string) {
				err = consumerGroup.client.Consume(
					context.Background(),
					consumerGroup.topicNames,
					handler.New(handlerFunc, consumerGroup.errorHandler, consumerGroup.stoppers[0]),
				)
				if err != nil {
					log.Fatal("Start consuming topic ", topic, " error: ", err)
				}
			}(topic)
		}
	}

	return nil
}
