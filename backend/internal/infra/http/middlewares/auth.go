package middlewares

import (
	infra_services "app/internal/infra/services"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type IAuthMiddleware interface {
	Handle() gin.HandlerFunc
}

type AuthMiddleware struct {
	tokenService infra_services.ITokenService
}

var _ IAuthMiddleware = (*AuthMiddleware)(nil)

func NewAuthMiddleware(ts infra_services.ITokenService) *AuthMiddleware {
	return &AuthMiddleware{
		tokenService: ts,
	}
}

func (am *AuthMiddleware) Handle() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header missing"})
			return
		}

		parts := strings.SplitN(header, " ", 2)
		if len(parts) != 2 || strings.ToLower(parts[0]) != "bearer" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header format must be Bearer <token>"})
			return
		}

		requestToken := parts[1]

		token, err := am.tokenService.Validate(requestToken)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		if token == nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "invalid token"})
			return
		}

		c.Set("userId", token.UserId)
		c.Next()
	}
}
