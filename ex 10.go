package main

import (
    "fmt"
    "sort"
)

func main() {
    var str1, str2 string
    fmt.Print("Digite a primeira string: ")
    fmt.Scanln(&str1)
    fmt.Print("Digite a segunda string: ")
    fmt.Scanln(&str2)

    str1Sorted := sortString(str1)
    str2Sorted := sortString(str2)

    if str1Sorted == str2Sorted {
        fmt.Printf("\"%s\" e \"%s\" são anagramas.\n", str1, str2)
    } else {
        fmt.Printf("\"%s\" e \"%s\" não são anagramas.\n", str1, str2)
    }
}

func sortString(str string) string {
    runes := []rune(str)
    sort.Slice(runes, func(i, j int) bool { return runes[i] < runes[j] })
   
    return string(runes)
}
