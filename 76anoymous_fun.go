package main

import "fmt"

func main() {
	fmt.Println("\n\n ===>72 fun <====")

	a := 2
	fmt.Println(a)

	func() {
		fmt.Println("\n\n ===>76 anoymouse func <====")
	}()

	func(a, b int) int {
		fmt.Println("\n\n ===>76 anoymouse func <====")

		fmt.Println("a + b", a+b)
		return a + b
	}(1, 2)
}

func fun11(a int) int {
	fmt.Println("72 fun1")

	return a
}
