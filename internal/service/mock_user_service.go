// internal/service/mock_user_service.go

package service

import "fmt"

type MockUserService struct {
	// will do stuff here later
}

func NewMockUserService() *MockUserService {
	return &MockUserService{}
}

func (m *MockUserService) CreateUser(name, email string) error {
	fmt.Printf("[Mock] Creating user: Name - %s, Email - %s\n", name, email)
	return nil
}

