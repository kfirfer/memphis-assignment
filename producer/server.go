package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"
	"io"
	"net/http"
	"strings"
	"time"
)

func Start(jsClient nats.JetStreamContext) error {
	e := echo.New()
	e.POST("sendMessage", func(c echo.Context) error {
		bodyReader := c.Request().Body
		buf := new(strings.Builder)
		_, err := io.Copy(buf, bodyReader)

		_, err = jsClient.PublishAsync("ORDERS.*", []byte(buf.String()))
		select {
		case <-jsClient.PublishAsyncComplete():
		case <-time.After(5 * time.Second):
			fmt.Println("Did not resolve in time")
		}
		if err != nil {
			fmt.Println(err)
			return err
		}
		if err != nil {
			fmt.Println(err)
			return err
		}
		return c.String(http.StatusOK, "Sent to queue")
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
	js, _ := natsClient.JetStream(nats.PublishAsyncMaxPending(256))
	js.AddStream(&nats.StreamConfig{
		Name:     "ORDERS",
		Subjects: []string{"ORDERS.*"},
	})

	if natsErr != nil {
		fmt.Println(natsErr)
		return
	}
	echoErr := Start(js)
	if echoErr != nil {
		fmt.Println(echoErr)
		return
	}
}
