package main

import (
	"go-gin/global"
	"go-gin/models"
	"log"
)

func init() {
	global.InitDB()
	err := global.DB.AutoMigrate(&models.GirlModel{}, &models.BoyModel{})
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func one2Many() {
	// 添加数据
	// 来了一个女神 带俩舔狗
	//err := global.DB.Create(&models.GirlModel{
	//	Name: "林志玲",
	//	BoyList: []models.BoyModel{
	//		{
	//			Name: "小张",
	//		},
	//		{
	//			Name: "小王",
	//		},
	//	},
	//}).Error
	//if err != nil {
	//	log.Fatalln(err)
	//}
	// 来了一个女神 选了一个舔狗
	//err := global.DB.Create(&models.GirlModel{
	//	Name: "张曼玉",
	//	BoyList: []models.BoyModel{
	//		{
	//			ID: 2,
	//		},
	//	},
	//}).Error
	//if err != nil {
	//	log.Fatalln(err)
	//}
	// 来了一个舔狗 选了一个女神
	//err := global.DB.Create(&models.BoyModel{
	//	Name: "小李",
	//	GirlModel: &models.GirlModel{
	//		ID: 1,
	//	},
	//}).Error
	//if err != nil {
	//	log.Fatalln(err)
	//}

	// 查询
	// 查询林志玲的舔狗
	//var girl models.GirlModel
	//global.DB.Preload("BoyList").First(&girl, "name = ?", "林志玲")
	//log.Println(girl, girl.BoyList, len(girl.BoyList))
	//
	// 专门查关联
	//girl = models.GirlModel{
	//	ID: 1,
	//}
	//var boyList []models.BoyModel
	//global.DB.Debug().Model(&girl).Association("BoyList").Find(&boyList)
	//log.Println(boyList)

	// 关联操作
	// 林志玲舔狗不舔了 换了一个人舔
	// replace
	//girl := models.GirlModel{
	//	ID: 1,
	//}
	//global.DB.Model(&girl).Association("BoyList").Replace([]models.BoyModel{
	//	{
	//		ID: 2,
	//	},
	//})

	// 都不舔了
	// Clear
	//girl := models.GirlModel{
	//	ID: 1,
	//}
	//global.DB.Model(&girl).Association("BoyList").Clear()

	// 又来了两个人舔狗
	// Append
	//girl := models.GirlModel{
	//	ID: 1,
	//}
	//global.DB.Model(&girl).Association("BoyList").Append([]models.BoyModel{
	//	{
	//		ID: 1,
	//	},
	//	{
	//		ID: 2,
	//	},
	//})

	// 只退出了一个舔狗
	// Delete
	girl := models.GirlModel{
		ID: 1,
	}
	global.DB.Model(&girl).Association("BoyList").Delete([]models.BoyModel{
		{
			ID: 2,
		},
	})
}

func main() {
	one2Many()
}
