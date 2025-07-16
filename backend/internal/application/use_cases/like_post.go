package use_cases

import (
	"app/backend/internal/infra/repositories"
	"errors"
	"fmt"
)

type LikePost struct {
	PostRepo repositories.IPostRepository
	UserRepo repositories.IUserRepository
}

func NewLikePost(pr repositories.IPostRepository, ur repositories.IUserRepository) *LikePost {
	return &LikePost{
		PostRepo: pr,
		UserRepo: ur,
	}
}

func (lp *LikePost) Execute(userId, postId string) error {
	u := lp.UserRepo.FindById(userId)

	fmt.Println(userId, "userId")

	if u == nil {
		return errors.New("user not found")
	}

	p := lp.PostRepo.FindById(postId)
	if p == nil {
		return errors.New("post not found")
	}

	err := p.AddLike(userId)
	if err != nil {
		return err
	}

	lp.PostRepo.Update(p)
	return nil
}
