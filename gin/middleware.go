package ginratelimiter

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hsraktu17/Distributed-Rate-Limiter/limiter"
)

// Middleware returns a Gin middleware using the provided Limiter.
func Middleware(l limiter.Limiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Query("user_id")
		if userID == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing user_id"})
			return
		}

		if !l.Allow(userID) {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Too Many Requests"})
			return
		}

		c.Next()
	}
}
