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
		CheckErr(err)
		_, err = jsClient.PublishAsync("messages", []byte(buf.String()))
		CheckErr(err)
		select {
		case <-jsClient.PublishAsyncComplete():
		case <-time.After(10 * time.Second):
			fmt.Println("Did not resolve in time")
		}
		return c.String(http.StatusOK, "Sent to queue")
	})
	e.Logger.Fatal(e.Start(":1323"))
	return nil
}

func GetNatsClient() (*nats.Conn, error) {
	natsClient, err := nats.Connect(nats.DefaultURL)
	CheckErr(err)
	return natsClient, nil

}
func main() {
	natsClient, err := GetNatsClient()
	CheckErr(err)
	js, err := natsClient.JetStream(nats.PublishAsyncMaxPending(256))
	CheckErr(err)
	_, err = js.AddStream(&nats.StreamConfig{
		Name:     "MESSAGES",
		Subjects: []string{"messages"},
	})
	CheckErr(err)
	echoErr := Start(js)
	CheckErr(echoErr)
}

func CheckErr(err error) {
	if err != nil {
		panic(err)
	}
}
