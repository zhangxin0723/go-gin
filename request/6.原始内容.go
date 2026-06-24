package main

import (
	"log"
	"net/url"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/", func(c *gin.Context) {
		//byteData, err := io.ReadAll(c.Request.Body)
		//if err != nil {
		//	return
		//}
		//log.Println(string(byteData))
		//
		//// 读了body之后 body就没了 阅后即焚
		//// 解决方法
		//c.Request.Body = io.NopCloser(bytes.NewBuffer(byteData))
		//
		//name := c.PostForm("name")
		//log.Println(name)

		valus, _ := url.ParseQuery("?")
		name1 := valus.Get("name=")
		log.Println(name1)
	})
	r.Run(":8080")
}
