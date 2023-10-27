package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

func main() {
	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	for i := 0; i < 10; i++ {
		go performTask(i)
		time.Sleep(time.Second)
	}
}

func performTask(id int) {
	if id == 5 {
		runtime.MemProfileRate = 1
	}
	defer profileRuntime()()
	
	data := make([]byte, 1024*1024)
	time.Sleep(2 * time.Second)
	_ = data
}

func profileRuntime() func() {
	start := time.Now()
	return func() {
		fmt.Printf("Tempo de execução da rotina: %v\n", time.Since(start))
	}
}
