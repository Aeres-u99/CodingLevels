package main

import (
    "fmt"
    "bufio"
    "os"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Enter Your String: ")
    s, _ := reader.ReadString('\n')
    rev := Reverse(s)
    fmt.Println(rev)
}

func Reverse(s string) string {
    n := len(s)
    r := make([]rune, len(s))
    for _, rune := range s {
        n--
        r[n] = rune
    }
    return string(r[n:])
}
