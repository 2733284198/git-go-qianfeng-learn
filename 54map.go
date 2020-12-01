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

	// todo:map如何赋值，

}
