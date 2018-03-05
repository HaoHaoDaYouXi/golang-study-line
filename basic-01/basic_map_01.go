package main

import "fmt"

func main() {

        // 语法：var map_variable map[key_data_type]value_data_type
        var map_hostip map[string]string

        /* 创建集合 */
        map_hostip = make(map[string]string)

        map_hostip["host31"] = "192.168.32.131"
        map_hostip["host32"] = "192.168.32.132"
        map_hostip["host33"] = "192.168.32.133"
        map_hostip["host34"] = "192.168.32.134"

        fmt.Println("map_hostip = ", map_hostip)
        fmt.Println("map_hostip[host31] = ", map_hostip["host31"])

}