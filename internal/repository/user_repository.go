package repository

import (
	"context"
	"database/sql"

	// "byte-bird/internal/db"
	"byte-bird/pkg/errors"
)

type UserRepository interface {
	CreateUser(name string, email string, hashedPassword string) error
  GetUserByEmail(ctx context.Context, email string) (string, string, string, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) CreateUser(name string, email string, hashedPassword string) error {
	_, err := ur.db.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", name, email, hashedPassword)
	if err != nil {
		return errors.Wrap(err, "failed to create user")
	}
	return nil
}

func (ur *userRepository) AuthenticateUser(ctx context.Context, email string, password string) (string, error) {
	var userID string
	var hashedPassword string
	err := ur.db.QueryRow("SELECT id, password FROM users WHERE email = $1", email).Scan(&userID, &hashedPassword)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.Wrap(err, "user not found")
		}
		return "", errors.Wrap(err, "failed to authenticate user")
	}

	return userID, nil
}

// get the user by email and return the userID, password, and the email
func (ur *userRepository) GetUserByEmail(ctx context.Context, email string) (string, string, string, error) {
  var userID string
  var hashedPassword string
  var userEmail string
  err := ur.db.QueryRow("SELECT id, password, email FROM users WHERE email = $1", email).Scan(&userID, &hashedPassword, &userEmail)
  if err != nil {
    if err == sql.ErrNoRows {
      return "", "", "", errors.Wrap(err, "user not found")
    }
    return "", "", "", errors.Wrap(err, "failed to get user by email")
  }
  return userID, hashedPassword, userEmail, nil
}
