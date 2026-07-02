package main

import (
	"go-gin/global"
	"go-gin/models"
	"log"
)

func init() {
	global.InitDB()
	err := global.DB.AutoMigrate(&models.User1Model{}, &models.Article1Model{}, &models.User2ArticleModel{})
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func userMany2Many() {
	// 创建用户 连带文章
	// 必须要加这个才会走第三张表的创建钩子
	//global.DB.SetupJoinTable(&models.User1Model{}, "CollArticleList", &models.User2ArticleModel{})
	//err := global.DB.Create(&models.User1Model{
	//	Name: "zx5",
	//	CollArticleList: []models.Article1Model{
	//		{ID: 2},
	//		{ID: 3},
	//	},
	//}).Error
	//if err != nil {
	//	log.Fatalln(err, "error===")
	//}

	// 查询
	// 这样查不到文章收藏时间
	//var user models.User1Model
	//err := global.DB.Preload("CollArticleList").First(&user).Error
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//log.Println(user)

	//type UserCollArticleResponse struct {
	//	Name         string
	//	UserID       int64
	//	ArticleTitle string
	//	ArticleID    int64
	//	Date         time.Time
	//}
	//var userId = 4
	//var userArticleList []models.User2ArticleModel
	//var collList []UserCollArticleResponse
	//err := global.DB.Preload("UserModel").Preload("ArticleModel").Find(&userArticleList, "user_id = ?", userId).Error
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//for _, model := range userArticleList {
	//	collList = append(collList, UserCollArticleResponse{
	//		Name:         model.UserModel.Name,
	//		UserID:       model.UserID,
	//		ArticleTitle: model.ArticleModel.Title,
	//		ArticleID:    model.ArticleID,
	//		Date:         model.CreatedAt,
	//	})
	//}
	//fmt.Println(collList)

	// 删除
	//var user = models.User1Model{ID: 4}
	//global.DB.First(&user)
	//global.DB.Model(&user).Association("CollArticleList").Clear()
}

func main() {
	userMany2Many()
}
