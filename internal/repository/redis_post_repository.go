// internal/repository/redis_post_repository.go

package repository

import (
	"context"
	"encoding/json"
	"fmt"

	"byte-bird/internal/domain/post"
	"github.com/go-redis/redis/v8"
)

// RedisPostRepository is a Redis-based repository for posts
type RedisPostRepository struct {
	client *redis.Client
}

// NewRedisPostRepository creates a new Redis-based repository for posts
func NewRedisPostRepository(redisAddress string) *RedisPostRepository {
	return &RedisPostRepository{
		client: redis.NewClient(&redis.Options{
			Addr: redisAddress,
		}),
	}
}

// GetPostByID retrieves a post by ID from the Redis database
func (r *RedisPostRepository) GetPostByID(ctx context.Context, postID string) (*post.Post, error) {
	// Implement logic to retrieve a post from Redis
	return nil, nil
}

// CreatePost inserts a new post into the Redis database
func (r *RedisPostRepository) CreatePost(ctx context.Context, newPost *post.Post) error {
	key := fmt.Sprintf("post:%s", newPost.ID)
	postJSON, err := json.Marshal(newPost)
	if err != nil {
		return err
	}

	// Use the context for handling timeouts or cancellations
	return r.client.WithContext(ctx).Set(ctx, key, postJSON, 0).Err()
}

