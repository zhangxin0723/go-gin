package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 查询参数
	r.GET("", func(c *gin.Context) {
		type User struct {
			Name string `form:"name"`
			Age  int    `form:"age"`
		}
		var user User
		err := c.ShouldBindQuery(&user)
		log.Println(user, err)
	})

	// 路径参数
	r.GET("/user/:name/:id", func(c *gin.Context) {
		type User struct {
			Name string `uri:"name"`
			Id   int    `uri:"id"`
		}
		var user User
		err := c.ShouldBindUri(&user)
		log.Println(user, err)
	})

	// formData
	r.POST("/upload", func(c *gin.Context) {
		type User struct {
			Name string `form:"name"`
			Age  int    `form:"age"`
		}
		var user User
		err := c.ShouldBind(&user)
		log.Println(user, err)
	})

	r.POST("/json", func(c *gin.Context) {
		type User struct {
			Name string `json:"name"`
			Age  int    `json:"age"`
		}
		var user User
		err := c.ShouldBind(&user)
		log.Println(user, err)
	})

	r.POST("/header", func(c *gin.Context) {
		type Headers struct {
			AppToken  string `header:"apptoken"`
			UserToken string `header:"usertoken"`
		}
		var headers Headers
		err := c.ShouldBindHeader(&headers)
		log.Println(headers, err)
	})
	r.Run(":8080")
}
