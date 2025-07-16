package entities

import (
	"errors"

	nanoid "github.com/matoous/go-nanoid/v2"
)

type UserLikePost struct {
	Id     string
	UserId string
	PostId string
}

func NewUserLikePost(userId, postId string) (*UserLikePost, error) {
	id, err := nanoid.New(21)
	if err != nil {
		return nil, errors.New("failed to generate user_like_post id")
	}

	return &UserLikePost{
		Id:     id,
		UserId: userId,
		PostId: postId,
	}, nil
}
