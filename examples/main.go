package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	ginratelimiter "github.com/hsraktu17/Distributed-Rate-Limiter/gin"
	"github.com/hsraktu17/Distributed-Rate-Limiter/limiter"
)

func main() {
	r := gin.Default()

	lim := limiter.NewMemoryLimiter(5, 10)
	r.Use(ginratelimiter.Middleware(lim))

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong"})
	})

	log.Println("Server running on http://localhost:8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("failed to run server: %v", err)
	}
}
