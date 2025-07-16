package routes

import (
	"app/backend/internal/infra/http/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterDebugRoutes(r *gin.Engine, dc controllers.IDebugController) {
	r.GET("/debug", dc.DebugAll)
}
