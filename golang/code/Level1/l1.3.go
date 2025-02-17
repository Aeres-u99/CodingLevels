package main

import (
	"fmt"
)

func main() {
	var number int64
	var val bool
	fmt.Println("Enter the Number: ")
	fmt.Scan(&number)
	val = (number & 1) != 0
	if val {
		fmt.Println("Number is odd")
	} else {
		fmt.Println("Number is Even")
	}
}
