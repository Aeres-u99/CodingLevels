package main

import "fmt"

func main() {
	var number int64
	var sum int64
	var next int64
	var i int64
	sum = 0
	next = 1
	fmt.Print("Enter a number: ")
	fmt.Scan(&number)
	fmt.Printf("Printing %d digits of Fibonacci", number)
	for i = 0; i < number; i = i + 2 {

		fmt.Printf("%d %d ", sum, next)
		sum = next + sum
		next = sum + next
	}
}
