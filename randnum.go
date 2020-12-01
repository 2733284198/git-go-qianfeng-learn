package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	num1 := rand.Int()
	fmt.Println("\n")

	fmt.Println(num1)

	fmt.Println("\n")

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 10; i++ {
		//num2 = rand.Intn(10)
		fmt.Println("==>", i, rand.Intn(10))
	}

	fmt.Println("\n")
	//rand.Seed(3)
	num3 := rand.Intn(10)
	fmt.Println(num3)

}
