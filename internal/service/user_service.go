package service

import (
  "golang.org/x/crypto/bcrypt"

	"byte-bird/internal/repository"
)

type UserService interface {
	CreateUser(name string, email string, password string) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository}
}

func (us *userService) CreateUser(name string, email string, password string) error {
  hashedPassword := hashedPassword(password)

	return us.userRepository.CreateUser(name, email, hashedPassword)
}


func hashedPassword(password string) string {
  bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
  return string(bytes)
}
