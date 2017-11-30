package config

import (
	"os"

	"github.com/garyburd/redigo/redis"
)

// RedisConnect connects to redis
func RedisConnect() redis.Conn {
	c, err := redis.Dial("tcp", os.Getenv("REDIS_URL"))
	handleError(err)
	return c
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
