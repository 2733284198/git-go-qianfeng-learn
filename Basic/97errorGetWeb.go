package main

import (
	"fmt"
	"net"
)

func main() {
	fmt.Println("\n ==>")

	addr, err := net.LookupHost("www.baidu.com")

	fmt.Println(err)
	fmt.Println(addr)
}
