package main

import "fmt"

func main() {
   var  ptr *int

   fmt.Printf("ptr 的值为 : %x\n", ptr  )

   // 空指针判断
   if ptr != nil {
		fmt.Printf("ptr 不是空指针")
   }

   if ptr == nil {
		fmt.Printf("ptr 是空指针")
	}
}