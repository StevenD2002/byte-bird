// internal/domain/post/post.go

package post

import (
	"time"
)

// Post represents a post in the application.
type Post struct {
	ID        string
	AuthorID  string
	Content   string
	Timestamp time.Time
	// Add other relevant fields
}

