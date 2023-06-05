package main

import (
    "fmt"
    "strings"
)

func main() {
    var str string
    fmt.Print("Digite uma string: ")
    fmt.Scanln(&str)

    words := strings.Fields(str)
    count := len(words)

    fmt.Printf("A string \"%s\" tem %d palavra(s).\n", str, count)
}
