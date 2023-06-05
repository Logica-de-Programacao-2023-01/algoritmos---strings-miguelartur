package main

import (
    "fmt"
    "unicode"
)

func main() {
    var str string
    fmt.Print("Digite uma string: ")
    fmt.Scanln(&str)

    hasNumber := false
    for _, char := range str {
        if unicode.IsDigit(char) {
            hasNumber = true
            break
        }
    }

    if hasNumber {
        fmt.Printf("A string \"%s\" contém pelo menos um número.\n", str)
    } else {
        fmt.Printf("A string \"%s\" não contém números.\n", str)
    }
}
