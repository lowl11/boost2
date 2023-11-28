package destroyer

type Service struct {
	destroyFunctions []func()
}

var instance *Service

func Get() *Service {
	if instance != nil {
		return instance
	}

	instance = &Service{
		destroyFunctions: make([]func(), 0),
	}
	return instance
}
