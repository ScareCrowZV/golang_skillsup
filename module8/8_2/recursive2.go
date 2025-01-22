package main

import "fmt"

func main() {
	fmt.Println(recursive(5))
}

func recursive(number int) int {
	if number > 1 {
		return number * recursive(number-1)
	}
	return 1
}

/*
5
5 * (5-1)=120 = ?
4 * (4-1) = 24 = ?
3 * (3-1) = 6 = 18
2 * (2-1)=2 = 4
1 = 1
*/