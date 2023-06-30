package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
)

func main() {

	// 事务
	dsn := "root:root@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4"
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		fmt.Println("数据库连接失败！")
		return
	}

	type User struct {
		Id      int64
		Name    string
		Salt    string
		Age     int
		Passwd  string    `xorm:"varchar(200)"`
		Created time.Time `xorm:"created"`
		Updated time.Time `xorm:"updated"`
	}

	session := engine.NewSession()
	defer session.Close()

	// 开启事务
	session.Begin()

	// 处理错误业务
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
			// 回滚
			session.Rollback()
		} else {
			session.Commit()
		}
	}()

	user1 := User{Id: 10, Name: "laoba", Age: 66, Passwd: "123456"}
	if _, err := session.Insert(&user1); err != nil {
		panic(err)
	}

	user2 := User{Name: "sxs002", Age: 28}
	if _, err := session.Where("id=?", 7).Update(&user2); err != nil {
		panic(err)
	}

	if _, err := session.Exec("delete from user where name=?", "555"); err != nil {
		panic(err)
	}

}
