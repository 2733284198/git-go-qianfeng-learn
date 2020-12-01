package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("\n ==>60 string pack ")

	s1 := "hello world"
	s1sub := "he"
	s1sub2 := "hewo"

	s1contain := strings.Contains(s1, s1sub)
	if s1contain {
		fmt.Println("\n => s1包含s1sub ")
	} else {
		fmt.Println("\n =》 s1不包含s1sub ")
	}

	s2contain := strings.ContainsAny(s1, s1sub2)
	if s2contain {
		fmt.Println("\n => s1包含s1sub2 ", s1, s1sub2)
	} else {
		fmt.Println("\n =》 s1不包含s1sub2 ", s1, s1sub2)
	}

	//	 split
	fmt.Println("\n ==>60 string split === ")

	str1 := "s1,s2,s3,s4"
	str2 := strings.Split(str1, ",")

	str3join := strings.Join(str2, "-")

	fmt.Println(str2)
	fmt.Printf("%T, %s", str2, str2)

	fmt.Println(str3join)
	fmt.Printf("%T, %s", str3join, str3join)

	// up, lowwer
	fmt.Println("\n ==>60 string split === ")
	str1up := strings.ToUpper(str1)
	fmt.Println(str1up)

}
