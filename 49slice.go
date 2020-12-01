package main

import "fmt"

func main() {
	p()

	// 另外数组定义
	var a = [4]int{1, 2, 3, 4}
	fmt.Println(a)

	var s1 []int
	fmt.Println(s1)

	s2 := []int{1, 2, 3, 4} // 变长
	fmt.Println(s2)

	p()

	s4 := make([]int, 0, 5)
	fmt.Println(s4)

	s4 = append(s4, 1, 2)
	fmt.Println(s4)

	p()
}
