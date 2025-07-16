package use_cases

import (
	"app/backend/internal/domain/entities"
	"app/backend/internal/infra/repositories"
	"errors"
)

type AddCommentToPost struct {
	userRepo repositories.IUserRepository
	postRepo repositories.IPostRepository
}

func NewAddCommentToPost(
	ur repositories.IUserRepository,
	pr repositories.IPostRepository,
) *AddCommentToPost {
	return &AddCommentToPost{
		userRepo: ur,
		postRepo: pr,
	}
}

func (ap *AddCommentToPost) Execute(userId, postId, content string) (*entities.Comment, error) {
	u := ap.userRepo.FindById(userId)
	if u == nil {
		return nil, errors.New("user not found")
	}

	p := ap.postRepo.FindById(postId)
	if p == nil {
		return nil, errors.New("post not found")
	}

	c, err := u.CreateComment(postId, content)
	if err != nil {
		return nil, err
	}

	p.Comments = append(p.Comments, c)
	ap.postRepo.Update(p)

	return c, nil
}
