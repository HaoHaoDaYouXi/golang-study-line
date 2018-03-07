package main

import "fmt"

// 交换
func swap(a, b int) (int, int){
	return b, a
}

// 传值
func add(a int) (int){
	a = a + 1
	return a
}

//传指针
func addPointer(a *int) (*int){
	// 地址+1
	*a = *a + 1
	return a
}

func main(){
	a :=10
	b :=20
	a, b = swap(a,b)
	fmt.Printf("a=%d b=%d \n",a,b) // a=20 b=10

	add(a)
	fmt.Printf("a=%d\n",a) // a=20

	addPointer(&a)
	fmt.Printf("a=%d\n",a) // a=21
}