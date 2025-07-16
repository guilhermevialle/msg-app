package api

import (
	"app/internal/application/di"
	"app/internal/infra/http/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func NewApp() *gin.Engine {
	r := gin.Default()
	c := di.NewContainer()

	// CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:5173"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Authorization"}
	r.Use(cors.New(config))

	// Routes registry
	routes.RegisterAuthRoutes(r, c.AuthController)
	routes.RegisterUserRoutes(r, c.AuthMiddleware, c.UserController)
	routes.RegisterPostRoutes(r, c.AuthMiddleware, c.PostController)
	routes.RegisterDebugRoutes(r, c.DebugController)

	return r
}
