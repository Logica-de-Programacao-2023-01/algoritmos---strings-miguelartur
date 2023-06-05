package main

import (
	"fmt"
	"strings"
)

func main() {
	var str, padrao string
	fmt.Print("Digite uma string: ")
	fmt.Scanln(&str)
	fmt.Print("Digite um padrão: ")
	fmt.Scanln(&padrao)

	posicoes := make([]int, 0)

	for i := 0; i < len(str); i++ {
		indice := strings.Index(str[i:], padrao)
		if indice == -1 {
			break
		}
		posicoes = append(posicoes, i+indice)
		i += indice + len(padrao) - 1
	}

	if len(posicoes) == 0 {
		fmt.Println("Padrão não encontrado na string.")
	} else {
		fmt.Println("Padrão encontrado nas seguintes posições:", posicoes)
	}
}
