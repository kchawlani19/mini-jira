package service

import (
	"errors"

	"mini-jira/model"
	"mini-jira/repository"
)

type TaskService struct {
	taskRepo repository.TaskRepository
	userRepo repository.UserRepository
}

func NewTaskService(
	taskRepo repository.TaskRepository,
	userRepo repository.UserRepository,
) *TaskService {
	return &TaskService{
		taskRepo: taskRepo,
		userRepo: userRepo,
	}
}

// CREATE TASK
func (s *TaskService) CreateTask(title string, description string) (model.Task, error) {

	if title == "" {
		return model.Task{}, errors.New("title is required")
	}

	task := model.Task{
		Title:       title,
		Description: description,
		Status:      model.StatusOpen, // DEFAULT STATUS
		AssigneeID:  nil,              // unassigned
	}

	return s.taskRepo.Save(task)
}

// GET ALL TASKS
func (s *TaskService) GetAllTasks() ([]model.Task, error) {
	return s.taskRepo.FindAll()
}

// GET TASK BY ID
func (s *TaskService) GetTaskByID(id int) (model.Task, error) {
	return s.taskRepo.FindByID(id)
}

// ASSIGN TASK TO USER
func (s *TaskService) AssignTask(taskID int, assigneeID int) error {

	// 1️⃣ check task exists
	_, err := s.taskRepo.FindByID(taskID)
	if err != nil {
		return err
	}

	// 2️⃣ check user exists (UserRepository returns bool)
	_, found := s.userRepo.FindByID(assigneeID)
	if !found {
		return errors.New("assignee user not found")
	}

	return s.taskRepo.UpdateAssignee(taskID, &assigneeID)
}

// UPDATE TASK STATUS (WITH RULES)
func (s *TaskService) UpdateTaskStatus(taskID int, newStatus model.TaskStatus) error {

	task, err := s.taskRepo.FindByID(taskID)
	if err != nil {
		return err
	}

	if !isValidStatusTransition(task.Status, newStatus) {
		return errors.New("invalid status transition")
	}

	return s.taskRepo.UpdateStatus(taskID, newStatus)
}

// DELETE TASK
func (s *TaskService) DeleteTask(taskID int) error {

	// ensure task exists
	_, err := s.taskRepo.FindByID(taskID)
	if err != nil {
		return err
	}

	return s.taskRepo.Delete(taskID)
}

// PRIVATE HELPER — STATUS RULES
func isValidStatusTransition(oldStatus, newStatus model.TaskStatus) bool {

	switch oldStatus {
	case model.StatusOpen:
		return newStatus == model.StatusInProgress

	case model.StatusInProgress:
		return newStatus == model.StatusDone

	case model.StatusDone:
		return false
	}

	return false
}
