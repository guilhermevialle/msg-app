package controllers

import (
	"app/internal/infra/repositories"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IDebugController interface {
	DebugAll(c *gin.Context)
}

type DebugController struct {
	userRepo repositories.IUserRepository
	postRepo repositories.IPostRepository
}

func NewDebugController(ur repositories.IUserRepository, pr repositories.IPostRepository) *DebugController {
	return &DebugController{
		userRepo: ur,
		postRepo: pr,
	}
}

func (dc *DebugController) DebugAll(c *gin.Context) {
	u := dc.userRepo.FindAll()
	p := dc.postRepo.FindAll()

	c.JSON(http.StatusOK,
		gin.H{
			"users": u,
			"posts": p,
		},
	)
}
