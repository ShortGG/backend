package handlers

import (
	b64 "encoding/base64"
	"net/http"
	"strconv"

	logic "../logic"

	"github.com/labstack/echo"
)

// ShortenURL shortens url right
func ShortenURL(c echo.Context) error {
	url := c.Param("url")

	decodedURL, _ := b64.StdEncoding.DecodeString(url)

	shortenedURL, err := logic.Shorten(string(decodedURL))
	if err != nil {
		return echo.NewHTTPError(http.StatusTeapot)
	}

	// Encode in B58 the url key
	key, _ := strconv.Atoi(shortenedURL)
	hash := logic.EncodeB58(key)

	return c.JSON(http.StatusOK, hash)
}

// FindURL find an URL by its key
func FindURL(c echo.Context) error {
	key := c.Param("key")

	decodedKey := logic.DecodeB58(key)
	index := strconv.Itoa(decodedKey)

	url, err := logic.Find(index)
	if err != nil {
		return echo.NewHTTPError(http.StatusTeapot)
	}

	return c.JSON(http.StatusOK, url)
}


// Redirect find an URL by its key
func Redirect(c echo.Context) error {
	key := c.Param("key")

	url, err := logic.Find(key)
	if err != nil {
		// instead return there 404 page
		return echo.NewHTTPError(http.StatusTeapot)
	}
	return c.Redirect(http.StatusMovedPermanently, url.RedirectTo)
}
