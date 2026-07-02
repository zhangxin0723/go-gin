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

	//// 通过用户详情查用户 正查
	//var id = 18
	//var detail models.UserDetailModel
	//global.DB.Preload("UserModel").Take(&detail, "user_id = ?", id)
	//log.Println(detail.Email, detail.UserModel.Name)
	//
	//// 反查
	//var user models.UserModel
	//global.DB.Preload("UserDetailModel").First(&user, id)
	//log.Println(user.Name, user.UserDetailModel.Email)

	// 删除数据
	// 级联删除
	//var user models.UserModel
	//global.DB.First(&user, 18)
	//global.DB.Select("UserDetailModel").Delete(&user)

	// 非级联删除
	//var user models.UserModel
	//global.DB.Take(&user, 19)
	//global.DB.Delete(&user)
	//global.DB.Model(&user).Association("UserDetailModel").Clear()

	// 或者设置model  前提得生成实体外键才行，而且修改关系之后要重新删掉实体外键再生成
	//UserDetailModel *UserDetailModel `gorm:"foreignKey:UserID;constraint:OnDelete:CASCADE"`  // CASCADE   或者  SET NULL
}

func main() {
	one2one()
}
