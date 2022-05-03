package main

import (
	"fmt"
	"github.com/nats-io/nats.go"
)

func Start(natsClient *nats.Conn) error {
	//sub, err := natsClient.SubscribeSync("foo")
	//if err != nil {
	//	return err
	//}
	_, err := natsClient.Subscribe("foo", func(m *nats.Msg) {
		fmt.Printf("Received a message: %s\n", string(m.Data))
	})
	if err != nil {
		return err
	}

	//message, err := sub.NextMsg(60)
	//fmt.Println("message: ", message)
	return nil
}

func GetNatsClient() (*nats.Conn, error) {
	natsClient, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, err
	}
	return natsClient, nil

}
func main() {
	natsClient, natsConnectErr := GetNatsClient()
	if natsConnectErr != nil {
		fmt.Println(natsConnectErr)
		return
	}
	natsErr := Start(natsClient)
	if natsErr != nil {
		fmt.Println(natsErr)
		return
	}
}
