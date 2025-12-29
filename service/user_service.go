package service

import (
	"errors"
	"mini-jira/model"
	"mini-jira/repository"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) GetAllUsers() []model.User {
	return s.repo.FindAll()
}

func (s *UserService) CreateUser(name, email string) (model.User, error) {
	if name == "" || email == "" {
		return model.User{}, errors.New("name or mail missing")
	}

	user := model.User{
		Name:  name,
		Email: email,
	}
	return s.repo.Save(user), nil
}

func (s *UserService) GetUserByID(id int) (model.User, error) {

	user, found := s.repo.FindByID(id)
	if !found {
		return model.User{}, errors.New("user not found")
	}

	return user, nil
}
