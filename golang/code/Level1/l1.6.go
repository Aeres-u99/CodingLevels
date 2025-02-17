package main

import (
    "fmt"
)

func main() {
    var number int64
    fmt.Println("Enter a number: ")
    fmt.Scan(&number)
    if (number > 0) {
        fmt.Println("Positive")
    } else if (number == 0) {
        fmt.Println("Zero")
    } else {
        fmt.Println("Negative Number")
    }
}
