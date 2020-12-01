package main

import "fmt"

func main() {

	fmt.Println("\n\n ===>")

	/*var arra := [2][2] {
		{[11,12], [21,22] },
		{[11,12], [21,22] },
	}*/

	fmt.Println("\n\n range loop ===>")
	var aint = [2][2]int{
		{11, 12},
		{21, 22},
	}

	for k, v := range aint {
		fmt.Println(k, "=>", v)
	}

	fmt.Println("\n\n for loop ===>")

	for i := 0; i < len(aint); i++ {
		for j := 0; j < len(aint[i]); j++ {
			fmt.Println(aint[i][j], "\t")
		}

		fmt.Println()
	}

	fmt.Println("\n\n ===>")

	var bstr = [2][2]string{
		{"aa", "ab"},
		{"ba", "bb"},
	}

	for k1, v1 := range bstr {
		fmt.Println(k1, "=>", v1)
	}

	/*for i,v := range arr1{
		fmt.Println(i, '=', v)
	}*/

}
