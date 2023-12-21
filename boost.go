package boost2

import (
	"context"
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
	ctx      context.Context
	cancel   func()
	infinite chan bool
}

func New(config ...Config) *Boost {
	ctx, cancel := context.WithCancel(context.Background())
	app := &Boost{
		infinite: make(chan bool, 1),
		ctx:      ctx,
		cancel:   cancel,
	}
	stopper.Get().Add(app.Stop)
	return app
}

func (boost *Boost) Run() {
	log.Info("App started")
	<-boost.infinite
	log.Infof("App stopped")
}

func (boost *Boost) Context() context.Context {
	return boost.ctx
}

func (boost *Boost) Stop() {
	boost.infinite <- true
	boost.cancel()
}
