package service

import (
  "context"
  "byte-bird/internal/domain/post"
  "byte-bird/internal/repository"
)

// PostService provides methods for post-related operations.
type PostService interface {
  CreatePost(ctx context.Context, newPost *post.Post) error
  GetPosts(ctx context.Context) ([]*post.PostWithUser, error)
  // Add other relevant methods
}

type postService struct {
  postRepository repository.PostRepository
}

// NewPostService creates a new instance of PostService.
func NewPostService(postRepository repository.PostRepository) PostService {
  return &postService{postRepository}
}

func (ps *postService) CreatePost(ctx context.Context, newPost *post.Post) error {
  // Implement any additional logic before creating the post
  return ps.postRepository.CreatePost(ctx, newPost)
}

func (ps *postService) GetPosts(ctx context.Context) ([]*post.PostWithUser, error) {
  // Implement any additional logic before getting the posts
  return ps.postRepository.GetPosts(ctx)
}
