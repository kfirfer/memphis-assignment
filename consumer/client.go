package main

import (
	"fmt"
)

func Start(natsClient *nats.Conn) error {
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
