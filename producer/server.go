package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func Start() error {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":1323"))
	return nil
}

func main() {
	err := Start()
	if err != nil {
		return
	}
}
