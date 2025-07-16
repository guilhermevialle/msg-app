package entities

import (
	"errors"

	nanoid "github.com/matoous/go-nanoid/v2"
)

type Profile struct {
	Id        string
	UserId    string
	Bio       string
	AvatarUrl string
}

func NewProfile(userId string) (*Profile, error) {
	id, err := nanoid.New(21)
	if err != nil {
		return nil, errors.New("failed to generate profile id")
	}

	return &Profile{
		Id:        id,
		UserId:    userId,
		Bio:       "My bio",
		AvatarUrl: "https://cdn.pixabay.com/photo/2023/02/18/11/00/icon-7797704_640.png",
	}, nil
}
