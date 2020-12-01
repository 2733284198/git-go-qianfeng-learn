package main

import "fmt"

/*func p()  {
	fmt.Println("\n ===>")
}*/

func main() {

	//p()

	// 另外数组定义
	//var a = []int{1, 2, 3} // 输出结果不同

	var a = []int{1, 2, 3}

	fmt.Println(len(a), cap(a))
	fmt.Printf("%p\n", a)

	b := append(a, 5, 6)
	fmt.Println(len(b), cap(b))
	fmt.Printf("%p\n", b)

	b = append(b, 7, 8)
	fmt.Println(len(b), cap(b))
	fmt.Printf("%p\n", b)

	//p()
}
