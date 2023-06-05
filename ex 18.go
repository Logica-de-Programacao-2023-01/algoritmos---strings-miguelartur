package main

import (
    "fmt"
    "unicode"
)

func main() {
    var str string
    fmt.Print("Digite uma string: ")
    fmt.Scanln(&str)

    for _, char := range str {
        if !unicode.IsDigit(char) {
            fmt.Println("A string não contém apenas números.")
            return
        }
    }

    fmt.Println("A string contém apenas números.")
}
