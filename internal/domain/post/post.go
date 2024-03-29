package post

import (
	"time"
)

// Post represents a post in the application.
type Post struct {
	ID        string
	UserID    string
	Content   string
	Timestamp time.Time
}

// this type is used to query posts and have access to the username
type PostWithUser struct {
  ID        string
  UserID    string
  Authorname  string
  Content   string
  Timestamp time.Time
}
