package main

import "fmt"

func main() {
        fmt.Println("function begins ... ")
        c := make(chan bool)
        go func() {
                fmt.Println("func has been called.")
                /* 内置的 close 函数关闭通道
                1、如果一个通道被关闭了，我们就不能往这个通道里发送数据了，如果发送的话，会引起 painc异常,
                2、但是，我们还可以接收通道里的数据，如果通道里没有数据的话，接收的数据是 nil 。
                */
                close(c)
        }()
        <-c // 通道连接信息
        fmt.Println("Completed.")
}