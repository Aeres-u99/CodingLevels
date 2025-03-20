package main

import "fmt"

func main() {
    var l1 int
    var rotatefactor int
    var storage int
    fmt.Println("Insert Length of Array: ")
    fmt.Scan(&l1)
    input := make([]int, l1)
    fmt.Println("Enter elements of array:")
    for i:=0; i < l1; i++ {
        fmt.Scan(&input[i])
    }
    fmt.Println("By how much should we rotate to the right? ")
    fmt.Scan(&rotatefactor)
    if rotatefactor > l1 {
        rotatefactor = rotatefactor % l1
    } else {
        for i := 0; i < rotatefactor; i++ {
            /* Swap the element to be rotated with the new index */
            /* New index = (currentindex + length) % rotatefactor */
            storage = input[0]
            for j := 1; j < l1; j++ {
                input[j-1] = input[j] 
                fmt.Print("Debug: ")
                fmt.Println(input)
            }
            input[l1-1] = storage
        }
    }
    fmt.Println(input)
}
