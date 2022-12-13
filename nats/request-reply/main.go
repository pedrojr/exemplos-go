package main

import (
	"fmt"
	"runtime"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		fmt.Println(err)
	}

	subj := "my_subj"
	Subscribe(nc, subj)
	Request(nc, subj)

	runtime.Goexit()
}

func Subscribe(nc *nats.Conn, subj string) {
	nc.Subscribe(subj, func(msg *nats.Msg) {
		fmt.Println("Subscribe.received: ", string(msg.Data))
		nc.Publish(msg.Reply, []byte("1.0"))
		msg.Respond([]byte("1.0"))
	})
}

func Request(nc *nats.Conn, subj string) {
	reply, err := nc.Request(subj, []byte("-v"), 4*time.Second)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Request.reply: ", string(reply.Data))
}
