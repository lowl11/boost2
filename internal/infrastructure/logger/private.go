package logger

import (
	"github.com/lowl11/boost2/pkg/stopper"
	"time"
)

func (logger *Logger) listen() {
	stopper.Add(func() {
		time.Sleep(time.Millisecond * 50)
	})

	go func() {
		for log := range logger.logChannel {
			log()
		}
	}()
}

func (logger *Logger) printLog(logFunc func()) {
	logger.logChannel <- logFunc
}
