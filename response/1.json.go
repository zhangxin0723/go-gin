package main

import (
	"go-gin/response/res"

	"github.com/gin-gonic/gin"
)

func Index(c *gin.Context) {
	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "Hello World",
		"data": nil,
	})
	res.Ok(c, nil, "成功")
	res.OkWithData(c, "成功")
	res.OkWithMsg(c, "成功")
	res.FailWithData(c, gin.H{"code": 500})
}

func Login(c *gin.Context) {
	//res.OkWithMsg(c, "登录成功")
	res.OkWithData(c, map[string]any{"name": "zx"})
}

func Logout(c *gin.Context) {
	res.FailWithMsg(c, "登录失败")
}

func main() {
	r := gin.Default()

	r.GET("/index", Index)
	r.POST("/login", Login)
	r.POST("/logout", Logout)

	r.Run(":8080")
}
