package main

import "fmt"

// struct定义
type Person struct {
	name   string
	age  int
}

type Student struct {
	Person // 匿名字段，默认Student 就包含了Person的所有字段，这就是 Go 的嵌入方式
	speciality  string //专业
}

func main() {
	student := Student{Person{"ShaoBo Wan",24,},"QinHua"}
	fmt.Printf(" %v",student.Person.age)
}

