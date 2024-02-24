package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mutex   sync.Mutex
	counter int
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

	start := time.Now()
	mutex.Lock()
	defer mutex.Unlock()
	elapsed := time.Since(start)

	time.Sleep(time.Second)

	counter++

	fmt.Printf("Goroutine %d adquiriu o lock após %s e incrementou o contador\n", id, elapsed)
}
