package repository

import (
	"database/sql"

	// "byte-bird/internal/db"
	"byte-bird/pkg/errors"
)

type UserRepository interface {
	CreateUser(name string, email string) error
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (ur *userRepository) CreateUser(name string, email string) error {
	_, err := ur.db.Exec("INSERT INTO users (name, email) VALUES ($1, $2)", name, email)
	if err != nil {
		return errors.Wrap(err, "failed to create user")
	}
	return nil
}

func (ur *userRepository) GetUser(id int) (string, string, error) {
  var name, email string
  err := ur.db.QueryRow("SELECT name, email FROM users WHERE id = $1", id).Scan(&name, &email)
  if err != nil {
    return "", "", errors.Wrap(err, "failed to get user")
  }
  return name, email, nil
}


