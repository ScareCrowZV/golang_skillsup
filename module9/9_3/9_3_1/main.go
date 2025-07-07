package main

import "fmt"

func main() {
	/*
		Объявление слайса
	*/
	var x []float64

	/*
		Определение слайса на определённое количество элементов, ёмкость при этом равняется длине
	*/
	x2 := make([]float64, 5)

	/*
		Определение слайса с 3 элементами и ёмкостью 5
	*/
	x3 := make([]float64, 3, 5)

	fmt.Println(x, cap(x), len(x))
	fmt.Println(x2, cap(x2), len(x2))
	fmt.Println(x3, cap(x3), len(x3))
}
