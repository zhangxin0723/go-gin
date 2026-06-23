package res

import "github.com/gin-gonic/gin"

type Response struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func response(c *gin.Context, code int, data any, msg string) {
	c.JSON(200, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}

func Ok(c *gin.Context, data any, msg string) {
	response(c, 200, data, msg)
}

func OkWithData(c *gin.Context, data any) {
	Ok(c, data, "成功")
}

func OkWithMsg(c *gin.Context, msg string) {
	Ok(c, gin.H{}, msg)
}

func Fail(c *gin.Context, code int, data any, msg string) {
	response(c, code, data, msg)
}

func FailWithData(c *gin.Context, data any) {
	Fail(c, 500, data, "失败")
}

func FailWithMsg(c *gin.Context, msg string) {
	Fail(c, 500, gin.H{}, msg)
}
