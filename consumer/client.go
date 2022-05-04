package main

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"time"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL, nats.RetryOnFailedConnect(true),
		nats.MaxReconnects(10),
		nats.ReconnectWait(time.Second),
	)
	CheckErr(err)
	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	CheckErr(err)
	sub, err := js.PullSubscribe("messages", "messages-durable", nats.PullMaxWaiting(128))
	CheckErr(err)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	for {
		select {
		case <-ctx.Done():
			return
		default:
		}
		msgs, _ := sub.Fetch(10, nats.Context(ctx))
		for _, msg := range msgs {
			err := msg.Ack()
			CheckErr(err)
			fmt.Println(string(msg.Data[:]))
		}
	}
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
