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

		//src, err := fileHeader.Open()
		//if err != nil {
		//	log.Println(err)
		//	return
		//}
		//defer src.Close()
		//byteData, err := io.ReadAll(src)
		//err = os.WriteFile("upload.jpg", byteData, 0777)
		//log.Println(err)

		err = c.SaveUploadedFile(fileHeader, "./upload/"+fileHeader.Filename)
		log.Println(fileHeader.Filename, err)
	})
	r.Run(":8080")
}
