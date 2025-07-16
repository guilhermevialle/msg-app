package routes

import (
	"app/internal/infra/http/controllers"
	"app/internal/infra/http/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, am middlewares.IAuthMiddleware, uc controllers.IUserController) {
	r.GET("/me/profile", am.Handle(), uc.GetProfile)
}
