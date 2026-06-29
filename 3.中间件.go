package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func authMiddleware(c *gin.Context) {
	if Authorization := c.Request.Header.Get("Authorization"); Authorization == "" {
		c.String(500, "请登录")
		return
	}
	log.Println("authMiddleware1")
	c.Next()
	log.Println("authMiddleware2")
}

func M1(c *gin.Context) {
	log.Println("M1")
	c.Next()
	log.Println("M1-end")
}

func M2(c *gin.Context) {
	log.Println("M2")
	c.Next()
	//c.Abort()
	log.Println("M2-end")
}

func main() {
	r := gin.Default()

	//r.Use(authMiddleware)
	r.GET("/", authMiddleware, func(c *gin.Context) {
		c.JSON(200, gin.H{
			"code": 200,
		})
	})
	r.GET("/a", M1, M2, func(c *gin.Context) {
		log.Println("/a")
	})

	r.Run(":8080")
}
