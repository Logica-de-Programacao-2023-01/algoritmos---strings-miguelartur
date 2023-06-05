package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Digite uma string:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		output := strings.ReplaceAll(scanner.Text(), " ", "")
		fmt.Println("String sem espaços em branco:", output)
	}
}
