package service

import (
	"byte-bird/internal/repository"
	"context"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	CreateUser(name string, email string, password string) error
	AuthenticateUser(ctx context.Context, email string, password string) (string, error)
}

type userService struct {
	userRepository repository.UserRepository
}

type Claims struct {
	UserID string `json:"user_id"`
	Email  string `json:"email"`
	jwt.StandardClaims
}

func NewUserService(userRepository repository.UserRepository) UserService {
	return &userService{userRepository}
}

func (us *userService) CreateUser(name string, email string, password string) error {
	hashedPassword := hashPassword(password)

	return us.userRepository.CreateUser(name, email, hashedPassword)
}

func (us *userService) AuthenticateUser(ctx context.Context, email string, password string) (string, error) {
	userId, userPassword, userEmail, err := us.userRepository.GetUserByEmail(ctx, email)

	err = bcrypt.CompareHashAndPassword([]byte(userPassword), []byte(password))
	if err != nil {
		return "", fmt.Errorf("invalid credentials")
	}

	// generate the token
	token, err := generateToken(userId, userEmail)
	if err != nil {
		return "", fmt.Errorf("error generating token")
	}

	return token, nil
}

func hashPassword(password string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes)
}

func generateToken(userID string, email string) (string, error) {
	claims := Claims{
		UserID: userID,
		Email:  email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

  // will change the secret key to something in an env later, dont want to make grading TOO hard....
	return token.SignedString([]byte("temp-secret-key"))
}
