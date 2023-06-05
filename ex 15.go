package main

import "fmt"

func main() {
	var palavra string
	fmt.Println("Digite uma palavra:")
	fmt.Scan(&palavra)

	palavra = strings.ReplaceAll(palavra, "a", "*")
	palavra = strings.ReplaceAll(palavra, "e", "*")
	palavra = strings.ReplaceAll(palavra, "i", "*")
	palavra = strings.ReplaceAll(palavra, "o", "*")
	palavra = strings.ReplaceAll(palavra, "u", "*")

	fmt.Println(palavra)
}
