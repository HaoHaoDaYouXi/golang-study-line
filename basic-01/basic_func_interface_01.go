package main

import "fmt"

// 定义结构体
type person struct {
        id      int
        name    string
        country string
}

// 定义接口
type interface_person interface {
        introduction()
}

//实现接口方法 使用指针访问
func (this *person) introduction() {
        fmt.Println("My name is : ", this.name)
}

func main() {

        var Tinywan person
        Tinywan.id = 13669361192
        Tinywan.name = "ShaoBo Wan"
        Tinywan.country = "China"

        fmt.Println("Tinywan = ", Tinywan)

        Tinywan.introduction()
}