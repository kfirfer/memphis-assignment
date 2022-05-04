package main

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"log"
	"time"
)

func main() {
	nc, err := nats.Connect(nats.DefaultURL, nats.RetryOnFailedConnect(true),
		nats.MaxReconnects(10),
		nats.ReconnectWait(time.Second),
	)
	if err != nil {
		log.Fatal(err)
		return
	}
	js, err := nc.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		log.Fatal(err)
		return
	}
	sub, err := js.PullSubscribe("messages", "messages-durable", nats.PullMaxWaiting(128))
	if err != nil {
		log.Fatal(err)
		return
	}
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
			if err != nil {
				log.Fatal(err)
				return
			}
			fmt.Println(string(msg.Data[:]))
		}
	}
}
