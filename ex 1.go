package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("Digite uma string:")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		output := strings.ToUpper(scanner.Text())
		fmt.Println("String convertida:", output)
	}
}
