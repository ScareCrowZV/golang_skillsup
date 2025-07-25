package main

import "fmt"

func main() {
	var intArr = []int{1, 2, 4, 5, 6}
	var stringArr = []string{"one", "two", "three"}

	// элементом слайса может быть любой тип, в т.ч. слайс
	var matrix = [][]int{
		[]int{0, 1},
		[]int{3, 4},
	}

	fmt.Println(intArr, stringArr, matrix)

	//Доступ к элементу слайса можно получить по индексу, например:
	firstElement := intArr[0] // firstElement == 1
	fmt.Println(firstElement)

	//Можно перебрать все элементы слайса с помощью цикла for:
	for i, value := range stringArr {
		fmt.Println(i, value)
	}
	//Можно изменить значение по индексу после инициализации массива:
	fmt.Println(matrix) // Output: [[0 1] [3 4]]

	matrix[1] = []int{5, 6}
	fmt.Println(matrix) // Output: [[0 1] [5 6]]

	// 	Длину слайса можно узнать с помощью функции len(), а вместимость с помощью cap().

	// Инициализировать слайс с заранее заданной длиной и вместимостью можно с помощью функции make():

	ar := make([]int, 2, 4)
	fmt.Println(len(ar), cap(ar), ar) // Output: 2 4 [0 0]

	sar := make([]string, 2, 4)
	fmt.Println(len(sar), cap(sar), sar) // Output: 2 4 [ ] ← two empty strings

	par := make([]*struct{}, 2, 4)
	fmt.Println(len(par), cap(par), par) // Output: 2 4 [<nil> <nil>]
}
