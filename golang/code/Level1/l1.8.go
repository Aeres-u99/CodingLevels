package main

import "fmt"

func main() {
    var year int64
    fmt.Println("Enter the year: ")
    fmt.Scan(&year)
    if (year % 4 == 0) {
        if (year % 100 == 0) {
            fmt.Printf("%d is Not Leap year", year)
        } else {
            fmt.Printf("%d is Leap year", year)
        }
    } else {
        fmt.Printf("%d is not leap year", year)
    }
}
