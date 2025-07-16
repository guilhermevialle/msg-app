package routes

import (
	"app/backend/internal/infra/http/controllers"
	"app/backend/internal/infra/http/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterPostRoutes(r *gin.Engine, am middlewares.IAuthMiddleware, pc controllers.IPostController) {
	r.POST("/post/new", am.Handle(), pc.CreateUserPost)
	r.GET("/post/:post_id/like", am.Handle(), pc.LikeUserPost)
	r.POST("/post/:post_id/comment", am.Handle(), pc.CreateCommentOnPost)
}
