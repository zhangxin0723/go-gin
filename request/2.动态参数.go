package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/users/:id", func(c *gin.Context) {
		userId := c.Param("id")
		log.Println(userId)
	})
	r.Run(":8080")
}
