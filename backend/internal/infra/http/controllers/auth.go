package controllers

import (
	app_services "app/internal/application/services"
	"app/internal/dtos"
	"net/http"

	"github.com/gin-gonic/gin"
)

type IAuthController interface {
	AuthenticateUser(c *gin.Context)
	RegisterUser(c *gin.Context)
}

type AuthController struct {
	authService app_services.IAuthService
}

var _ IAuthController = (*AuthController)(nil)

func NewAuthController(as app_services.IAuthService) *AuthController {
	return &AuthController{authService: as}
}

func (ac *AuthController) AuthenticateUser(c *gin.Context) {
	var body *dtos.LoginDto

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := ac.authService.Authenticate(body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

func (ac *AuthController) RegisterUser(c *gin.Context) {
	var body *dtos.RegisterDto

	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := ac.authService.Register(body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "user registered successfully"})
}
