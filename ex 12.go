package main

import (
	"fmt"
	"strings"
)

func main() {
	var palavra string
	fmt.Println("Digite uma palavra:")
	fmt.Scan(&palavra)

	if isPalindromo(palavra) {
		fmt.Println("A palavra é um palíndromo.")
	} else {
		fmt.Println("A palavra não é um palíndromo.")
	}
}

func isPalindromo(palavra string) bool {
	palavra = strings.ToLower(palavra)
	for i := 0; i < len(palavra)/2; i++ {
		if palavra[i] != palavra[len(palavra)-1-i] {
			return false
		}
	}
	return true
}
