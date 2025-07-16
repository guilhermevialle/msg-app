package controllers

import (
	dtos "app/backend/internal"
	"app/backend/internal/application/use_cases"
	"app/backend/internal/infra/http/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IPostController interface {
	CreateUserPost(c *gin.Context)
	LikeUserPost(c *gin.Context)
	CreateCommentOnPost(c *gin.Context)
}

type PostController struct {
	CreatePost       *use_cases.CreatePost
	LikePost         *use_cases.LikePost
	AddCommentToPost *use_cases.AddCommentToPost
}

var _ IPostController = (*PostController)(nil)

func NewPostController(
	cp *use_cases.CreatePost,
	lp *use_cases.LikePost,
	actp *use_cases.AddCommentToPost,
) *PostController {
	return &PostController{
		CreatePost:       cp,
		LikePost:         lp,
		AddCommentToPost: actp,
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

	p, err := pc.CreatePost.Execute(id, body.Content)
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

	err := pc.LikePost.Execute(id, postId)
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

	comment, err := pc.AddCommentToPost.Execute(id, postId, body.Content)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, comment)
}
