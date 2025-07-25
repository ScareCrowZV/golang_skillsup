package main

import "fmt"

func main() {
	var n int = 101

	a, b := 0, n
	for b > 0 {
		a *= 10
		a += b % 10
		b /= 10
		fmt.Println("-----")
		fmt.Println(a)
		fmt.Println(b)
	}

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(n)
}
