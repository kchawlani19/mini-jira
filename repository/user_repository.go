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

// 👑 ADMIN-ONLY
func (r *UserRepository) GetAll() ([]models.User, error) {
	rows, err := r.DB.Query(
		"SELECT id,email,role FROM users",
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		rows.Scan(&u.ID, &u.Email, &u.Role)
		users = append(users, u)
	}
	return users, nil
}
