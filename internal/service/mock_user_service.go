// internal/service/mock_user_service.go

package service

import "fmt"

type MockUserService struct {
	// You can add fields or methods needed for testing purposes
}

func NewMockUserService() *MockUserService {
	return &MockUserService{}
}

func (m *MockUserService) CreateUser(name, email string) error {
	// Implement mock logic for CreateUser (e.g., print details for testing)
	fmt.Printf("[Mock] Creating user: Name - %s, Email - %s\n", name, email)
	return nil
}

