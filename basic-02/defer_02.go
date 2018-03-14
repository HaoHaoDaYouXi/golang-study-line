package main

import "fmt"

func f1() (r int) {
    r = 1
    defer func() {
        r++
        fmt.Println(" r value = ",r)
    }()
    r = 2
    return
}

func main() {
    f1()
}