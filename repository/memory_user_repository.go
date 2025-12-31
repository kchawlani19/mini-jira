package repository

import "mini-jira/model"

type MemoryUserRepository struct {
	users  map[int]model.User
	nextID int
}

func NewMemoryUserRepository() *MemoryUserRepository {
	return &MemoryUserRepository{
		users:  make(map[int]model.User),
		nextID: 1,
	}
}

func (r *MemoryUserRepository) Save(user model.User) model.User {
	user.ID = r.nextID
	r.users[user.ID] = user
	r.nextID++
	return user
}

func (r *MemoryUserRepository) FindAll() []model.User {
	users := make([]model.User, 0)

	for _, user := range r.users {
		users = append(users, user)
	}

	return users
}

func (r *MemoryUserRepository) FindByID(id int) (model.User, bool) {
	user, found := r.users[id]
	return user, found
}

func (r *MemoryUserRepository) Update(user model.User) model.User {
	r.users[user.ID] = user
	return user
}

func (r *MemoryUserRepository) Delete(id int) {
	delete(r.users, id)
}
