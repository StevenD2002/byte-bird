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
	GetPosts(ctx context.Context) ([]*post.PostWithUser, error)
}
type postRepository struct {
	db *sql.DB
}

func NewPostRepository(db *sql.DB) PostRepository {
	return &postRepository{db}
}

func (r *postRepository) CreatePost(ctx context.Context, newPost *post.Post) error {
	_, err := r.db.Exec("INSERT INTO posts (user_id, content, timestamp) VALUES ($1, $2, $3)", newPost.UserID, newPost.Content, newPost.Timestamp)
	if err != nil {
		return errors.Wrap(err, "failed to create post")
	}

	return nil
}

func (r *postRepository) GetPosts(ctx context.Context) ([]*post.PostWithUser, error) {
	rows, err := r.db.Query("SELECT posts.id, posts.user_id, users.name, posts.content, posts.timestamp FROM posts INNER JOIN users ON posts.user_id = users.id")
	if err != nil {
		return nil, errors.Wrap(err, "failed to get posts")
	}
	defer rows.Close()

	var posts []*post.PostWithUser
	for rows.Next() {
		var post post.PostWithUser
		if err := rows.Scan(&post.ID, &post.UserID, &post.Authorname, &post.Content, &post.Timestamp); err != nil {
			return nil, errors.Wrap(err, "failed to scan row into PostWithUser struct")
		}
		posts = append(posts, &post)
	}

	if err := rows.Err(); err != nil {
		return nil, errors.Wrap(err, "failed to iterate over rows")
	}

	return posts, nil
}
