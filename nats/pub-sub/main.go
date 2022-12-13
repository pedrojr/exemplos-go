package main

import (
	"encoding/json"
	"fmt"
	"runtime"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Println(err)
	}

	subj := "my_subj"
	SubscribeTest(nc, subj)
	PublishTest(nc, subj)
	
	runtime.Goexit()
}

type Test struct {
	Id  int
	Msg string
}

func PublishTest(nc *nats.Conn, subj string) {
	test := Test{Id: 1001, Msg: "Test"}
	jsonTest, _ := json.Marshal(test)
	fmt.Printf("PublishTest: %v\n", test)
	nc.Publish(subj, jsonTest)
}

func SubscribeTest(nc *nats.Conn, subj string) {
	nc.Subscribe(subj, func(m *nats.Msg) {
		var test Test
		json.Unmarshal(m.Data, &test)
		fmt.Printf("SubscribeTest: %v\n", test)
	})
}
