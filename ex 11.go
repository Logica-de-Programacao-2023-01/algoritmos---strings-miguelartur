package main

import "fmt"

func main() {
    var str string
    fmt.Print("Digite uma string: ")
    fmt.Scanln(&str)

    result := ""

    for _, char := range str {
        if !isVowel(char) {
            result += string(char)
        }
    }

    fmt.Printf("String sem vogais: %s\n", result)
}

func isVowel(char rune) bool {
    switch char {
    case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
        return true
    default:
        return false
    }
}
