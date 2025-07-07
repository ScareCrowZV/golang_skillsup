package main

import "fmt"

func main() {
	var arr = [2][4]int{
		[4]int{1, 2, 3, 4},
		[4]int{5, 6, 7, 8},
	}

	fmt.Println(arr)
	fmt.Println(arr[0][1])
	fmt.Println(arr[1][2])
}
