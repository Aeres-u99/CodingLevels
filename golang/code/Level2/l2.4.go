package main

import (
	"fmt"
	"math"
)

func isPrime(n uint64) bool {

	var i uint64
	var sqrt uint64
	sqrt = uint64(math.Sqrt(float64(n)))
	for i = 2; i <= sqrt; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var number uint64
	fmt.Println("Enter a number: ")
	fmt.Scan(&number)
	fmt.Printf("is Number %d Prime? %t", number, isPrime(number))
}
