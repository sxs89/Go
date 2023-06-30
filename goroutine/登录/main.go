package main

import "fmt"

type User struct {
	Username string
	Password string
}

var users = []User{
	{"user1", "password1"},
	{"user2", "password2"},
	{"user3", "password3"},
	{"user4", "password4"},
}

func main() {

	// 模拟用户输入数据
	username := "user2"
	password := "password2"

	if login(username, password) {
		fmt.Println("登录成功！")
	} else {
		fmt.Println("登录失败！")
	}
}

func login(username string, password string) bool {
	// 遍历用户列表，查找匹配用户信息
	for _, user := range users {
		if user.Username == username && user.Password == password {
			return true
		}
	}
	return false
}
