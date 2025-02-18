package main

import "fmt"

func main() {
    var number int64
    var i int64
    fmt.Println("Enter a number: ")
    fmt.Scan(&number)
    for i = 1; i <= 10; i++ {
        fmt.Printf("%d X %d = %d\n", number, i, number * i)
    }
}
