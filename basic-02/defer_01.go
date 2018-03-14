package main

import "fmt"

func main() {
	i := 1
	defer fmt.Println("Deferred print:",i)
	i++
	fmt.Println("Normal print:", i)
}