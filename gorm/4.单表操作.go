package main

import (
	"go-gin/global"
	"go-gin/models"
	"log"

	"gorm.io/gorm"
)

func init() {
	global.InitDB()
}

func insert() {
	//err := global.DB.Create(&models.UserModel{
	//	Name: "张三",
	//}).Error
	//if err != nil {
	//	log.Fatalln(err)
	//	return
	//}

	// 回填式创建
	//user := models.UserModel{
	//	Name: "张三三",
	//}
	//err := global.DB.Create(&user).Error
	//if err != nil {
	//	log.Fatalln(err)
	//	return
	//}
	//log.Println(user, user.Id)

	// 分片添加
	var userList = []models.UserModel{
		{
			Name: "李四3",
		},
		{
			Name: "王五3",
		},
	}
	err := global.DB.Create(&userList).Error
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println(userList)
}

func query() {
	var userList []models.UserModel
	global.DB.Find(&userList, "age >= ?", 19)
	log.Println(userList)

	//var user models.UserModel
	//global.DB.Take(&user)
	// 根据主键去查
	//global.DB.Take(&user, 14)
	// 根据主键查询第一个
	//global.DB.First(&user, 1)
	//global.DB.Last(&user, 1)
	//log.Println(user)

	//user.Id = 6
	//global.DB.Take(&user)
	//log.Println(user)

	var user models.UserModel
	// take查询不到会报错
	//err := global.DB.Take(&user, 111111).Error
	//if err == gorm.ErrRecordNotFound {
	//	log.Println("记录不存在", err)
	//}
	// find 查询不会报错
	global.DB.Debug().Limit(1).Find(&user, 1111)
	log.Println(user)
}

func updateSave() {
	// 根据主键 无则创建 有则更新
	var user models.UserModel
	//user.Name = "李四4"

	//user.Id = 15
	//user.Name = "李四4-1"
	//user.Age = 10
	//user.CreatedAt = time.Now()
	global.DB.Save(&user)
}

func update() {
	//var user = models.UserModel{Id: 15}
	//global.DB.Model(&user).Update("age", 11).Update("name", "李四4-1-1") // 唯一区别会执行更新钩子函数
	//global.DB.Model(&user).UpdateColumn("age", 11).UpdateColumn("name", "李四4-1-1")
}

func updates() {
	var user = models.UserModel{Id: 15}
	// 如果是结构体，不可以更新非零值，map可以更新零值
	global.DB.Model(&user).Updates(models.UserModel{
		Name: "李四4-1-1-1",
		Age:  0,
	})
	//global.DB.Model(&user).Updates(map[string]interface{}{
	//	"name": "李四4-1-1-1",
	//	"age":  0,
	//})
}

func expr() {
	var user = models.UserModel{Id: 15}
	global.DB.Model(&user).Update("age", gorm.Expr("age + 1"))
}

func deleteUser() {
	//var user = models.UserModel{Id: 14}
	//global.DB.Delete(&user)

	//global.DB.Delete(&models.UserModel{}, 13)

	global.DB.Delete(&models.UserModel{}, "age = ?", []int64{25})

	// 软删除
	// 模型带有deleteAt字段默认开启软删除 不可被查询 可通过Unscoped()查到
}

func highQuery() {
	var user models.UserModel
	var userList []models.UserModel

	//global.DB.Where("name = ?", "张三").Find(&user)

	//query := global.DB.Where("age > ?", 18)
	//global.DB.Where(query).Take(&user)

	//global.DB.Or("name = ?", "李四4-1-1-1").Or("age = 18").Find(&userList)

	//global.DB.Not("age < 20").Find(&userList)

	//global.DB.Order("age desc").Find(&userList)

	//var names []string
	//global.DB.Model(&user).Select("name").Scan(&names)
	//global.DB.Model(&user).Pluck("name", &names)
	//log.Println(names)

	//type Info struct {
	//	Name string
	//	Age  int
	//}
	//type Info struct {
	//	Label string `gorm:"column:name"`
	//	Value int    `gorm:"column:age"`
	//}
	//var list []Info
	//global.DB.Model(&models.UserModel{}).Scan(&list)
	//log.Println(list)

	//type UserCount struct {
	//	Age   int
	//	Count int
	//}
	//var userCount []UserCount
	//global.DB.Model(&models.UserModel{}).Group("age").Select("age", "count(id) as count").Scan(&userCount)
	//log.Println(userCount)

	//去重
	//var ageList []int64
	//global.DB.Model(&models.UserModel{}).Distinct("age").Pluck("age", &ageList)
	//log.Println(ageList)

	log.Println(user, userList)
}

func pageQuery(limit int, offset int) {
	var userList []models.UserModel
	if limit < 1 {
		limit = 1
	}
	if offset < 1 {
		offset = 1
	}
	global.DB.Limit(limit).Offset(limit * (offset - 1)).Find(&userList)
	log.Println(userList)
}

func main() {
	//insert()
	//query()
	//updateSave()
	//update()
	//updates()
	//expr()
	//deleteUser()
	//highQuery()
	//pageQuery(1, 1)
	//pageQuery(1, 2)
	//pageQuery(0, 0)
}
