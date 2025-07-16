package controllers

import (
	"app/backend/internal/application/use_cases"
	"app/backend/internal/infra/http/utils"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IUserController interface {
	GetProfile(c *gin.Context)
}

type UserController struct {
	getUserProfile *use_cases.GetUserProfile
}

var _ IUserController = (*UserController)(nil)

func NewUserController(gup *use_cases.GetUserProfile) *UserController {
	return &UserController{
		getUserProfile: gup,
	}
}

func (uc *UserController) GetProfile(c *gin.Context) {
	id, ok := utils.GetUserIdOrAbort(c)
	if !ok {
		return
	}

	p, err := uc.getUserProfile.Execute(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, p)
}
