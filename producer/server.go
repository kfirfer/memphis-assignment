package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"
	"net/http"
)

func Start(natsClient *nats.Conn) error {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		err := natsClient.Publish("foo", []byte("Hello World"))
		if err != nil {
			return err
		}
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
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
	natsClient, natsErr := GetNatsClient()
	if natsErr != nil {
		fmt.Println(natsErr)
		return
	}
	echoErr := Start(natsClient)
	if echoErr != nil {
		fmt.Println(echoErr)
		return
	}
}
