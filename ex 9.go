package main

import (
    "fmt"
    "strings"
)

func main() {
    var str string
    fmt.Print("Digite uma string: ")
    fmt.Scanln(&str)

    var oldChar, newChar string
    fmt.Print("Digite a letra a ser substituída: ")
    fmt.Scanln(&oldChar)
    fmt.Print("Digite a letra substituta: ")
    fmt.Scanln(&newChar)

    replaced := strings.Replace(str, oldChar, newChar, -1)

    fmt.Printf("A string \"%s\" com \"%s\" substituído por \"%s\" é \"%s\".\n", str, oldChar, newChar, replaced)
}
