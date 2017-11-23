package config

import (
	"github.com/garyburd/redigo/redis"
)

// RedisConnect connects to redis
func RedisConnect() redis.Conn {
	c, err := redis.Dial("tcp", ":6379")
	handleError(err)
	return c
}

func handleError(err error) {
	if err != nil {
		panic(err)
	}
}
