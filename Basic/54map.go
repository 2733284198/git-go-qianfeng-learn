package main

import "fmt"

func main() {

	fmt.Println("\n ==>")
	// create  map
	var map1 map[int]string
	var map2 = make(map[int]string)
	var map3 = map[string]int{"Go": 98, "Python": 88, "Java": 77, "Html": 93}

	fmt.Println(map1)
	fmt.Println(map2)
	fmt.Println(map3)

	fmt.Println("\n ==>")

	fmt.Println(map1 == nil)
	fmt.Println(map2 == nil)
	fmt.Println(map3 == nil)

	//map1[3] = {100:"hello"};
	//map1[2] = "hello"

	fmt.Println("\n ==>")
	// todo:map如何赋值，
	map2[1] = "map1"
	map2[2] = "map2"
	map2[100] = "map100"

	fmt.Println(map2)
	//fmt.Println(map2[100])

	fmt.Println("\n ==>取值 ")
	// 值不存在
	value1, ok1 := map2[100]
	//value2, ok2 := map2[200]

	//fmt.Println(value1, ok1, map2[100])

	if ok1 == true {
		fmt.Println("v2值:", value1)
	} else {
		fmt.Println("操作的key不存在，获取到的是零值")
	}

	//	 map3
	fmt.Println("\n ==>mp3取值 ")
	map3["news"] = 900
	v3, ok3 := map3["news"]

	if ok3 == true {
		fmt.Println("v3值:", v3)
	} else {
		fmt.Println("v3操作的key不存在，获取到的是零值")
	}

	// 5 修改数据

	// 6 删除数据
	fmt.Println("\n ==>mp3 删除 ")

	fmt.Println("删除前", map3["news"])
	delete(map3, "news")
	fmt.Println("删除后", map3["news"])

}
