package main

import (
    "fmt"
)

func main() {
    var str string
    fmt.Print("Digite uma string: ")
    fmt.Scanln(&str)

    reversed := ""
    for i := len(str)-1; i >= 0; i-- {
        reversed += string(str[i])
    }

    fmt.Printf("A string \"%s\" invertida é \"%s\".\n", str, reversed)
}
