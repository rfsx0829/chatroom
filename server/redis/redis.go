package redis

import (
	"fmt"

	"github.com/go-redis/redis"
)

// Option expose config of redis
type Option struct {
	Host string `json:"host"`
	Port int    `json:"port"`
	Pass string `json:"pass"`
	DB   int    `json:"db"`
}

// Client for redis.Client
type Client struct {
	cli *redis.Client
}

// InitClient init client
func InitClient(o *Option) Client {
	opt := redis.Options{
		Addr:     fmt.Sprintf("%s:%d", o.Host, o.Port),
		Password: o.Pass,
		DB:       o.DB,
	}

	return Client{redis.NewClient(&opt)}
}
