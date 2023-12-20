package batch

import "github.com/IBM/sarama"

func (batch *Batch) Append(messages ...*sarama.ProducerMessage) bool {
	batch.mx.Lock()
	defer batch.mx.Unlock()

	batch.messages = append(batch.messages, messages...)
	return len(batch.messages) == Size
}

func (batch *Batch) Get() []*sarama.ProducerMessage {
	batch.mx.Lock()
	defer batch.mx.Unlock()

	return batch.messages
}

func (batch *Batch) Clear() {
	batch.messages = make([]*sarama.ProducerMessage, 0, Size)
}
