package main

import (
	"net/http"
	"os"

	handlers "./handlers"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("CORS_URL")},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
	}))

	e.GET("/api/", func(c echo.Context) error {
		return c.String(http.StatusOK, "short.gg API")
	})

	e.GET("/api/shorten/:url", handlers.ShortenURL)
	e.GET("/api/find/:key", handlers.FindURL)

	e.Logger.Fatal(e.Start(":1323"))
}
