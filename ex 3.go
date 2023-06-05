package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    fmt.Print("Digite uma string: ")
    scanner.Scan()
    input := scanner.Text()

    fmt.Print("Digite o caractere a ser substituído: ")
    scanner.Scan()
    oldChar := scanner.Text()

    fmt.Print("Digite o novo caractere: ")
    scanner.Scan()
    newChar := scanner.Text()

    output := strings.ReplaceAll(input, oldChar, newChar)

    fmt.Println("String modificada:", output)
}
