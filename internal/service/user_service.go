package service

import (
	"byte-bird/internal/repository"
)

type UserService interface {
	CreateUser(name string, email string) error
}

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository}
}

func (us *userService) CreateUser(name string, email string) error {
	return us.userRepository.CreateUser(name, email)
}
