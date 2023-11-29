package handler

import "github.com/lowl11/boost2/data/types"

type Handler struct {
	handlerFunc  types.KafkaConsumerHandler
	stopper      chan bool
	errorHandler types.ErrorHandler
}

func New(handlerFunc types.KafkaConsumerHandler, errorHandler types.ErrorHandler, stopper chan bool) *Handler {
	return &Handler{
		handlerFunc:  handlerFunc,
		errorHandler: errorHandler,
		stopper:      stopper,
	}
}
