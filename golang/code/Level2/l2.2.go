package main

import "fmt"

func main() {
	var number int64
	var fact int64
	var i int64
	fmt.Println("Enter a number:")
	fmt.Scan(&number)
	if (number == 0) {
		fmt.Println("1")
	} else if (number == 1) {
		fmt.Println("1")
	} else {
		fact = 1
		for i=1; i<=number; i++ {
			fact = fact * i
		}
		fmt.Println(fact)
	}
}

