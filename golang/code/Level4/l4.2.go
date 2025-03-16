package main

import (
    "fmt"
    "bufio"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Enter Your String: ")
    s, _ := reader.ReadString('\n')
    input := strings.TrimRight(s, "\n")
    output := PallindromeCheck(input)
    fmt.Printf("%t", output)
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

func PallindromeCheck(s string) bool {
    rev := Reverse(s)
    if rev == s {
        return true
    } else {
        return false
    }
}
