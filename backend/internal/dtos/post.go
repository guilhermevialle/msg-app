package dtos

type PostAuthor struct {
	Id        string `json:"id"`
	Username  string `json:"username"`
	AvatarUrl string `json:"avatar_url"`
}

type PostComment struct {
	Id        string      `json:"id"`
	PostId    string      `json:"post_id"`
	ParentId  *string     `json:"parent_id"`
	Content   string      `json:"content"`
	CreatedAt string      `json:"created_at"`
	Author    *PostAuthor `json:"author"`
	Replies   []*PostComment
}

type PostResponseDto struct {
	Id        string      `json:"id"`
	Content   string      `json:"content"`
	CreatedAt string      `json:"created_at"`
	Likes     int         `json:"likes"`
	Author    *PostAuthor `json:"author"`
}
