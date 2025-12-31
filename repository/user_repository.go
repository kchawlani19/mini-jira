package repository

import "mini-jira/model"

type UserRepository interface {
	Save(user model.User) model.User
	FindAll() []model.User
	FindByID(id int) (model.User, bool)
	Update(user model.User) model.User
	Delete(id int)
}
