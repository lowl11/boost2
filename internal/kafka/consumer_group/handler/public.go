package handler

import (
	"errors"
	"fmt"
	"github.com/IBM/sarama"
	"github.com/lowl11/boost2/data/types"
	"github.com/lowl11/boost2/log"
)

func (handler *Handler) SetAlwaysCommit() {
	handler.alwaysCommit = true
}

// Setup is run at the beginning of a new session, before ConsumeClaim.
func (handler *Handler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
// but before the offsets are committed for the very last time.
func (handler *Handler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish its processing
// loop and exit.
func (handler *Handler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		if err := callHandlerFunc(handler.handlerFunc, message, handler.errorHandler); err != nil {
			log.Error("Kafka handler func error: ", err)
		}

		session.MarkMessage(message, "")
	}

	return nil
}

func callHandlerFunc(handlerFunc types.KafkaConsumerHandler, message *sarama.ConsumerMessage, errorHandler types.ErrorHandler) (err error) {
	defer func() {
		errRecover := recover()
		if errRecover == nil {
			return
		}

		errStr, ok := errRecover.(string)
		if ok {
			err = errors.New(errStr)
			return
		}

		errErr, ok := errRecover.(error)
		if ok {
			err = errErr
			return
		}

		err = errors.New(fmt.Sprintf("%s", err))
	}()

	if err = handlerFunc(message); err != nil {
		if errorHandler != nil {
			errorHandler(err)
		}

		return err
	}

	return nil
}
