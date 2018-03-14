package main

import "fmt"

type User struct {
	Name string
	Age int
	sex string
 }

 type Lover struct {
	User
	sex time.Time
	int
	Age int
}

/*
1.方法的访问控制也是通过大小写控制 
2.init函数是通过传入指针实现，这样改变struct字段值，因为是值类型 
3.小写包外无法引用
*/
func (this *User) init(name string, age int, sex string) {
	this.Name = name
	this.Age = age
	this.sex = sex
}

func (this User) GetName() string {
	return this.Name
}

func main() {
	var user User /* 声明 user 为 User 类型 */
	user.init("nick", 18, "man")
	//(&user).init("nick", 18, "man")
	name := user.GetName()
	fmt.Println(name)
}

