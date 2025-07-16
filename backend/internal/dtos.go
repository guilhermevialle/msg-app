package dtos

type LoginDto struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type RegisterDto struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type CreatePostDto struct {
	Content string `json:"content"`
}

type CreateCommentOnPostDto struct {
	Content string `json:"content"`
}
