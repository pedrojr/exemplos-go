package main

import (
	"fmt"
	"sync"
)

var (
	mutex   sync.Mutex
	counter int
)

func count() {
	mutex.Lock()
	counter++
	mutex.Unlock()
}

func countList(n int, waitGroup *sync.WaitGroup) {
	for i := 0; i < n; i++ {
		count()
	}
	waitGroup.Done()
}

func main() {
	waitGroup := sync.WaitGroup{}
	waitGroup.Add(2)

	go countList(10000, &waitGroup)
	go countList(10000, &waitGroup)

	waitGroup.Wait()

	fmt.Printf("Total: %d\n", counter)
}
