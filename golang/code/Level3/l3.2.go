package main

import "fmt"

func main() {
    length := 0
    var max int
    var secondmax int
    fmt.Println("Enter the number of input: ")
    fmt.Scanln(&length)
    fmt.Println("Enter the input: ")
    numbers := make([]int, length)
    for i := 0; i < length; i++ {
        fmt.Scanln(&numbers[i])
    }
    fmt.Println(numbers)
    max = numbers[0]
    secondmax = numbers[0]
    for i := 0; i < length; i++ {
        if numbers[i] >= secondmax {
            secondmax = numbers[i]
        }

        if numbers[i] >= max {
            secondmax = max
            max = numbers[i]
        }

    }
    fmt.Printf("Max: %d\n", max)
    fmt.Printf("Second Max: %d\n", secondmax)
}
