package main

type Service interface {
	Execute() string
}

type ServiceA struct{}

func (s *ServiceA) Execute() string {
	return "Service A"
}

type ServiceB struct{}

func (s *ServiceB) Execute() string {
	return "Service B"
}
