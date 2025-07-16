package utils

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetUserIdOrAbort(c *gin.Context) (string, bool) {
	rawId, exists := c.Get("userId")
	if !exists {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "unauthorized"})
		return "", false
	}

	id, ok := rawId.(string)
	if !ok {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "userId is not a string"})
		return "", false
	}

	if id == "" {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "userId is required"})
		return "", false
	}

	return id, true
}
