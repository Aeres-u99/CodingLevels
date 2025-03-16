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
    output := CountVowelsAndConsonents(input)
    for key, value := range output {
        fmt.Printf("%s => %v\n", key, value)
    }
}

func CountVowelsAndConsonents(s string) map[string]int {
    result := make(map[string]int)
    var vowels int
    var consonents int
    for _, rune := range s {
        if ('a' <= rune && rune <= 'z') || ('A' <= rune && rune <= 'Z' ) {
            switch rune {
                case
                    'a',
                    'A',
                    'e',
                    'E',
                    'I',
                    'i',
                    'o',
                    'O',
                    'u',
                    'U':
                    vowels ++
                default:
                    consonents ++
            }
        }
    }
    result["Vowels"] = vowels
    result["Consonents"] = consonents
    return result
}
