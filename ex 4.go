package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Digite a primeira string: ")
	scanner.Scan()
	str1 := scanner.Text()

	fmt.Print("Digite a segunda string: ")
	scanner.Scan()
	str2 := scanner.Text()

	if str1 == str2 {
		fmt.Println("As strings são iguais.")
	} else {
		fmt.Println("As strings são diferentes.")
	}
}
