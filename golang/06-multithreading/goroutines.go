package main

import (
	"fmt"
	"time"
)

func task() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	// Thread 2
	go task()

	// Thread principal
	letras := []string{"A", "B", "C", "D", "E"}
	for i := 0; i < 5; i++ {
		fmt.Println(letras[i])
		time.Sleep(1 * time.Second)
	}
}
