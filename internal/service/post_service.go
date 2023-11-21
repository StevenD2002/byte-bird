// internal/service/post_service.go

package service

import (
	"byte-bird/internal/domain/post"
	"byte-bird/internal/repository"
	"context"
)

// PostService provides methods for post-related operations.
type PostService interface {
	CreatePost(ctx context.Context, newPost *post.Post) error
}

// PostServiceImpl is an implementation of the PostService interface.
type PostServiceImpl struct {
	postRepo repository.PostRepository
}

// NewPostServiceImpl creates a new instance of PostServiceImpl.
func NewPostServiceImpl(postRepo repository.PostRepository) *PostServiceImpl {
	return &PostServiceImpl{
		postRepo: postRepo,
	}
}

func (s *PostServiceImpl) CreatePost(ctx context.Context, newPost *post.Post) error {
	// Call the CreatePost method in the repository
	return s.postRepo.CreatePost(ctx, newPost)
}
