package logger

import (
	"time"

	"github.com/IBM/sarama"
	"github.com/lowl11/boost2/pkg/stopper"
)

type LogMessage struct {
	Service string `json:"service,omitempty"`
	Data    string `json:"data,omitempty"`
}

func (logger *Logger) listen() {
	stopper.Add(func() {
		time.Sleep(time.Millisecond * 50)
	})

	go func() {
		for log := range logger.logChannel {
			log()
		}
	}()
}

func (logger *Logger) printLog(logFunc func()) {
	logger.logChannel <- logFunc
}

// KafkaSink implements zapcore.WriteSyncer to send logs to a Kafka topic
type KafkaSink struct {
	producer    sarama.SyncProducer
	topic       string
	serviceName string
}

// NewKafkaSink creates a new KafkaSink
func NewKafkaSink(producer sarama.SyncProducer, topic, serviceName string) *KafkaSink {
	return &KafkaSink{
		serviceName: serviceName,
		producer:    producer,
		topic:       topic,
	}
}

// Write serializes the log entry and sends it to Kafka
func (k *KafkaSink) Write(p []byte) (n int, err error) {
	_, _, err = k.producer.SendMessage(&sarama.ProducerMessage{
		Topic: k.topic,
		Key:   sarama.ByteEncoder(k.serviceName),
		Value: sarama.ByteEncoder(p),
		Headers: []sarama.RecordHeader{{
			Key:   []byte("service"),
			Value: []byte(k.serviceName),
		}},
	})
	if err != nil {
		return 0, err
	}
	return len(p), nil
}

// Sync does nothing for KafkaSink
func (k *KafkaSink) Sync() error {
	return nil
}
