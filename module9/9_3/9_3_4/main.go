package main

import "fmt"

func main() {
	x := make([]float64, 5, 10)

	fmt.Println(x, cap(x), len(x))

	x = append(x, 1, 2, 3, 4, 5)

	fmt.Println(x, cap(x), len(x))

	x = append(x, 6)

	fmt.Println(x, cap(x), len(x))
}
