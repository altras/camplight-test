package application

import (
	"backend/internal/domain"
	"errors"
	"regexp"
)

var (
	ErrInvalidEmail = errors.New("invalid email address")
	ErrInvalidPhone = errors.New("invalid phone number")
)

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) ListUsers(page, limit int) ([]domain.User, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	return s.repo.FindAll(page, limit)
}

func (s *UserService) CreateUser(user *domain.User) error {
	if err := validateUser(user); err != nil {
		return err
	}
	return s.repo.Create(user)
}

func (s *UserService) DeleteUser(id string) error {
	if id == "" {
		return errors.New("invalid user ID")
	}
	return s.repo.Delete(id)
}

func validateUser(user *domain.User) error {
	if user.Name == "" {
		return errors.New("name is required")
	}
	if !isValidEmail(user.Email) {
		return ErrInvalidEmail
	}
	if !isValidPhone(user.Phone) {
		return ErrInvalidPhone
	}
	return nil
}

func isValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
	return emailRegex.MatchString(email)
}

func isValidPhone(phone string) bool {
	phoneRegex := regexp.MustCompile(`^[+]?[0-9]{10,14}$`)
	return phoneRegex.MatchString(phone)
}

func (s *UserService) SearchUsers(query string, page, limit int) ([]domain.User, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 || limit > 100 {
		limit = 10
	}
	return s.repo.Search(query, page, limit)
}