package entities

import (
	"errors"
	"time"

	nanoid "github.com/matoous/go-nanoid/v2"
)

type Post struct {
	Id        string
	UserId    string
	Content   string
	CreatedAt time.Time
	Likes     []*UserLikePost
	Comments  []*Comment
}

func NewPost(userId, content string) (*Post, error) {
	id, err := nanoid.New(21)
	if err != nil {
		return nil, errors.New("failed to generate post id")
	}

	return &Post{
		Id:        id,
		UserId:    userId,
		Content:   content,
		CreatedAt: time.Now(),
		Likes:     make([]*UserLikePost, 0),
		Comments:  make([]*Comment, 0),
	}, nil
}

func (p *Post) AddLike(userId string) error {
	for i, like := range p.Likes {
		if like.UserId == userId {
			p.Likes = append(p.Likes[:i], p.Likes[i+1:]...)
			return nil
		}
	}

	like, err := NewUserLikePost(userId, p.Id)
	if err != nil {
		return err
	}

	p.Likes = append(p.Likes, like)
	return nil
}

func (p *Post) AddComment(userId, content string) error {
	c, err := NewComment(userId, p.Id, content)
	if err != nil {
		return err
	}

	p.Comments = append(p.Comments, c)
	return nil
}
