package main

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/novalagung/natskeepalivesubscribe"
	"strings"
	"time"
)

func Start() error {
	natsURL := "nats://localhost:4222"
	subject := "messages"
	natskeepalivesubscribe.KeepAliveSubscribe(natsURL, subject, func(msg *nats.Msg) (interface{}, error) {
		// parse payload
		payload := make(map[string]interface{})
		err := json.Unmarshal(msg.Data, &payload)
		if err != nil {
			return nil, err
		}

		// handle the request
		switch strings.ToUpper(payload["method"].(string)) {
		case "OPTIONS":
			// ...
		case "GET":
			// ...
		case "POST":
			fmt.Println(payload["message"])
			msg.Reply = "ok"
		case "PATCH":
			// ...
		case "PUT":
			// ...
		case "DELETE":
			// ...
		}
		return nil, fmt.Errorf("invalid http method")
	})
	return nil
}

func main() {
	nc, _ := nats.Connect(nats.DefaultURL, nats.RetryOnFailedConnect(true),
		nats.MaxReconnects(10),
		nats.ReconnectWait(time.Second),
	)
	js, _ := nc.JetStream(nats.PublishAsyncMaxPending(256))

	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//sub, err := nc.SubscribeSync("ORDERS.*", nc.Durable("MONITOR"), nats.MaxDeliver(3))

	//_, err := nc.QueueSubscribe("messages", "messages", func(m *nats.Msg) {
	//	fmt.Printf("Received a message: %s\n", string(m.Data))
	//
	//})

	_, err := js.Subscribe("ORDERS.*", func(m *nats.Msg) {
		fmt.Printf("Received a JetStream message: %s\n", string(m.Data))
	})
	if err != nil {
		fmt.Println(err)
		return
	}
	for true {
		//sub, _ := js.SubscribeSync("ORDERS.*", nats.Durable("ORDERS"), nats.MaxDeliver(3))
		//m, _ := sub.NextMsg(60)
		//fmt.Println(m)

		time.Sleep(time.Millisecond * 100)
	}
	// Tests
	//natsErr := Start()
	//if natsErr != nil {
	//	fmt.Println(natsErr)
	//	return
	//}
}
