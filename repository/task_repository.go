package repository

import (
	"database/sql"
	"test_mini_jira/models"
)

type TaskRepository struct {
	DB *sql.DB
}

func (r *TaskRepository) Create(t models.Task) error {
	_, err := r.DB.Exec(
		"INSERT INTO tasks(title,description,status,assignee_id) VALUES (?,?,?,?)",
		t.Title, t.Description, t.Status, t.AssigneeID,
	)
	return err
}

// 👑 ADMIN → all NON-DELETED tasks
func (r *TaskRepository) GetAll(limit, offset int) ([]models.Task, error) {
	rows, err := r.DB.Query(
		`SELECT id,title,description,status,assignee_id
		 FROM tasks
		 WHERE deleted_at IS NULL
		 LIMIT ? OFFSET ?`,
		limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var t models.Task
		rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.AssigneeID)
		tasks = append(tasks, t)
	}
	return tasks, nil
}

// 🔐 USER → own NON-DELETED tasks
func (r *TaskRepository) GetByUser(userID, limit, offset int) ([]models.Task, error) {
	rows, err := r.DB.Query(
		`SELECT id,title,description,status,assignee_id
		 FROM tasks
		 WHERE assignee_id=? AND deleted_at IS NULL
		 LIMIT ? OFFSET ?`,
		userID, limit, offset,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []models.Task
	for rows.Next() {
		var t models.Task
		rows.Scan(&t.ID, &t.Title, &t.Description, &t.Status, &t.AssigneeID)
		tasks = append(tasks, t)
	}
	return tasks, nil
}

func (r *TaskRepository) Update(id int, t models.Task) error {
	_, err := r.DB.Exec(
		`UPDATE tasks
		 SET title=?,description=?,status=?,assignee_id=?
		 WHERE id=? AND deleted_at IS NULL`,
		t.Title, t.Description, t.Status, t.AssigneeID, id,
	)
	return err
}

func (r *TaskRepository) PatchStatus(id int, status string) error {
	_, err := r.DB.Exec(
		`UPDATE tasks
		 SET status=?
		 WHERE id=? AND deleted_at IS NULL`,
		status, id,
	)
	return err
}

// 🗑️ SOFT DELETE
func (r *TaskRepository) Delete(id int) error {
	res, err := r.DB.Exec(
		"UPDATE tasks SET deleted_at=NOW() WHERE id=? AND deleted_at IS NULL",
		id,
	)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}
	return nil
}

func (r *TaskRepository) IsOwner(taskID int, userID int) (bool, error) {
	var assigneeID *int
	err := r.DB.QueryRow(
		`SELECT assignee_id
		 FROM tasks
		 WHERE id=? AND deleted_at IS NULL`,
		taskID,
	).Scan(&assigneeID)

	if err != nil {
		return false, err
	}
	if assigneeID == nil {
		return false, nil
	}
	return *assigneeID == userID, nil
}
