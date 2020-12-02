package main

import "fmt"

func main() {

	//p()
	fmt.Println("\n ==>")

	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(a)

	fmt.Println("\n ==>")

	a1 := a[:5]
	fmt.Println(a1)

	fmt.Println("\n ==>")

	a2 := a[3:8]
	fmt.Println(a2)

	fmt.Println("\n ==>")

	a3 := a[:]
	fmt.Println(a3)

}
