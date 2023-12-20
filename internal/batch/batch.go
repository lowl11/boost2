package batch

import (
	"github.com/IBM/sarama"
	"sync"
	"sync/atomic"
	"time"
)

type Batch struct {
	messages    []*sarama.ProducerMessage
	ticker      *time.Ticker
	mx          sync.Mutex
	needProduce *atomic.Bool
}

var instance *Batch
var Size int
var ProducerFunc func(messages []*sarama.ProducerMessage) error

func Get() *Batch {
	if instance != nil {
		return instance
	}

	if Size == 0 {
		Size = 10000
	}

	needProduce := &atomic.Bool{}
	needProduce.Store(false)
	instance = &Batch{
		messages:    make([]*sarama.ProducerMessage, 0, Size),
		needProduce: needProduce,
		ticker:      time.NewTicker(time.Second * 5),
	}
	go instance.waitTicker()
	return instance
}
