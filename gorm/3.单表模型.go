package main

import (
	"go-gin/global"
	"go-gin/models"
	"log"
)

func init() {
	global.InitDB()
}

func main() {
	err := global.DB.AutoMigrate(&models.UserModel{})
	if err != nil {
		log.Fatalln(err)
		return
	}
}
