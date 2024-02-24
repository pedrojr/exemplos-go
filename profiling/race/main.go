package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	counter int
	mu      sync.Mutex
)

func main() {
	for i := 0; i < 10; i++ {
		go incrementCounter(i)
	}

	time.Sleep(time.Second * 5)
	fmt.Println("Valor final do contador:", counter)
}

func incrementCounter(id int) {
	fmt.Printf("Goroutine %d tentando adquirir o lock...\n", id)

	mu.Lock()
	defer mu.Unlock()
	time.Sleep(time.Second)

	counter++
	fmt.Printf("Goroutine %d incrementou o contador\n", id)
}
