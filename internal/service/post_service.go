// internal/service/post_service.go

package service

import (
	"context"
	"byte-bird/internal/domain/post"
)

// PostService provides methods for post-related operations.
type PostService interface {
	// Add methods for post-related business logic
	GetPostByID(ctx context.Context, postID string) (*post.Post, error)
	CreatePost(ctx context.Context, newPost *post.Post) error
	// Add other relevant methods
}

