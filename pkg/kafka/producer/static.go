package producer

import (
	"github.com/lowl11/boost2"
	"github.com/lowl11/boost2/internal/kafka/async_producer"
	"github.com/lowl11/boost2/internal/kafka/mock_producer"
	"github.com/lowl11/boost2/internal/kafka/sync_producer"
	"github.com/lowl11/boost2/pkg/kafka/configurator"
)

func NewSync(config *configurator.Configurator) (boost2.IProducer, error) {
	return sync_producer.New(config)
}

func NewAsync(config *configurator.Configurator) (boost2.IProducer, error) {
	return async_producer.New(config)
}

func NewMock(config *configurator.Configurator) (boost2.IProducer, error) {
	return mock_producer.New(config)
}
