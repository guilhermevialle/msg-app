package routes

import (
	"app/backend/internal/infra/http/controllers"
	"app/backend/internal/infra/http/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, am middlewares.IAuthMiddleware, uc controllers.IUserController) {
	r.GET("/me/profile", am.Handle(), uc.GetProfile)
}
