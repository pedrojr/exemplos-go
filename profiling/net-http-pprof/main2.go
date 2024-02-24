package main

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"time"
)

type MyStruct struct {
	data []byte
}

func leakingRoutine() {
	var data []*MyStruct
	for i := 1; i <= 10000; i++ {
		data = append(data, &MyStruct{data: make([]byte, 1024*1024)})
	}
}

func slowRoutine() {
	for i := 1; i <= 5; i++ {
		time.Sleep(1 * time.Second)
		fmt.Println("Iteração:", i)
	}
}

func main() {
	go func() {
		fmt.Println("Servidor pprof iniciado em http://localhost:6060/debug/pprof/")
		http.ListenAndServe("localhost:6060", nil)
	}()

	go slowRoutine()
	go leakingRoutine()

	runtime.Goexit()
}
