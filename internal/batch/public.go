package batch

import "github.com/IBM/sarama"

func (batch *Batch) Produce(messages ...*sarama.ProducerMessage) error {
	batch.mx.Lock()
	defer batch.mx.Unlock()

	if batch.needProduce.Load() {
		if err := ProducerFunc(batch.messages); err != nil {
			return err
		}

		batch.
			clearMessages().
			refreshTicker()
	}

	batch.messages = append(batch.messages, messages...)
	if len(batch.messages) >= Size {
		batch.needProduce.Store(true)
	}

	return nil
}
