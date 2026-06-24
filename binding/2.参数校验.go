package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("", func(c *gin.Context) {
		type User struct {
			//Name string `json:"name" binding:"required,min=2,max=5"`
			//Age  int    `json:"age" binding:"gt=0"`
			// 同级校验
			//Password        string `json:"password" binding:"required"`
			//ConfirmPassword string `json:"confirm_password" binding:"required,eqfield=Password"`
			// 枚举
			//Color string `json:"color" binding:"required,oneof=red blue"`
			// 字符串
			//Str string `json:"str" binding:"required,contains=zx"` // 包含
			//Str string `json:"str" binding:"required,excludes=zx"` // 不包含
			//Str string `json:"str" binding:"required,startswith=zx"` // 开头
			//Str string `json:"str" binding:"required,endswith=zx"` // 结尾
			// 网络校验
			//Ip string `json:"ip" binding:"required,ip"`
			//Url string `json:"url" binding:"required,url"`
			//Uri string `json:"uri" binding:"required,uri"`
			// 数组
			Arr []string `json:"arr" binding:"required,min=1,dive,ip"`
		}
		var user User
		err := c.ShouldBindJSON(&user)
		if err != nil {
			c.JSON(200, gin.H{
				"code": 400,
				"msg":  err.Error(),
			})
			return
		}
		c.JSON(200, user)
		log.Println(user, err)
	})
	r.Run(":8080")
}
