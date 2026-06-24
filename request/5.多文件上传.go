package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/uploads", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			log.Println(err)
			return
		}
		for _, headers := range form.File {
			for _, header := range headers {
				log.Println(header.Filename)
				c.SaveUploadedFile(header, "./uploads/"+header.Filename)
			}
		}
	})
	r.Run(":8080")
}
