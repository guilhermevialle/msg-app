package repositories

import "app/backend/internal/domain/entities"

type IPostRepository interface {
	Save(post *entities.Post)
	FindById(id string) *entities.Post
	Update(post *entities.Post)
	FindAll() []*entities.Post
}

type PostRepository struct {
	posts []*entities.Post
}

var _ IPostRepository = (*PostRepository)(nil)

func NewPostRepository() *PostRepository {
	return &PostRepository{
		posts: make([]*entities.Post, 0),
	}
}

func (pr *PostRepository) Save(post *entities.Post) {
	pr.posts = append(pr.posts, post)
}

func (pr *PostRepository) FindById(id string) *entities.Post {
	for _, p := range pr.posts {
		if p.Id == id {
			return p
		}
	}
	return nil
}

func (pr *PostRepository) Update(post *entities.Post) {
	for i, p := range pr.posts {
		if p.Id == post.Id {
			pr.posts[i] = post
			return
		}
	}
}

func (pr *PostRepository) FindAll() []*entities.Post {
	return pr.posts
}
