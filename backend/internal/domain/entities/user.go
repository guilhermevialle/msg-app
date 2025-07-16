package entities

import (
	"errors"

	nanoid "github.com/matoous/go-nanoid/v2"
)

type User struct {
	Id       string
	Username string
	Email    string
	Password string
	Profile  *Profile
}

func NewUser(username, email, password string) (*User, error) {
	id, err := nanoid.New(21)
	if err != nil {
		return nil, errors.New("failed to generate user id")
	}

	p, err := NewProfile(id)
	if err != nil {
		return nil, err
	}

	return &User{
		Id:       id,
		Username: username,
		Email:    email,
		Password: password,
		Profile:  p,
	}, nil
}

func (u *User) CreatePost(content string) (*Post, error) {
	return NewPost(u.Id, content)
}

func (u *User) CreateComment(postId, content string) (*Comment, error) {
	return NewComment(u.Id, postId, content)
}

func (u *User) CreateReply(postId, parentId, content string) (*Comment, error) {
	return NewReply(u.Id, postId, parentId, content)
}
