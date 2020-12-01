package main

import "fmt"

func main() {
	fmt.Println("\n\n ===>")

	var num1 int
	num1 = 100

	num1 = 200
	fmt.Println(num1)

	// 创建数组
	//var arr1[4] int
	var arr1 [4]int

	arr1[0] = 10
	arr1[1] = 11
	arr1[2] = 12
	arr1[3] = 13

	fmt.Println("数组的值: ", arr1)
	fmt.Println("数组的值[0]: ", arr1[0])
	fmt.Println("长度: ", len(arr1))
	fmt.Println("容量: ", cap(arr1))

	fmt.Println("\n\n ===>")

	str1 := "str11"
	fmt.Println(str1)

	// 另外数组定义
	var a = [4]int{1, 2, 3, 4}
	fmt.Println(a)

	var bstr = [3]string{"rose", "王二狗", "ruby"}
	fmt.Println(bstr)
	fmt.Println(bstr[1])

	fmt.Println("\n\n ===>")
	//for i,v := range bstr{
	for i, v := range arr1 {
		fmt.Println(i, '=', v)
	}

	fmt.Println("\n\n ===>")
	//	 数组的类型
	fmt.Printf("%T\n", arr1)
	fmt.Printf("%T\n", bstr)

}
