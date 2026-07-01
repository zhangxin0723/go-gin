package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/sql-test?charset=utf8"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	//defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatalln("ping", err)
	}
	log.Println(db)

	res, err := db.Exec("insert into users (name, age, email) values ('张三', 18, 'zhangsan@example.com'), ('李四', 20, 'lisi@example.com')")
	if err != nil {
		log.Fatalln("sql-error", err)
	}
	log.Println(res)

	rows, _ := db.Query("select id,name from users where name = '张三'")
	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		log.Println(id, name, err)
	}
	var id int
	var name string
	err = db.QueryRow("select id,name from users where name = '张三'").Scan(&id, &name)
	log.Println(id, name, err)
}
