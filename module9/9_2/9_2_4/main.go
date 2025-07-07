package main

import "fmt"

func main() {
	arr := [5]float64{1, 1, 1, 1, 1}
	changeArray(&arr)
	fmt.Println(arr)
	fmt.Println(len(arr))
}

func changeArray(arr *[5]float64) {
	arr[4] = 100
}
