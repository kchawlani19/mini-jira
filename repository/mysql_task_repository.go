package repository

import (
	"database/sql"
	"errors"
	"mini-jira/model"
)

type MySQLTaskRepository struct {
	db *sql.DB
}

func NewMySQLTaskRepository(db *sql.DB) TaskRepository {
	return &MySQLTaskRepository{db: db}
}

func (r *MySQLTaskRepository) Save(task model.Task) (model.Task, error) {
	query := `
		INSERT INTO tasks (title,description,status,assignee_id)
		VALUES( ?,?,?,?)
		`

	result, err := r.db.Exec(
		query,
		task.Title,
		task.Description,
		task.Status,
		task.AssigneeID,
	)

	if err != nil {
		return task, err
	}

	id, err := result.LastInsertId()
	if err != nil {
		return task, err
	}
	task.ID = int(id)
	return task, nil
}

func (r *MySQLTaskRepository) FindAll() ([]model.Task, error) {

	query := `
		SELECT id, title, description, status, assignee_id
		FROM tasks
	`

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task

	for rows.Next() {
		var task model.Task
		var assigneeID sql.NullInt64

		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Status,
			&assigneeID,
		)
		if err != nil {
			return nil, err
		}

		if assigneeID.Valid {
			id := int(assigneeID.Int64)
			task.AssigneeID = &id
		} else {
			task.AssigneeID = nil
		}

		tasks = append(tasks, task)
	}

	return tasks, nil
}

func (r *MySQLTaskRepository) FindByID(id int) (model.Task, error) {

	query := `
		SELECT id, title, description, status, assignee_id
		FROM tasks
		WHERE id = ?
	`

	var task model.Task
	var assigneeID sql.NullInt64

	err := r.db.QueryRow(query, id).Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Status,
		&assigneeID,
	)

	if err == sql.ErrNoRows {
		return task, errors.New("task not found")
	}
	if err != nil {
		return task, err
	}

	if assigneeID.Valid {
		id := int(assigneeID.Int64)
		task.AssigneeID = &id
	}

	return task, nil
}

func (r *MySQLTaskRepository) UpdateAssignee(taskID int, assigneeID *int) error {

	query := `
		UPDATE tasks
		SET assignee_id = ?
		WHERE id = ?
	`

	_, err := r.db.Exec(query, assigneeID, taskID)
	return err
}

func (r *MySQLTaskRepository) UpdateStatus(taskID int, status model.TaskStatus) error {

	query := `
		UPDATE tasks
		SET status = ?
		WHERE id = ?
	`

	_, err := r.db.Exec(query, status, taskID)
	return err
}

func (r *MySQLTaskRepository) Delete(id int) error {

	query := `
		DELETE FROM tasks
		WHERE id = ?
	`

	_, err := r.db.Exec(query, id)
	return err
}
