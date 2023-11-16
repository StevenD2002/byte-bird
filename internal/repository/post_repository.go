// internal/repository/post_repository.go

package repository

import (
	"context"
	"byte-bird/internal/domain/post"
)

// PostRepository defines methods to interact with post-related data.
type PostRepository interface {
	// Add methods for CRUD operations on posts
	GetPostByID(ctx context.Context, postID string) (*post.Post, error)
	CreatePost(ctx context.Context, newPost *post.Post) error
	// Add other relevant methods
}

