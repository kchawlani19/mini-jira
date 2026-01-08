package repository

import (
	"database/sql"
	"test_mini_jira/models"
)

type UserRepository struct {
	DB *sql.DB
}

func (r *UserRepository) Create(u models.User) error {
	_, err := r.DB.Exec(
		"INSERT INTO users(email,password,role) VALUES (?,?,?)",
		u.Email, u.Password, u.Role,
	)
	return err
}

func (r *UserRepository) FindByEmail(email string) (models.User, error) {
	var u models.User
	err := r.DB.QueryRow(
		"SELECT id,email,password,role FROM users WHERE email=?",
		email,
	).Scan(&u.ID, &u.Email, &u.Password, &u.Role)
	return u, err
}
