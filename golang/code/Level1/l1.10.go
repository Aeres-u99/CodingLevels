package main

import (
	"fmt"
)

func main() {
	var number int64
    var i int64
	fmt.Println("Enter a Number: ")
	fmt.Scan(&number)
	for i = number; i >= 0; i-- {
		fmt.Printf("%d\n", i)
	}
}
