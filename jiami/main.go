package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	fmt.Println("ceshi -------------start")

	str := "www.topgoer.com"

	data := []byte(str)
	has := md5.Sum(data)
	md5str1 := fmt.Sprintf("%x", has)
	fmt.Println(md5str1)

	fmt.Println("ceshi -------------end")
}
