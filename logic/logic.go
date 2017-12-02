package logic

import (
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"time"

	config "../config"
	models "../models"
)

// Shorten shortens and store an url in redis
func Shorten(url string) (string, error) {
	/* open connection to redis */
	c := config.RedisConnect()
	/* without forgetting to close it afterwards */
	defer c.Close()

	/* we find the next key */
	nextKey, err := c.Do("INCR", "redirection")
	if err != nil {
		return "", err
	}

	/* filling the struct */
	var p models.ShortenedURL

	p.Key = fmt.Sprintf("%v", nextKey)
	p.Timestamp = time.Now()

	matched, _ := regexp.MatchString("^http[s]?://.*$", url)
	if !matched {
		url = "http://" + url
	}

	p.RedirectTo = url

	// marshal models.ShortenedURL to JSON blob
	b, err := json.Marshal(p)
	if err != nil {
		return "", err
	}

	// save json to redis
	reply, err := c.Do("SET", "redirection:"+p.Key, b)
	if err != nil {
		return "", err
	}
	/* we dont care actually about the reply if the set succeeded*/
	_ = reply

	return p.Key, nil
}

// Find find a shortened url in redis
func Find(key string) (models.ShortenedURL, error) {
	var shortenedURL models.ShortenedURL

	c := config.RedisConnect()
	defer c.Close()

	reply, err := c.Do("GET", "redirection:"+key)
	if reply == nil || err != nil {
		return shortenedURL, errors.New("not found")
	}
	if err := json.Unmarshal(reply.([]byte), &shortenedURL); err != nil {
		return shortenedURL, err
	}
	return shortenedURL, nil
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
