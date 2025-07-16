package use_cases

import (
	"app/internal/domain/entities"
	"app/internal/infra/repositories"
)

type GetAllPosts struct {
	postRepo repositories.IPostRepository
}

func NewGetAllPosts(pr repositories.IPostRepository) *GetAllPosts {
	return &GetAllPosts{postRepo: pr}
}

func (uc *GetAllPosts) Execute() []*entities.Post {
	return uc.postRepo.FindAll()
}
