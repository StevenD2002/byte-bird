package repository

import (
    "database/sql"

    // "github.com/stevend2002/tgp-bp/internal/db"
    "github.com/stevend2002/tgp-bp/pkg/errors"
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

