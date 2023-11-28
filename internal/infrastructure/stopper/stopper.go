package stopper

type Service struct {
	stoppers []func()
}

var instance *Service

func Get() *Service {
	if instance != nil {
		return instance
	}

	instance = &Service{
		stoppers: make([]func(), 0),
	}
	return instance
}
