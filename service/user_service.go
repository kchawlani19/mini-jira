package service

import (
	"errors"

	"golang.org/x/crypto/bcrypt"

	"mini-jira/model"
	"mini-jira/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(name, email, password string) (model.User, error) {

	if name == "" || email == "" || password == "" {
		return model.User{}, errors.New("missing fields")
	}

	hashedPwd, _ := bcrypt.GenerateFromPassword(
		[]byte(password), bcrypt.DefaultCost,
	)

	user := model.User{
		Name:     name,
		Email:    email,
		Password: string(hashedPwd),
	}

	return s.repo.Save(user), nil
}

func (s *UserService) GetAllUsers() []model.User {
	return s.repo.FindAll()
}

func (s *UserService) GetUserByID(id int) (model.User, error) {

	user, found := s.repo.FindByID(id)
	if !found {
		return model.User{}, errors.New("user not found")
	}
	return user, nil
}

func (s *UserService) UpdateUser(id int, name, email string) (model.User, error) {

	if name == "" || email == "" {
		return model.User{}, errors.New("missing fields")
	}

	user, found := s.repo.FindByID(id)
	if !found {
		return model.User{}, errors.New("user not found")
	}

	user.Name = name
	user.Email = email

	return s.repo.Update(user), nil
}

func (s *UserService) DeleteUser(id int) error {

	_, found := s.repo.FindByID(id)
	if !found {
		return errors.New("user not found")
	}

	s.repo.Delete(id)
	return nil
}

func (s *UserService) GetUserByEmail(email string) (model.User, error) {

	users := s.repo.FindAll()
	for _, u := range users {
		if u.Email == email {
			return u, nil
		}
	}

	return model.User{}, errors.New("user not found")
}
