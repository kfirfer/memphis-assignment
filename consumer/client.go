package main

import (
	"encoding/json"
	"fmt"
	"github.com/nats-io/nats.go"
	"github.com/novalagung/natskeepalivesubscribe"
	"strings"
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
	natsErr := Start()
	if natsErr != nil {
		fmt.Println(natsErr)
		return
	}
}
