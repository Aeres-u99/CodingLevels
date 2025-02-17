package main

import (
	"fmt"
)

func main() {
	var number int64
    var i int64
    var sum int64
	fmt.Println("Enter a Number: ")
	fmt.Scan(&number)
    sum = 0
	for i = 0; i <= number; i++ {
        sum += i
	}
    fmt.Printf("%d\n", sum)
}
