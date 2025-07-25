package main

import (
	"fmt"
	"strconv"
)

func main() {
	var input string
	var archSymbol string
	var countSymbols int = 0
	var resultArchive string
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println("Ошибка!")
	}

	for i := 0; i < len(input); i++ {

		// fmt.Println(input[i])
		// fmt.Println(archSymbol)
		// fmt.Println(i)
		// fmt.Println(len(input))
		if archSymbol == "" {
			archSymbol = string(input[i])
			countSymbols = 1
		} else if archSymbol != string(input[i]) {
			resultArchive += archSymbol
			resultArchive += strconv.Itoa(countSymbols)
			archSymbol = string(input[i])
			countSymbols = 1
		} else if archSymbol == string(input[i]) {
			countSymbols += 1
		} else {
			fmt.Println("Незапланированное поведение")
			return
		}

		if i == len(input)-1 {
			resultArchive += archSymbol
			resultArchive += strconv.Itoa(countSymbols)
		}

	}

	fmt.Println(resultArchive)
}
