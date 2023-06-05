package main

import (
    "fmt"
    "unicode"
)

func main() {
    var str string
    fmt.Print("Digite uma string: ")
    fmt.Scanln(&str)

    var count int
    for i, char := range str {
        if i == 0 {
            if !unicode.IsUpper(char) {
                fmt.Println("A string não está em camelCase.")
                return
            }
            count++
        } else if unicode.IsUpper(char) {
            count++
        }
    }

    fmt.Printf("A string está em camelCase e possui %d palavras.\n", count)
}
