package controllers

import (
	"app/internal/application/use_cases"
	"app/internal/dtos"
	"app/internal/infra/http/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IPostController interface {
	CreateUserPost(c *gin.Context)
	LikeUserPost(c *gin.Context)
	CreateCommentOnPost(c *gin.Context)
	GetPosts(c *gin.Context)
}

type PostController struct {
	createPost       *use_cases.CreatePost
	likePost         *use_cases.LikePost
	addCommentToPost *use_cases.AddCommentToPost
	getAllPosts      *use_cases.GetAllPosts
}

var _ IPostController = (*PostController)(nil)

func NewPostController(
	cp *use_cases.CreatePost,
	lp *use_cases.LikePost,
	actp *use_cases.AddCommentToPost,
	gap *use_cases.GetAllPosts,
) *PostController {
	return &PostController{
		createPost:       cp,
		likePost:         lp,
		addCommentToPost: actp,
		getAllPosts:      gap,
	}
}

func (pc *PostController) CreateUserPost(c *gin.Context) {
	id, ok := utils.GetUserIdOrAbort(c)
	if !ok {
		return
	}

	var body dtos.CreatePostDto

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	p, err := pc.createPost.Execute(id, body.Content)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, p)
}

func (pc *PostController) LikeUserPost(c *gin.Context) {
	id, ok := utils.GetUserIdOrAbort(c)
	if !ok {
		return
	}

	postId := c.Param("post_id")
	if postId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "post_id is required"})
		return
	}

	err := pc.likePost.Execute(id, postId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "post liked successfully"})
}

func (pc *PostController) CreateCommentOnPost(c *gin.Context) {
	id, ok := utils.GetUserIdOrAbort(c)
	if !ok {
		return
	}

	postId := c.Param("post_id")
	if postId == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "post_id is required"})
		return
	}

	var body dtos.CreateCommentOnPostDto

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if body.Content == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "content is required"})
		return
	}

	comment, err := pc.addCommentToPost.Execute(id, postId, body.Content)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comment)
}

func (pc *PostController) GetPosts(c *gin.Context) {
	posts := pc.getAllPosts.Execute()
	c.JSON(http.StatusOK, posts)
}
