package batch

import (
	"github.com/IBM/sarama"
	"sync"
)

type Batch struct {
	messages []*sarama.ProducerMessage
	mx       sync.Mutex
}

var instance *Batch
var Size int

func Get() *Batch {
	if instance != nil {
		return instance
	}

	if Size == 0 {
		Size = 10000
	}

	instance = &Batch{
		messages: make([]*sarama.ProducerMessage, 0, Size),
	}
	return instance
}
