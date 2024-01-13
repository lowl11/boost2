package boost2

import (
	"context"
	"github.com/lowl11/boost2/internal/infrastructure/stopper"
	"github.com/lowl11/boost2/log"
)

type Config struct {
	//
}

type Boost struct {
	ctx    context.Context
	cancel func()
}

func New(config ...Config) *Boost {
	ctx, cancel := context.WithCancel(context.Background())
	app := &Boost{
		ctx:    ctx,
		cancel: cancel,
	}
	stopper.Get().Add(app.Stop)
	return app
}

func (boost *Boost) Run() {
	<-stopper.Get().GetSignals()
	stopper.Get().Run()
}

func (boost *Boost) Context() context.Context {
	return boost.ctx
}

func (boost *Boost) Stop() {
	boost.cancel()
	log.Info("App stopped")
}
