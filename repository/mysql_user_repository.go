package repository

import (
	"database/sql"

	"mini-jira/model"
)

type MySQLUserRepository struct {
	db *sql.DB
}

func NewMySQLUserRepository(db *sql.DB) *MySQLUserRepository {
	return &MySQLUserRepository{db: db}
}

func (r *MySQLUserRepository) Save(user model.User) model.User {

	result, _ := r.db.Exec(
		"INSERT INTO users(name, email, password) VALUES (?, ?, ?)",
		user.Name, user.Email, user.Password,
	)

	id, _ := result.LastInsertId()
	user.ID = int(id)
	return user
}

func (r *MySQLUserRepository) FindAll() []model.User {

	rows, _ := r.db.Query("SELECT id, name, email, password FROM users")
	defer rows.Close()

	users := []model.User{}
	for rows.Next() {
		var u model.User
		rows.Scan(&u.ID, &u.Name, &u.Email, &u.Password)
		users = append(users, u)
	}

	return users
}

func (r *MySQLUserRepository) FindByID(id int) (model.User, bool) {

	row := r.db.QueryRow(
		"SELECT id, name, email, password FROM users WHERE id=?",
		id,
	)

	var u model.User
	err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password)
	if err != nil {
		return model.User{}, false
	}

	return u, true
}

func (r *MySQLUserRepository) Update(user model.User) model.User {

	r.db.Exec(
		"UPDATE users SET name=?, email=? WHERE id=?",
		user.Name, user.Email, user.ID,
	)

	return user
}

func (r *MySQLUserRepository) Delete(id int) {
	r.db.Exec("DELETE FROM users WHERE id=?", id)
}
