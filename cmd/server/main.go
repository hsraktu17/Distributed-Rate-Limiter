package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.GET("/api/check", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})
}
