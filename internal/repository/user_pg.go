package repository

import (
	"database/sql"
	"user-service/internal/domain"
)

type UserRepository interface {
	Create(user *domain.User) error
	GetByID(id int64) (*domain.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Create(user *domain.User) error {
	err := r.db.QueryRow("INSERT INTO users (name, email) VALUES ($1, $2) RETURNING id",
		user.Name, user.Email).Scan(&user.ID)
	return err
}

func (r *userRepository) GetByID(id int64) (*domain.User, error) {
	user := &domain.User{}
	row := r.db.QueryRow("SELECT id, name, email FROM users WHERE id = $1", id)
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	return user, err
}
