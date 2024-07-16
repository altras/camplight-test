package postgres

import (
	"backend/internal/domain"
	"database/sql"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) FindAll(page, limit int) ([]domain.User, error) {
	offset := (page - 1) * limit
	rows, err := r.db.Query("SELECT id, name, email, phone, created_at, updated_at FROM users ORDER BY created_at DESC LIMIT $1 OFFSET $2", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var u domain.User
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Phone, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *UserRepository) FindByID(id string) (*domain.User, error) {
	var u domain.User
	err := r.db.QueryRow("SELECT id, name, email, phone, created_at, updated_at FROM users WHERE id = $1", id).
		Scan(&u.ID, &u.Name, &u.Email, &u.Phone, &u.CreatedAt, &u.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) Create(user *domain.User) error {
	_, err := r.db.Exec("INSERT INTO users (id, name, email, phone, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)",
		user.ID, user.Name, user.Email, user.Phone, user.CreatedAt, user.UpdatedAt)
	return err
}

func (r *UserRepository) Update(user *domain.User) error {
	_, err := r.db.Exec("UPDATE users SET name = $1, email = $2, phone = $3, updated_at = $4 WHERE id = $5",
		user.Name, user.Email, user.Phone, user.UpdatedAt, user.ID)
	return err
}

func (r *UserRepository) Delete(id string) error {
	_, err := r.db.Exec("DELETE FROM users WHERE id = $1", id)
	return err
}

func (r *UserRepository) Search(query string, page, limit int) ([]domain.User, error) {
	offset := (page - 1) * limit
	rows, err := r.db.Query(`
		SELECT id, name, email, phone, created_at, updated_at 
		FROM users 
		WHERE name ILIKE $1 OR email ILIKE $1
		ORDER BY created_at DESC 
		LIMIT $2 OFFSET $3`,
		"%"+query+"%", limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []domain.User
	for rows.Next() {
		var u domain.User
		err := rows.Scan(&u.ID, &u.Name, &u.Email, &u.Phone, &u.CreatedAt, &u.UpdatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}