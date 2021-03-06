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

	e.Use(middleware.Logger())

	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("CORS_URL"), "chrome-extension://lalhakoplnjfhbeigdgcnefcfmhiignc"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE, echo.OPTIONS},
	}))

	e.GET("/api/", func(c echo.Context) error {
		return c.String(http.StatusOK, "short.gg API")
	})

	e.POST("/api/shorten", handlers.ShortenURL)
	e.GET("/api/find/:key", handlers.FindURL)

	e.GET("/:key", handlers.Redirect)

	e.Logger.Fatal(e.Start(":1323"))
}
