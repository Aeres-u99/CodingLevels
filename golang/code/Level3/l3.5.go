package main

import "fmt"

func main() {
    l1 := 0
    l2 := 0
    fmt.Println("Enter the number of input for Array 1: ")
    fmt.Scanln(&l1)
    arr1 := make([]int, l1)
    fmt.Println("Enter the input for Array 1: ")
    for i := 0; i < l1; i++ {
        fmt.Scanln(&arr1[i])
    }
    fmt.Println("Enter the number of input for Array 2: ")
    fmt.Scanln(&l2)
    arr2 := make([]int, l2)
    fmt.Println("Enter the input for Array 2: ")
    for i := 0; i < l2; i++ {
        fmt.Scanln(&arr2[i])
    }
    arr1 = Sort(arr1)
    arr2 = Sort(arr2)
    arr3 := SoftMerge(arr1, arr2)
    fmt.Println(arr3)
}

func SoftMerge(arr1, arr2 []int) []int {
    result := make([]int, len(arr1)+len(arr2))
    i, j, k := 0, 0, 0

    for i < len(arr1) && j < len(arr2){
        if arr1[i] < arr2[j] {
            result[k] = arr1[i]
            i++
        } else {
            result[k] = arr2[j]
            j++
        }
        k++
    }
    for i < len(arr1) {
        result[k] = arr1[i]
        i++
        k++
    }
    for j < len(arr2) {
        result[k] = arr2[j]
        j++
        k++
    }

    return result
}

func Sort(arr[] int) []int {
    for i := 0; i < len(arr); i++ {
        for j := len(arr) - i - 1; i < j; j-- {
            if arr[i] > arr[j] {
                arr[i], arr[j] = arr[j], arr[i]
            }
        }
    }
    return arr
}
