package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/nats-io/nats.go"
	"io"
	"net/http"
	"strings"
)

func Start(natsClient *nats.Conn) error {
	e := echo.New()
	e.POST("sendMessage", func(c echo.Context) error {
		bodyReader := c.Request().Body
		buf := new(strings.Builder)
		_, err := io.Copy(buf, bodyReader)
		err = natsClient.Publish("messages", []byte(buf.String()))
		if err != nil {
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
