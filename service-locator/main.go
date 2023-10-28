package main

import (
	"fmt"
)

func main() {
	locator := initServiceLocator()

	serviceA := locator.GetService("ServiceA")
	resultA := serviceA.Execute()
	fmt.Println("Resultado Service A:", resultA)

	serviceB := locator.GetService("ServiceB")
	resultB := serviceB.Execute()
	fmt.Println("Resultado Service B:", resultB)
}

func initServiceLocator() *ServiceLocator {
	locator := NewServiceLocator()
	locator.RegisterService("ServiceA", &ServiceA{})
	locator.RegisterService("ServiceB", &ServiceB{})
	return locator
}
