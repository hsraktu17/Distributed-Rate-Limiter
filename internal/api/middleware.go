package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/hsraktu17/Distributed-Rate-Limiter/internal/limiter"
)

var globalLimiterMap = make(map[string]*limiter.Limiter)

func RateLimiterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.Query("user_id")
		if userID == "" {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"error": "Missing user_id"})
			return
		}

		mu.Lock()
		l, exists := globalLimiterMap[userID]
		if !exists {
			l = limiter.NewLimiter(5, 10) // Customize here
			globalLimiterMap[userID] = l
		}
		mu.Unlock()

		if !l.Allow() {
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{"error": "Too Many Requests"})
			return
		}

		c.Next() // Allow request to continue
	}
}
