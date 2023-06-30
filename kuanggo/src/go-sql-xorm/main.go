package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"time"
	"xorm.io/xorm"
)

type User struct {
	Id      int64
	Name    string
	Salt    string
	Age     int
	Passwd  string    `xorm:"varchar(200)"`
	Created time.Time `xorm:"created"`
	Updated time.Time `xorm:"updated"`
}

func main() {
	/*var (
		userName  string = "root"
		password  string = "root"
		ipAddress string = "127.0.0.1"
		port      int    = 3306
		dbName    string = "go_test"
		charset   string = "utf8mb4"
	)*/

	//dsn := fmt.Print("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddress, port, dbName, charset)

	dsn := "root:root@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4"
	engine, err := xorm.NewEngine("mysql", dsn)
	if err != nil {
		fmt.Println("数据库连接失败！")
		return
	}

	// 创建User的表结构
	//errCrete := engine.Sync(new(User))
	//if errCrete != nil {
	//	fmt.Println("同步数据失败！")
	//}

	// 添加User的数据 Insert() 插入 对象，返回值：受影响的行数
	user := User{Id: 4, Name: "sxs", Age: 18, Passwd: "123456"}
	user1 := User{Id: 5, Name: "kuangshen", Age: 18, Passwd: "123456"}
	n, _ := engine.Insert(&user, &user1)
	if n > 0 {
		fmt.Println("数据插入成功")
		fmt.Println("n的值=>", n)
	}

	// 使用切片进行插入多条数据
	var users []User
	users = append(users, User{Id: 6, Name: "zhangsan", Age: 10, Passwd: "123456"})
	users = append(users, User{Id: 7, Name: "lisi", Age: 30, Passwd: "123456"})
	n, err = engine.Insert(&users)
	if n > 0 {
		fmt.Println("插入数据成功")
		fmt.Println(n)
	} else {
		fmt.Println(err)
	}

	// 更新
	useradd := User{Name: "wangwu"}
	ures, _ := engine.ID(1).Update(&useradd)
	if ures > 0 {
		fmt.Println("更新成功")
		fmt.Println(ures)
	}
	// 删除
	userDel := User{Name: "sxs"}
	dres, _ := engine.ID(4).Delete(&userDel)
	if dres > 0 {
		fmt.Println("删除成功")
		fmt.Println(dres)
	}

	// Exec sql版
	engine.Exec("update user set age=? where id=?", 18, 7)

	// 查询
	result, err := engine.Query("select * from user")
	fmt.Println(result)

	result2, err := engine.QueryString("select * from user")
	fmt.Println(result2)
	fmt.Println("-------------------")

	result3, err := engine.QueryInterface("select * from user")
	fmt.Println(result3)

	// Get
	userGet := User{}
	engine.Get(&userGet)
	fmt.Println(userGet)

	// 根据条件进行查询
	user10 := User{Name: "sxs"}
	engine.Where("name=?", user1.Name).Desc("id").Get(&user10)
	fmt.Println(user10)
	fmt.Println("-------------------")

	// 根据条件查询多条数据
	var user11 []User
	engine.Table(&user).Where("passwd=?", "123456").And("age=?", 18).Find(&user11)
	fmt.Println(user11)

	// 获取记录总数
	var user12 = User{}
	counts, _ := engine.Count(&user12)
	fmt.Println(counts)

	// Iterate 和 Rows 根据条件遍历数据
	user13 := User{Passwd: "123456"}
	engine.Iterate(&user13, func(idx int, bean interface{}) error {
		user := bean.(*User)
		fmt.Println(user)
		return nil
	})

}
