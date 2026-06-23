package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func queryParams(c *gin.Context) {
	name := c.Query("name")
	age := c.DefaultQuery("age", "25")
	keyList := c.QueryArray("key")
	log.Println(name, age, keyList)
}

func main() {
	r := gin.Default()
	r.GET("/", queryParams)
	r.Run(":8080")
	// http://127.0.0.1:8080/?name=1&age=2&key=1&key=2&key=3
}
