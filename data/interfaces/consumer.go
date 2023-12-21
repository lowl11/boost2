package interfaces

import (
	"context"
	"github.com/lowl11/boost2/data/types"
)

type IConsumer interface {
	SetErrorHandler(errorHandler types.ErrorHandler) IConsumer
	StartConsume(ctx context.Context, handlerFunc types.KafkaConsumerHandler) error
	StartConsumeAsync(ctx context.Context, handlerFunc types.KafkaConsumerHandler)
}
