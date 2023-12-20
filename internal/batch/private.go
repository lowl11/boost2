package batch

import (
	"github.com/IBM/sarama"
	"github.com/lowl11/boost2/log"
	"time"
)

func (batch *Batch) clearMessages() *Batch {
	batch.messages = make([]*sarama.ProducerMessage, 0, Size)
	return batch
}

func (batch *Batch) refreshTicker() *Batch {
	batch.ticker.Reset(time.Second * 5)
	batch.needProduce.Store(false)
	return batch
}

func (batch *Batch) waitTicker() {
	<-batch.ticker.C
	batch.needProduce.Store(true)

	if err := batch.Produce(); err != nil {
		log.Error("Produce batch messages after expiring time error: ", err)
	}
}
