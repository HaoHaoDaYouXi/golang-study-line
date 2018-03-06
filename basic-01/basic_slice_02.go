package main

import "fmt"

func main() {

	/* [1] 初始化数组，初始化数组中 {} 中的元素个数不能大于 [] 中的数字*/
	var arr1 = [5]float32{1.2,3.4,5.6,4.4,6.6} //必须指定数组大小

	// [2] 如果忽略 [] 中的数字不设置数组大小，Go 语言会根据元素的个数来设置数组的大小
	var arr2 = [...]int{1,3,5,7,9,11}

	/* [3] 数组元素可以通过索引（位置）来读取。格式为数组名后加中括号，中括号中为索引的值*/
	fmt.Println("-------打印数组1：",arr1[1])
	fmt.Println("-------打印数组2：",arr2[1])
	fmt.Printf("-------开始切片----------\n")

	/* [4] 切片的字面值和数组字面值很像，不过切片没有指定元素个数*/
	slice1 := []float32{1.2,3.4,5.6,4.4,6.6}

	/* 使用make函数创建切片 */
	// func make([]T, len, cap) []T : T 代表被创建的切片元素的类型。函数 make 接受一个类型、一个长度和一个可选的容量参数。调用 make 时，内部会分配一个数组，然后返回数组对应的切片。
	slice2 := make([]int,5,10)

	// 使用内置函数 len 和 cap 获取切片的长度和容量信息
	fmt.Printf("------获取切片 slice1 的长度和容量信息 len=%d cap=%d slice=%v\n",len(slice1),cap(slice1),slice1)
	fmt.Printf("------获取切片 slice2 的长度和容量信息 len=%d cap=%d slice=%v\n",len(slice2),cap(slice2),slice2)

	// 零值的切片类型变量为 nil。对于零值切片变量，len 和 cap 都将返回 0。
	b := []byte{'g', 'o', 'l', 'a', 'n', 'g'} // len=6 cap=6 slice=[103 111 108 97 110 103]
	slice3 := []int{1,2,3,4,5,6,7}
	fmt.Printf("------获取切片 b 的长度和容量信息 len=%d cap=%d slice=%v\n",len(b),cap(b),b)
	fmt.Printf("------获取切片 slice3 的长度和容量信息 len=%d cap=%d slice=%v\n",len(slice3),cap(slice3),slice3)
}

