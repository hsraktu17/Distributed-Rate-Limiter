package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/hsraktu17/Distributed-Rate-Limiter/internal/api"
)

func main() {
	r := gin.Default()
	r.Use(api.RateLimiterMiddleware())
	r.GET("/api/check", api.RateLimitHandle)

	log.Println("Server is running on port 8080")
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Failed to run server: %v", err)
	}
}
