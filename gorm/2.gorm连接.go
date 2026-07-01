package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:root@tcp(8.141.120.178:3306)/sql-test?charset=utf8"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
		return
	}
	log.Println(db)
	type Info struct {
		Name  string
		Email string
	}
	var UserList []Info
	db.Table("users").Find(&UserList)
	log.Println(UserList)
	for _, user := range UserList {
		log.Println(user.Name, user.Email)
	}
}
