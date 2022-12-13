package main

import (
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
	QueueSubscribe(nc, subj)

	Publish(nc, subj, "Msg1")
	Publish(nc, subj, "Msg2")
	Publish(nc, subj, "Msg3")

	runtime.Goexit()
}

func QueueSubscribe(nc *nats.Conn, subj string) {
	nc.QueueSubscribe(subj, "Sub1", func(msg *nats.Msg) {
		fmt.Println("Sub1: ", string(msg.Data))
	})

	nc.QueueSubscribe(subj, "Sub2", func(msg *nats.Msg) {
		fmt.Println("Sub2: ", string(msg.Data))
	})

	nc.QueueSubscribe(subj, "Sub3", func(msg *nats.Msg) {
		fmt.Println("Sub3: ", string(msg.Data))
	})
}

func Publish(nc *nats.Conn, subj string, msg string) {
	nc.Publish(subj, []byte(msg))
	fmt.Println("Publish: ", msg)
}
