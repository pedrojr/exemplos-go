package main

type ServiceLocator struct {
	services map[string]Service
}

func NewServiceLocator() *ServiceLocator {
	return &ServiceLocator{
		services: make(map[string]Service),
	}
}

func (sl *ServiceLocator) RegisterService(name string, service Service) {
	sl.services[name] = service
}

func (sl *ServiceLocator) GetService(name string) Service {
	return sl.services[name]
}
