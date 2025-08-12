package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайное значение от [-100;100]
	}

	fmt.Println("Исходный порядок")
	fmt.Println(ar)

	fmt.Println("По возрастанию")
	bubbleSort(ar)
	fmt.Println(ar)

	fmt.Println("По убыванию")
	bubbleSortReserved(ar)
	fmt.Println(ar)
}

func bubbleSort(ar []int) {
	var is_already_sort bool = true
	for i := 0; i < len(ar)-1; i++ {
		for j := 0; j < len(ar)-1; j++ {
			if ar[j] > ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
				is_already_sort = false
			}
		}
		if is_already_sort {
			fmt.Println("Перемещения объектов не было")
			return
		}
	}
}

func bubbleSortReserved(ar []int) {
	var is_already_sort bool = true
	for i := 0; i < len(ar)-1; i++ {
		for j := 0; j < len(ar)-1; j++ {
			if ar[j] < ar[j+1] {
				ar[j], ar[j+1] = ar[j+1], ar[j]
				is_already_sort = false
			}
		}
		if is_already_sort {
			fmt.Println("Перемещения объектов не было")
			return
		}
	}
}
