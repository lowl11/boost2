package producer

import (
	"github.com/lowl11/boost2/internal/kafka/async_producer"
	"github.com/lowl11/boost2/internal/kafka/mock_producer"
	"github.com/lowl11/boost2/internal/kafka/sync_producer"
	"github.com/lowl11/boost2/pkg/kafka/configurator"
)

func NewSync(config *configurator.Configurator) (*sync_producer.Producer, error) {
	return sync_producer.New(config)
}

func NewAsync(config *configurator.Configurator) (*async_producer.Producer, error) {
	return async_producer.New(config)
}

func NewMock(config *configurator.Configurator) (*mock_producer.Producer, error) {
	return mock_producer.New(config)
}
