package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.Static("static", "response/static")
	r.StaticFile("test", "response/static/test.txt")
	r.Run(":8080")
}
