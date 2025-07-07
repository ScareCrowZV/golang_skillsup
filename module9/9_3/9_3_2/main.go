package main

import "fmt"

func main() {
	slice := make([]float64, 5)
	fmt.Println(slice, cap(slice), len(slice))
	slice = append(slice, 1)
	fmt.Println(slice, cap(slice), len(slice))
}
