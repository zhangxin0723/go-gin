package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func controller(c *gin.Context) {
	log.Println(c.Request.Method, c.Request.URL)
}

func main() {
	r := gin.Default()

	apiGroup := r.Group("/api")
	{
		apiGroup.GET("/", controller)
		apiGroup.POST("/", controller)
		apiGroup.DELETE("/", controller)
		apiGroup.PUT("/", controller)
		apiGroup.PATCH("/", controller)
	}
	r.Run(":8080")
}
