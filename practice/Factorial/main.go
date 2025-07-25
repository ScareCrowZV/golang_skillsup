package main

import "fmt"

func factorialIterate(num int) int {
	var result int = 1
	for i := 1; i <= num; i++ {
		result *= i
	}
	return result
}

func factorialRecursive(num int) int {
	if num < 2 {
		return 1
	}

	return num * factorialRecursive(num-1)

}

func main() {
	var input int

	fmt.Print("Введите факториал: ")
	_, err := fmt.Scanln(&input)
	if err != nil {
		fmt.Println(fmt.Sprintf("Не получилось прочитать строку: %s", err))

	}

	fmt.Println(factorialRecursive(input))
	fmt.Println(factorialIterate(input))
}
