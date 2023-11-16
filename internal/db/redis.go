// internal/db/redis.go

package db

import (

	"github.com/go-redis/redis/v8"
	// "context"
)

// RedisClient is a global Redis client.
var RedisClient *redis.Client

// InitRedis initializes the Redis client.
func InitRedis() {
	RedisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		// Add other Redis configuration options
	})
}

// CloseRedis closes the Redis client.
func CloseRedis() {
	_ = RedisClient.Close()
}

