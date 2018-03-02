package main

import "fmt"

func main() {
	//使用 var 定义一个布尔类型并且初始化
	var flag bool = false
	fmt.Println("flag = ", flag)

	//int var
	var i, j, k int
	i = 1
	j = 2
	k = 3
	fmt.Printf("i=%d, j=%d, k=%d\n", i, j, k)

	//var rune and init go特有数据类型，专用于存储Unicode编码
	var r1, r2, r3 rune = 4, 5, 6
	fmt.Printf("r1=%d r3=%d r3=%d\n", r1, r2, r3)

	//var float32 and init
	var (
		f1 float32 = 2.1
		f2 float32 = 3.4
	)
	fmt.Printf("f1=%f f2=%f\n", f1, f2)

	//var string and init
	var str1 string
	str1 = "Tinywan"
	fmt.Println("Hello, ", str1)

	// :=
	str2 := "hello str2"
	fmt.Printf("str2 = %s \n", str2)
}
