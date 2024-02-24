package main

import (
	"fmt"
	"log"
	"os"
	"runtime/pprof"
	"sync"
	"time"
)

var (
	counter int
	mu      sync.Mutex
)

func main() {
	f, err := os.Create("mutex_profile.prof")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	if err := pprof.StartCPUProfile(f); err != nil {
		log.Fatal(err)
	}
	defer pprof.StopCPUProfile()

	for i := 0; i < 10; i++ {
		go incrementCounter(i)
	}
	time.Sleep(time.Second * 5)

	fmt.Println("Valor final do contador:", counter)
}

func incrementCounter(id int) {
	fmt.Printf("Goroutine %d tentando adquirir o lock...\n", id)

	start := time.Now()
	mu.Lock()
	defer mu.Unlock()
	elapsed := time.Since(start)
	time.Sleep(time.Second)

	counter++
	fmt.Printf("Goroutine %d adquiriu o lock após %s e incrementou o contador\n", id, elapsed)
}
