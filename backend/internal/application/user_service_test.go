package application

import (
	"backend/internal/domain"
	"errors"
	"testing"
)

// MockUserRepository is a mock implementation of domain.UserRepository
type MockUserRepository struct {
	users []domain.User
}

func (m *MockUserRepository) FindAll(page, limit int) ([]domain.User, error) {
	start := (page - 1) * limit
	end := start + limit
	if start >= len(m.users) {
		return []domain.User{}, nil
	}
	if end > len(m.users) {
		end = len(m.users)
	}
	return m.users[start:end], nil
}

func (m *MockUserRepository) Create(user *domain.User) error {
	m.users = append(m.users, *user)
	return nil
}

func (m *MockUserRepository) Delete(id string) error {
	for i, user := range m.users {
		if user.ID == id {
			m.users = append(m.users[:i], m.users[i+1:]...)
			return nil
		}
	}
	return errors.New("user not found")
}

func TestUserService_ListUsers(t *testing.T) {
	mockRepo := &MockUserRepository{
		users: []domain.User{
			{ID: "1", Name: "Alice", Email: "alice@example.com", Phone: "1234567890"},
			{ID: "2", Name: "Bob", Email: "bob@example.com", Phone: "0987654321"},
		},
	}

	service := NewUserService(mockRepo)

	tests := []struct {
		name     string
		page     int
		limit    int
		expected int
	}{
		{"First page", 1, 1, 1},
		{"Second page", 2, 1, 1},
		{"All users", 1, 10, 2},
		{"Empty page", 3, 1, 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			users, err := service.ListUsers(tt.page, tt.limit)
			if err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
			if len(users) != tt.expected {
				t.Errorf("Expected %d users, got %d", tt.expected, len(users))
			}
		})
	}
}

func TestUserService_CreateUser(t *testing.T) {
	mockRepo := &MockUserRepository{}
	service := NewUserService(mockRepo)

	tests := []struct {
		name        string
		user        domain.User
		expectedErr error
	}{
		{
			name: "Valid user",
			user: domain.User{Name: "Alice", Email: "alice@example.com", Phone: "1234567890"},
			expectedErr: nil,
		},
		{
			name: "Invalid email",
			user: domain.User{Name: "Bob", Email: "invalid-email", Phone: "0987654321"},
			expectedErr: ErrInvalidEmail,
		},
		{
			name: "Invalid phone",
			user: domain.User{Name: "Charlie", Email: "charlie@example.com", Phone: "123"},
			expectedErr: ErrInvalidPhone,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.CreateUser(&tt.user)
			if err != tt.expectedErr {
				t.Errorf("Expected error %v, got %v", tt.expectedErr, err)
			}
		})
	}
}

func TestUserService_DeleteUser(t *testing.T) {
	mockRepo := &MockUserRepository{
		users: []domain.User{
			{ID: "1", Name: "Alice", Email: "alice@example.com", Phone: "1234567890"},
		},
	}
	service := NewUserService(mockRepo)

	tests := []struct {
		name        string
		id          string
		expectedErr bool
	}{
		{"Existing user", "1", false},
		{"Non-existing user", "2", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := service.DeleteUser(tt.id)
			if (err != nil) != tt.expectedErr {
				t.Errorf("Expected error: %v, got: %v", tt.expectedErr, err)
			}
		})
	}
}