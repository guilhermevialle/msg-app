package repositories

import "app/backend/internal/domain/entities"

type IUserRepository interface {
	Save(user *entities.User)
	FindByUsername(username string) *entities.User
	FindById(id string) *entities.User
	FindByEmail(email string) *entities.User
	FindAll() []*entities.User
}

type UserRepository struct {
	users []*entities.User
}

var _ IUserRepository = (*UserRepository)(nil)

func NewUserRepository() *UserRepository {
	return &UserRepository{
		users: make([]*entities.User, 0),
	}
}

func (r *UserRepository) Save(user *entities.User) {
	r.users = append(r.users, user)
}

func (r *UserRepository) FindByUsername(username string) *entities.User {
	for _, u := range r.users {
		if u.Username == username {
			return u
		}
	}
	return nil
}

func (r *UserRepository) FindById(id string) *entities.User {
	for _, u := range r.users {
		if u.Id == id {
			return u
		}
	}
	return nil
}

func (r *UserRepository) FindByEmail(email string) *entities.User {
	for _, u := range r.users {
		if u.Email == email {
			return u
		}
	}
	return nil
}

func (r *UserRepository) FindAll() []*entities.User {
	return r.users
}
