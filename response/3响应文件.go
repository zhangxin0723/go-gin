package main

import "github.com/gin-gonic/gin"

// 浏览器下载
func fileRes(c *gin.Context) {
	c.Header("Content-Type", "application/octet-stream")              // 表示是文件流，唤起浏览器下载，一般设置了这个，就要设置文件名
	c.Header("Content-Disposition", "attachment; filename=3.响应文件.go") // 用来指定下载下来的文件名
	c.File("response/3.响应文件.go")
}

func main() {
	r := gin.Default()
	r.GET("/", fileRes)
	r.Run(":8080")
}
