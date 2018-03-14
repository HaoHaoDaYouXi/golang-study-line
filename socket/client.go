package main

import (
	"net"
	"fmt"
)

func connHandler(c net.Conn){
	// 当一个函数调用前有关键字 defer 时, 那么这个函数的执行会推迟到包含这个 defer 语句的函数即将返回前才执行
	defer c.Close()
}

func main() {
  	defer fmt.Println("我是最后执行的")
    fmt.Println("我是第一个")
    fmt.Println("我是第二个")
}