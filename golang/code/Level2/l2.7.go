package main

import (
    "fmt"
)

func main() {
    var number uint64
    var count uint64
    fmt.Print("Enter a number: ")
    fmt.Scan(&number)
    count = 0
    for number != 0 {
        count += (number % 10)
        number /= 10
    }
    fmt.Printf("sum of digits: %d", count)
}
