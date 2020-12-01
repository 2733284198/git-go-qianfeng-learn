package main

import "fmt"

func main() {
	fmt.Println("\n\n ===>72 fun <====")

	a := 2
	fmt.Println(a)

	defer fun1(a)
	a++
	fmt.Println(a)

	//fun2()
}

func fun1(a int) int {
	fmt.Println("72 fun1")

	return a
}

func fun2() {
	fmt.Println("72 fun2")

}
