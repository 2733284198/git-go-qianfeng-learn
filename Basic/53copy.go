package main

import "fmt"

func main() {

	fmt.Println("\n ==>")

	s1 := []int{1, 2, 3, 4}
	s2 := s1[:]

	s3 := make([]int, 0)
	for i := 0; i < len(s1); i++ {
		s3 = append(s3, s1[i])
	}

	s1[0] = 100

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)

	fmt.Println("\n ==>")
	s4 := []int{7, 8, 9}
	copy(s3, s4)

	fmt.Println("s1 ", s1)
	fmt.Println("s2", s2)
	fmt.Println("s3", s3)
	fmt.Println("s4", s4)

	fmt.Println("\n ==>")

}
