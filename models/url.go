package models

import (
	"time"
)

// ShortenedURL stored in redis
type ShortenedURL struct {
	Key        string    `json:"key"`
	RedirectTo string    `json:"redirectTo"`
	Timestamp  time.Time `json:"timestamp"`
}
