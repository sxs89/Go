package utils

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"gopkg.in/ini.v1"
	"os"
)

var DB *xorm.Engine

func init() {
	fmt.Println("初始化db中的init")

	// 加载 app.ini文件配置动态获取文件信息内容
	//config, er := ini.Load("gin_zhenghe/conf/app.ini")
	config, er := ini.Load("D:\\WEB\\Go\\gin_zhenghe\\conf\\app.ini")
	if er != nil {
		fmt.Printf("Fail to read file: %v", er)
		os.Exit(1)
	}
	ip := config.Section("mysql").Key("ip").String()
	database := config.Section("mysql").Key("database").String()
	user := config.Section("mysql").Key("user").String()
	password := config.Section("mysql").Key("password").String()
	port := config.Section("mysql").Key("port").String()
	sqlStr := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True", user, password, ip, port, database)

	// 连接数据库
	//sqlStr := "root:root@tcp(127.0.0.1:3306)/go_test?charset=utf8mb4"
	var err error
	DB, err = xorm.NewEngine("mysql", sqlStr)
	if err != nil {
		fmt.Println("连接数据库失败！", err)
		return
	}
	// 会在控制台打印执行的SQL
	DB.ShowSQL()
}
