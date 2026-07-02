package main

import (
	"go-gin/global"
	"go-gin/models"
	"log"
)

func init() {
	global.InitDB()
	err := global.DB.AutoMigrate(&models.ArticleModel{}, &models.TagModel{})
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func many2Many() {
	// 创建一篇文章，新增tag
	//err := global.DB.Create(&models.ArticleModel{
	//	Title: "Go多对多关系",
	//	TagModel: []models.TagModel{
	//		{
	//			Title: "GO",
	//		},
	//		{
	//			Title: "GO-gorm",
	//		},
	//	},
	//}).Error
	//if err != nil {
	//	log.Fatalln(err)
	//}

	// 创建一篇文章选择tag
	//var tagIdList = []int{1, 2}
	//var tagList []models.TagModel
	//global.DB.Find(&tagList, "id in ?", tagIdList)
	//
	//err := global.DB.Create(&models.ArticleModel{
	//	Title:    "多对多6",
	//	TagModel: tagList,
	//}).Error
	//if err != nil {
	//	log.Fatalln(err)
	//}

	// 查询文章携带标签
	var articles []models.ArticleModel
	err := global.DB.Preload("TagModel").Find(&articles).Error
	if err != nil {
		log.Fatalln(err)
	}
	for _, article := range articles {
		for _, tag := range article.TagModel {
			log.Println(article.Title, "---", tag.Title)
		}
	}

	// 更新
	//var article models.ArticleModel
	//global.DB.Take(&article, 4)
	//global.DB.Take(&article).Association("TagModel").Replace([]models.TagModel{
	//	{ID: 2},
	//})
}

func main() {
	many2Many()
}
