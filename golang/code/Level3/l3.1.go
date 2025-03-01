package main

import "fmt"

func main() {
    length := 0
    var max int
    fmt.Println("Enter the number of input: ")
    fmt.Scanln(&length)
    fmt.Println("Enter the input: ")
    numbers := make([]int, length)
    for i := 0; i < length; i++ {
        fmt.Scanln(&numbers[i])
    }
    fmt.Println(numbers)
    max = numbers[0]
    for i := 0; i < length; i++ {
        if numbers[i] >= max {
            max = numbers[i]
        }

    }
    fmt.Printf("Max: %d", max)
}
