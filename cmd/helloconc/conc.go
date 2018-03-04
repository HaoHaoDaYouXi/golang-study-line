// Hello World 并发版
package main

import (
	"fmt"
	"time"
)

func main() {
	// 通道
	ch := make(chan string)
	// go starts a goroutine 其他语言通过异步IO来完成并发
	for i := 0; i < 5000; i++ {
		// 并发
		go printHello(i)
		time.Sleep(time.Millisecond)
	}

	//通信
	for {
		msg := <-ch
		fmt.Println(msg)
	}
}

func printHello(i int) {
	for {
		fmt.Printf("Hello World from goroutine %d\n", i)
	}
}
