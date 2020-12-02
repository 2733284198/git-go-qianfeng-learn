package main

import (
	"fmt"
	"time"
)

func init() {
	fmt.Println(" \n ===> init")
}

func main() {
	fmt.Println(time.Now())
	fmt.Println(time.Now().UnixNano())
	fmt.Println(time.Minute)
	fmt.Println(time.Hour)

}
