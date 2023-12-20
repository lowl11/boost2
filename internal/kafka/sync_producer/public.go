package sync_producer

import (
	"encoding/json"
	"github.com/IBM/sarama"
	"github.com/lowl11/boost2/internal/batch"
)

func (producer *Producer) Publish(topic, key string, objects ...any) error {
	if len(objects) == 0 {
		return nil
	}

	messages := make([]*sarama.ProducerMessage, 0, len(objects))

	for _, obj := range objects {
		objectInBytes, err := json.Marshal(obj)
		if err != nil {
			return err
		}

		messages = append(messages, &sarama.ProducerMessage{
			Topic: topic,
			Key:   sarama.ByteEncoder(key),
			Value: sarama.ByteEncoder(objectInBytes),
		})
	}

	batch.Size = producer.batchSize
	batch.ProducerFunc = func(messages []*sarama.ProducerMessage) error {
		return producer.client.SendMessages(messages)
	}

	if producer.isBatch {
		return batch.Get().Produce(messages...)
	}

	return producer.client.SendMessages(messages)
}
