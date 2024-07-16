package domain

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)


type User struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Password  string    `json:"-"` // Password is not exposed in JSON
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *User) SetPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}

type UserRepository interface {
	FindAll(page, limit int) ([]User, error)
	FindByID(id string) (*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id string) error
	Search(query string, page, limit int) ([]User, error)
}