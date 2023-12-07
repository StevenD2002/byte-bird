// internal/repository/post_repository.go

package repository

import (
	"byte-bird/internal/domain/post"
	"byte-bird/pkg/errors"
	"context"
	"database/sql"
)

// PostRepository defines methods to interact with post-related data.
type PostRepository interface {
	// Add methods for CRUD operations on posts
	CreatePost(ctx context.Context, newPost *post.Post) error
	// Add other relevant methods
}
type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepository{db}
}

func (r *postRepository) CreatePost(ctx context.Context, newPost *post.Post) error {
	// Add implementation here
	_, err := r.db.Exec("INSERT INTO posts (user_id, content, timestamp) VALUES ($1, $2, $3)", newPost.UserID, newPost.Content, newPost.Timestamp)
	if err != nil {
		return errors.Wrap(err, "failed to create post")
	}



	return nil
}
