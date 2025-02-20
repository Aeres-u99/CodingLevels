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
        number /= 10
        count = count + 1
    }
    fmt.Printf("Number of digits: %d", count)
}
