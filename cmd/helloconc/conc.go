// Hello World 并发版
package main

import (
	"time"
	"fmt"
)

func main () {
	// go starts a goroutine
	for i:=0; i<5000; i++ {
		go printHello(i)
		time.Sleep(time.Millisecond)
	}
}

func printHello(i int){
	for {
		fmt.Printf("Hello World from goroutine %d\n",i)
	}
}