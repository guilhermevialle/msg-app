package use_cases

import (
	"app/internal/domain/entities"
	"app/internal/infra/repositories"
	"errors"
)

type CreatePost struct {
	UserRepo repositories.IUserRepository
	PostRepo repositories.IPostRepository
}

func NewCreatePost(ur repositories.IUserRepository, pr repositories.IPostRepository) *CreatePost {
	return &CreatePost{
		UserRepo: ur,
		PostRepo: pr,
	}
}

func (cp *CreatePost) Execute(userId, content string) (*entities.Post, error) {
	u := cp.UserRepo.FindById(userId)
	if u == nil {
		return nil, errors.New("user not found")
	}

	p, err := u.CreatePost(content)
	if err != nil {
		return nil, err
	}

	cp.PostRepo.Save(p)

	return p, nil
}
