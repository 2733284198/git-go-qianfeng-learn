package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("\n ==>60 string pack ")

	s1 := "hello world"
	s1sub := "he"

	s1contain := strings.Contains(s1, s1sub)
	if s1contain {
		fmt.Println("\ns1包含s1sub ")
	} else {
		fmt.Println("\n =》 s1不包含s1sub ")
	}
}
