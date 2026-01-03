package repository

import "mini-jira/model"

type TaskRepository interface {
	Save(task model.Task) (model.Task, error)

	FindAll() ([]model.Task, error)
	FindByID(id int) (model.Task, error)

	UpdateAssignee(taskID int, assigneeId *int) error
	UpdateStatus(taskID int, status model.TaskStatus) error

	Delete(id int) error
}
