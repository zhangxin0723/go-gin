package main

import (
	"go-gin/global"
	"go-gin/models"
	"log"
)

func init() {
	global.InitDB()
	err := global.DB.AutoMigrate(&models.UserModel{}, &models.UserDetailModel{})
	if err != nil {
		log.Fatalln(err)
		return
	}
}

// 一对一创建
func one2one() {
	// 创建用户并且创建用户详情
	//err := global.DB.Create(&models.UserModel{
	//	Name: "zx",
	//	Age:  18,
	//	UserDetailModel: &models.UserDetailModel{
	//		Email: "zx@qq.com",
	//	},
	//}).Error
	//if err != nil {
	//	log.Fatalln(err)
	//}

	// 创建详情 关联用户
	//err := global.DB.Create(&models.UserDetailModel{
	//	Email: "zx1@qq.com",
	//	//UserId: 18,
	//	UserModel: &models.UserModel{Id: 18},
	//}).Error
	//if err != nil {
	//	log.Fatalln(err)
	//}

	// 通过用户详情查用户 正查
	var id = 18
	var detail models.UserDetailModel
	global.DB.Preload("UserModel").Take(&detail, "user_id = ?", id)
	log.Println(detail.Email, detail.UserModel.Name)

	// 反查
	var user models.UserModel
	global.DB.Preload("UserDetailModel").First(&user, id)
	log.Println(user.Name, user.UserDetailModel.Email)
}

func main() {
	one2one()
}
