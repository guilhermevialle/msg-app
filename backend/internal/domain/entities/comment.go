package entities

import (
	"errors"
	"time"

	nanoid "github.com/matoous/go-nanoid/v2"
)

type Comment struct {
	Id        string
	UserId    string
	PostId    string
	ParentId  *string
	Content   string
	CreatedAt time.Time
	Replies   []*Comment
}

func NewComment(userId, postId, content string) (*Comment, error) {
	id, err := nanoid.New(21)
	if err != nil {
		return nil, errors.New("failed to generate comment id")
	}

	return &Comment{
		Id:        id,
		UserId:    userId,
		PostId:    postId,
		Content:   content,
		CreatedAt: time.Now(),
		Replies:   make([]*Comment, 0),
	}, nil
}

func NewReply(userId, postId, parentId, content string) (*Comment, error) {
	id, err := nanoid.New(21)
	if err != nil {
		return nil, errors.New("failed to generate reply id")
	}

	return &Comment{
		Id:        id,
		UserId:    userId,
		PostId:    postId,
		ParentId:  &parentId,
		Content:   content,
		CreatedAt: time.Now(),
		Replies:   make([]*Comment, 0),
	}, nil
}
