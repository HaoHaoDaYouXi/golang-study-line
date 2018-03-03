// hello 开发版
package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/",func(writer http.ResponseWriter,request *http.Request){
			fmt.Fprintln(writer,"Hello World!")
		})
	http.ListenAndServe(":8888",nil)		
}