package api

import (
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"github.com/hsraktu17/Distributed-Rate-Limiter/internal/limiter"
)

var limiterMap = make(map[string]*limiter.Limiter)
var mu sync.Mutex

func RateLimitHandle(c *gin.Context) {
	userID := c.Query("user_id")

	if userID == "" {
		c.JSON(400, gin.H{
			"error": "user_id is required",
		})
	}

	mu.Lock()
	l, exists := limiterMap[userID]
	if !exists {
		// Create a new limiter for the user: 5 req/sec, burst of 10
		l = limiter.NewLimiter(5, 10)
		limiterMap[userID] = l
	}
	mu.Unlock()

	if l.Allow() {
		c.JSON(http.StatusOK, gin.H{
			"message": "Request allowed",
		})
	} else {
		c.JSON(http.StatusTooManyRequests, gin.H{
			"error": "Rate limit exceeded",
		})
	}
}
