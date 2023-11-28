package stopper

import "github.com/lowl11/boost2/internal/infrastructure/stopper"

func Add(stoppers ...func()) {
	stopper.Get().Add(stoppers...)
}
