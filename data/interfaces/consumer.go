package interfaces

import "github.com/lowl11/boost2/data/types"

type IConsumer interface {
	SetErrorHandler(errorHandler types.ErrorHandler) IConsumer
	StartConsume(handlerFunc types.KafkaConsumerHandler) error
	StartConsumeAsync(handlerFunc types.KafkaConsumerHandler)
}
