package routes

import (
	"app/backend/internal/infra/http/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAuthRoutes(r *gin.Engine, ac controllers.IAuthController) {
	r.POST("/auth/login", ac.AuthenticateUser)
	r.POST("/auth/register", ac.RegisterUser)
}
