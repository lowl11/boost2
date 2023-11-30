package stopper

import (
	"github.com/lowl11/boost2/internal/infrastructure/stopper"
	"os"
)

func Add(stoppers ...func()) {
	stopper.Get().Add(stoppers...)
}

func GetSignals() chan os.Signal {
	return stopper.Get().GetSignals()
}
