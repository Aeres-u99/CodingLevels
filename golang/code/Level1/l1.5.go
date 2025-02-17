package main

import (
	"fmt"
)

func main() {
	var number1 int64
	var number2 int64
	var number3 int64
	var largest int64
	fmt.Println("Enter a Number: ")
	fmt.Scan(&number1)
	fmt.Println("Enter a Number: ")
	fmt.Scan(&number2)
	fmt.Println("Enter a Number: ")
	fmt.Scan(&number3)
	if number1 > number2 {
		if number1 > number3 {
			largest = number1
		} else {
			largest = number3
		}
	} else if number2 > number3 {
		largest = number2
	} else {
		largest = number3
	}
	fmt.Printf("%d is the Largest", largest)
}
