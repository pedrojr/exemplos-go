package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
	"time"
)

var (
	counter int
	mu      sync.Mutex
)

func main() {
	go func() {
		fmt.Println("Servidor pprof iniciado em http://localhost:6060/debug/pprof/")
		log.Fatal(http.ListenAndServe("localhost:6060", nil))
	}()

	go func() {
		for {
			incrementCounter()
			time.Sleep(time.Millisecond * 100)
		}
	}()

	go func() {
		for {
			readCounter()
			time.Sleep(time.Millisecond * 200)
		}
	}()

	select {}
}

func incrementCounter() {
	mu.Lock()
	defer mu.Unlock()
	counter++
}

func readCounter() {
	mu.Lock()
	defer mu.Unlock()
	fmt.Println("Valor do contador:", counter)
}
