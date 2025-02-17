package main

import (
	"fmt"
)

func main() {
	var number1 int64
	var number2 int64
	fmt.Println("Enter a Number: ")
	fmt.Scan(&number1)
	fmt.Println("Enter another Number: ")
	fmt.Scan(&number2)
	if number1 >= number2 {
		fmt.Printf("%d is greater than %d", number1, number2)
	} else {
		fmt.Printf("%d is greater than %d", number2, number1)
	}
}
