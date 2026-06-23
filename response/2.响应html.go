package main

import "github.com/gin-gonic/gin"

func htmlRes(c *gin.Context) {
	c.HTML(200, "index.html", gin.H{
		"title": "zx",
	})
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("response/templates/*")
	r.GET("/", htmlRes)

	r.Run(":8080")
}
