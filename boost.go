package boost2

import (
	"github.com/lowl11/boost2/data/interfaces"
	"github.com/lowl11/boost2/internal/infrastructure/stopper"
	"github.com/lowl11/boost2/log"
)

type (
	IConsumer = interfaces.IConsumer
	IProducer = interfaces.IProducer
)

type Config struct {
	//
}

type Boost struct {
	infinite chan bool
}

func New(config ...Config) *Boost {
	app := &Boost{
		infinite: make(chan bool, 1),
	}
	stopper.Get().Add(app.Stop)
	return app
}

func (boost *Boost) Run() {
	log.Info("App started")
	<-boost.infinite
	log.Infof("App stopped")
}

func (boost *Boost) Stop() {
	boost.infinite <- true
}
