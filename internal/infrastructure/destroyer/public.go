package destroyer

func (service *Service) Add(destroyFunc func()) *Service {
	service.destroyFunctions = append(service.destroyFunctions, destroyFunc)
	return service
}

func (service *Service) Run() {
	for _, destroyFunction := range service.destroyFunctions {
		runDestroyFunc(destroyFunction)
	}
}
