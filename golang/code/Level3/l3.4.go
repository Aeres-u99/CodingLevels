package main

import "fmt"

func main() {
    length := 0
    var temp int
    fmt.Println("Enter the number of input: ")
    fmt.Scanln(&length)
    fmt.Println("Enter the input: ")
    numbers := make([]int, length)
    // n[number] = count
    for i := 0; i < length; i++ {
        fmt.Scanln(&numbers[i])
    }
    for i := 0; i < length; i++ {
        fmt.Println(numbers[i])
    }
    fmt.Println("Reversing the array")
    for i := 0; i < ((length + 1) / 2); i++ {
        temp = numbers[i]
        numbers[i] = numbers[length - i - 1]
        numbers[length - i - 1] = temp
    }
    for i := 0; i < length; i++ {
        fmt.Println(numbers[i])
    }
}


/*
Index:   0 1 2 3 4 5 6 7 8 9

data:    0 1 2 3 4 5 6 7 8 9

reverse: 0 9 8 7 6 5 4 3 2 1



*/
