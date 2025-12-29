package repository

import "mini-jira/model"

type UserRepository struct {
	users  map[int]model.User
	nextID int
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users:  make(map[int]model.User),
		nextID: 1,
	}
}

func (r *UserRepository) Save(user model.User) model.User {
	user.ID = r.nextID
	r.users[user.ID] = user
	r.nextID++
	return user
}

func (r *UserRepository) FindAll() []model.User {

	users := make([]model.User, 0)

	for _, user := range r.users {
		users = append(users, user)
	}

	return users
}

func (r *UserRepository) FindByID(id int) (model.User, bool) {

	user, found := r.users[id]
	return user, found
}
