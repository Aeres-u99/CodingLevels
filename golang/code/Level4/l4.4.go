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
    output := RemoveSpaces(input)
    fmt.Printf("%s", output)
}

func RemoveSpaces(s string) string {
    n := len(s)
    r := make([]rune, len(s))
    for _, rune := range s {
        if rune != ' ' {
        n--
        r[n] = rune
        }
    }
    return string(r[n:])
}
