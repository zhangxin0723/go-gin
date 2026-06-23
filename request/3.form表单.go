package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/upload", func(c *gin.Context) {
		fileHeader, err := c.FormFile("file")
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(fileHeader.Filename, fileHeader.Size)

		//file, _ := fileHeader.Open()
		//byteData, _ := io.ReadAll(file)
		//
		//err = os.WriteFile("upload.jpg", byteData, 0777)
		//log.Println(err)

		err = c.SaveUploadedFile(fileHeader, "./upload/"+fileHeader.Filename)
		log.Println(err)
	})
	r.Run(":8080")
}
