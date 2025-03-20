package main

import "fmt"

func main() {
    var l1 int
    var l2 int
    fmt.Println("Insert Length of First Array: ")
    fmt.Scan(&l1)
    input1 := make([]int, l1)
    fmt.Println("Enter elements of array:")
    for i:=0; i < l1; i++ {
        fmt.Scan(&input1[i])
    }

    fmt.Println("Insert Length of First Array: ")
    fmt.Scan(&l2)
    input2 := make([]int, l2)
    fmt.Println("Enter elements of array:")
    for i:=0; i < l2; i++ {
        fmt.Scan(&input2[i])
    }
    intersection := make(map[int]int, 0)
    for elements := 0; elements < l1; elements++ {
        intersection[input1[elements]] += 1
    }
    for elements := 0; elements < l2; elements++ {
        intersection[input2[elements]] += 1
    }
    for key, _ := range intersection {
        if intersection[key] > 1 {
            fmt.Printf("%d ", key)
        }
    }
}
