package main

import "fmt"

// 定义Car 接口
type Car interface {
	NameGet() string
	Run(n int)
	Stop()
}

// 宝马结构体
type BMW struct {
	Name string
}

// 实现方法
func (this *BMW) NameGet() string {
	return this.Name
}

func (this *BMW) Run(n int) {
	fmt.Printf("BMW is running of num is %d \n", n)
}

func (this *BMW) Stop() {
	fmt.Printf("BMW is stop \n")
}

// 大奔 结构体
type Benz struct {
	Name string
}

func (this *Benz) NameGet() string {
	return this.Name
}

func (this *Benz) Run(n int) {
	fmt.Printf("Benz is running of num is %d \n", n)
}

func (this *Benz) Stop() {
	fmt.Printf("Benz is stop \n")
}

func (this *Benz) ChatUp() {
	fmt.Printf("ChatUp \n")
}

func main() {
	var car Car /* 声明 car 为 Car 类型 */
	fmt.Println(" car is = ", car)
	fmt.Println("-----------------------\r\n")

	var bmw BMW = BMW{Name: "宝马"}
	car = &bmw
	fmt.Println(car.NameGet()) //宝马
	car.Run(1)                 // BMW is running of num is 1
	car.Stop()                 // BMW is stop
	fmt.Println("-----------------------\r\n")

	benz := &Benz{Name: "大奔"}
	car = benz
	fmt.Println(car.NameGet()) //大奔
	car.Run(2)                 //Benz is running of num is 2
	car.Stop()                 //Benz is stop
}
