package handlers

import (
	"net/http"

  logic "../logic"

	"github.com/labstack/echo"
)

// ShortenURL
func ShortenURL(c echo.Context) error {
  url := c.Param("url")

  shortenedURL, err := logic.Shorten(url)
  if err != nil {
    return echo.NewHTTPError(http.StatusTeapot)
  }

  return c.JSON(http.StatusOK, shortenedURL)
}

// FindURL find an URL by its key
func FindURL(c echo.Context) error {
  key := c.Param("key")

  url, err := logic.Find(key)
  if err != nil {
    return echo.NewHTTPError(http.StatusTeapot)
  }

  return c.JSON(http.StatusOK, url)
}
