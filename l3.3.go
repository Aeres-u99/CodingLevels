package main

import "fmt"

func main() {
    length := 0
    fmt.Println("Enter the number of input: ")
    fmt.Scanln(&length)
    fmt.Println("Enter the input: ")
    numbers := make([]int, length)
    numHashMap := make(map[int]int)
    // n[number] = count
    for i := 0; i < length; i++ {
        fmt.Scanln(&numbers[i])
    }
    fmt.Println(numbers)
    for i := 0; i < length; i++ {
        numHashMap[numbers[i]] += 1

    }
    for key, value := range numHashMap {
        fmt.Printf("%d ==> %d\n", key, value)
    }
}
