package main

import (
	"net/http"

	handlers "./handlers"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "short.gg API")
	})

	e.GET("/shorten/:url", handlers.ShortenURL)
	e.GET("/find/:key", handlers.FindURL)

	e.Logger.Fatal(e.Start(":1323"))
}
