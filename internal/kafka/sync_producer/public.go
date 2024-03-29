package sync_producer

import (
	"encoding/json"
	"github.com/IBM/sarama"
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

	return producer.client.SendMessages(messages)
}
