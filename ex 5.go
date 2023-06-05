package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string

	fmt.Print("Digite um número em ponto flutuante: ")
	fmt.Scan(&input)

	if _, err := strconv.ParseFloat(input, 64); err == nil {
		fmt.Println("É um número em ponto flutuante válido.")
	} else {
		fmt.Println("Não é um número em ponto flutuante válido.")
	}
}
