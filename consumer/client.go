package main

import (
	"context"
	"fmt"
	"github.com/nats-io/nats.go"
	"time"
)

func main() {
	nc, _ := nats.Connect(nats.DefaultURL, nats.RetryOnFailedConnect(true),
		nats.MaxReconnects(10),
		nats.ReconnectWait(time.Second),
	)
	js, _ := nc.JetStream(nats.PublishAsyncMaxPending(256))

	sub, _ := js.PullSubscribe("ORDERS.*", "order-review", nats.PullMaxWaiting(128))
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
			msg.Ack()
			fmt.Println(string(msg.Data[:]))
		}
	}
}
