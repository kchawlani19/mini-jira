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

func (r *TaskRepository) GetAll() ([]models.Task, error) {
	rows, err := r.DB.Query(
		"SELECT id,title,description,status,assignee_id FROM tasks",
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
		"UPDATE tasks SET title=?,description=?,status=?,assignee_id=? WHERE id=?",
		t.Title, t.Description, t.Status, t.AssigneeID, id,
	)
	return err
}

func (r *TaskRepository) PatchStatus(id int, status string) error {
	_, err := r.DB.Exec(
		"UPDATE tasks SET status=? WHERE id=?", status, id,
	)
	return err
}

func (r *TaskRepository) Delete(id int) error {
	_, err := r.DB.Exec("DELETE FROM tasks WHERE id=?", id)
	return err
}

// 🔐 OWNERSHIP CHECK
func (r *TaskRepository) IsOwner(taskID int, userID int) (bool, error) {
	var assigneeID *int
	err := r.DB.QueryRow(
		"SELECT assignee_id FROM tasks WHERE id=?",
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
