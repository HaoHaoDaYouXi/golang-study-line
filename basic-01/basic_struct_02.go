package main

import "fmt"

type Persons struct {
	name string
	age  int
	blog string
	ID   int
}

func main() {
	var Tinywan Persons /* 声明 Tinywan 为 Persons 类型 */

	/* Tinywan 个人信息描述 */
	Tinywan.name = "ShaoBo Wan"
	Tinywan.age = 24
	Tinywan.blog = "www.tinywan.com"
	Tinywan.ID = 756684177

	/* 打印 Book1 信息 */
	fmt.Printf("Tinywan name : %s\n", Tinywan.name)
	fmt.Printf("Tinywan age : %d\n", Tinywan.age)
	fmt.Printf("Tinywan blog : %s\n", Tinywan.blog)
	fmt.Printf("Tinywan ID : %d\n", Tinywan.ID)
}
