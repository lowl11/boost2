package types

import "github.com/IBM/sarama"

type KafkaConsumerHandler = func(message *sarama.ConsumerMessage) error
