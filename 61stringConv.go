package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Println("\n ==>61 stringConv pack ")

	//s1 := "hello world"
	//s1sub := "he"
	//s1sub2 := "hewo"

	strt1 := "true"
	str1bool, err := strconv.ParseBool(strt1)

	if err != nil {
		fmt.Printf("%T,%b", str1bool, str1bool)
	} else {
		fmt.Printf("%T,%s", str1bool, str1bool)
	}

	// strToInt
	fmt.Println("\n ==>61 stringConv pack ")

	str1 := "100"
	str2, err := strconv.ParseInt(str1, 10, 2)

	if err != nil {
		fmt.Printf("%T,%T", str1, str2)
	} else {
		fmt.Println("err")
		//fmt.Printf("%T,%T", str1, str2)
	}

	fmt.Println("\n\n  ==>61 stringConv.itoa pack ")
	inta := 42
	//strinta, err = strconv.Itoa(inta)
	strinta := strconv.Itoa(inta)

	if err != nil {
		fmt.Printf("%T,%T", str1, strinta)
	}

}
