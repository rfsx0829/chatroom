package redis

import (
	"fmt"

	"github.com/go-redis/redis"
)

// Client redis.Client
var Client *redis.Client

// InitClient init client
func InitClient(host string, port int, password string, DB int) {
	opt := redis.Options{
		Addr:     fmt.Sprintf("%s:%d", host, port),
		Password: password,
		DB:       DB,
	}

	Client = redis.NewClient(&opt)
}
