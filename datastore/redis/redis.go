package redis

import (
	"github.com/go-redis/redis"
	"github.com/opentracing/opentracing-go"
)

// Configuration provides configuration for redis Client
type Configuration struct {
	Address  string `required:"true"`
	Password string `required:"true" split_words:"true"`
	Database int    `required:"true" split_words:"true"`
}

// CreateRedisCLient creates connection client for redis server
func CreateRedisCLient(configuration Configuration, tracer opentracing.Tracer) *redis.Client {

	client := redis.NewClient(&redis.Options{
		Addr:     configuration.Address,
		Password: configuration.Password,
		DB:       configuration.Database,
	})

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	return client
}
