package main

import "fmt"

func main() {

	fmt.Println("\n ==>")

	var map3 = map[string]int{"Go": 98, "Python": 88, "Java": 77, "Html": 93}

	var map1 = map[int]string{
		1: "str1",
		2: "str2",
		3: "str3",
	}

	fmt.Println(map1)
	fmt.Println(map3)

	//	 打印
	fmt.Println("\n ==>")
	for k1, v1 := range map1 {
		fmt.Println(k1, v1)
	}

}
