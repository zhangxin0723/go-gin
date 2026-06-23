package main

import "github.com/gin-gonic/gin"

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Index(c *gin.Context) {
	c.JSON(200, Response{
		Code: 200,
		Msg:  "Hello World",
		Data: nil,
	})
}

func main() {
	// 1.初始化
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	// 2.挂载路由
	r.GET("/index", Index)
	// 3.绑定端口
	r.Run(":8080")
}
