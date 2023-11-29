package handler

import (
	"errors"
	"github.com/IBM/sarama"
	"github.com/lowl11/boost2/data/types"
	"github.com/lowl11/boost2/log"
)

// Setup is run at the beginning of a new session, before ConsumeClaim.
func (handler *Handler) Setup(session sarama.ConsumerGroupSession) error {
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
// but before the offsets are committed for the very last time.
func (handler *Handler) Cleanup(session sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim must start a consumer loop of ConsumerGroupClaim's Messages().
// Once the Messages() channel is closed, the Handler must finish its processing
// loop and exit.
func (handler *Handler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	select {
	case message := <-claim.Messages():
		if err := callHandlerFunc(handler.handlerFunc, message, handler.errorHandler); err != nil {
			log.Error(err, "Kafka handler func error")
		} else {
			session.MarkMessage(message, "")
		}
	case <-handler.stopper:
		log.Info("Stopping consumer group")
		return nil
	}

	return nil
}

func callHandlerFunc(handlerFunc types.KafkaConsumerHandler, message *sarama.ConsumerMessage, errorHandler types.ErrorHandler) (err error) {
	defer func() {
		errRecover := recover()
		if errRecover == nil {
			return
		}

		err = errors.New(errRecover.(string))
	}()

	if err = handlerFunc(message); err != nil {
		if errorHandler != nil {
			errorHandler(err)
		}

		return err
	}

	return nil
}
