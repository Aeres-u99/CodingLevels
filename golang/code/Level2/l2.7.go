package main

import (
    "fmt"
    "os"
)
func reverse(n int64) int64 {
    var reversed int64
    for n > 0 {
        remainder := n % 10
        reversed = reversed * 10 + remainder
        n = n / 10
    }
    return reversed
}

func main(){
    var number int64
    var reversed int64
    fmt.Print("Enter a number: ")
    fmt.Scan(&number)
    if number == 0 {
        fmt.Print("0")
        os.Exit(0)
    }
    reversed = reverse(number)
    fmt.Printf("Reversed: %d", reversed)
}
