package main

import "fmt"

func main() {

	// [1] 声明map 语法：var map_variable map[key_data_type]value_data_type
	var map1 map[string]string

	/* 使用make函数创建一个非nil的map(创建集合)，nil map不能赋值 */
	map1 = make(map[string]string)

	/* 给已声明的map赋值 */
	map1["name"] = "Tinywan"
	map1["age"] = "24"

	/* [2] 直接创建map */
	map2 := make(map[string]string)
	/* 赋值 */
	map2["name"] = "Tinyaiai"
	map2["age"] = "25"

	/* [3] 初始化 + 赋值一体化 */
	map3 := map[string]string{
		"name": "ShaoBo Wan",
		"age":  "26",
	}

	/* 使用 key 输出 map 值 */
	fmt.Printf("-------遍历 map1---------\n")
	for name := range map1 {
		fmt.Println("map1 of", name, "is", map1[name])
	}

	// 遍历 map2
	fmt.Printf("-------遍历 map2---------\n")
	for k, v := range map2 {
		fmt.Println(k, v)
	}

	// 遍历 map3
	fmt.Printf("-------遍历 map3---------\n")
	for k, v := range map3 {
		fmt.Println(k, v)
	}

	// 查找键值是否存在
	fmt.Printf("-------查找键值是否存在---------\n")
	if v, ok := map1["name"]; ok {
		fmt.Println("存在", v)
	} else {
		fmt.Println("Key Not Found")
	}

	/* 删除元素 */
	delete(map3, "age")
	fmt.Printf("--------删除后的 map3--------\n")

	// 遍历 map3
	for k, v := range map3 {
		fmt.Println(k, v)
	}
}
