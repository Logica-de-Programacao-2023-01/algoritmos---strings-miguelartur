package main

import (
    "fmt"
)

func main() {
    var str string
    fmt.Print("Digite uma string: ")
    fmt.Scanln(&str)

    count := make(map[rune]int)
    for _, char := range str {
        count[char]++
    }

    fmt.Print("Letras únicas: ")
    for char, freq := range count {
        if freq == 1 {
            fmt.Printf("%c ", char)
        }
    }
    fmt.Println()
}
