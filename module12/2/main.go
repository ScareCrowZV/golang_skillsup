package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ar := make([]int, 50)
	// ar[0] = 63
	// ar[1] = -11
	// ar[2] = 53
	// ar[3] = 53
	// ar[4] = -32

	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайное значение от [-100;100]
	}
	fmt.Println("Исходный порядок")
	fmt.Println(ar)

	// fmt.Println("Сортированный слева на право")
	// selectionSort(ar)
	// fmt.Println(ar)

	// fmt.Println("Сортированный Справа на лево")
	// selectionSortMax(ar)
	// fmt.Println(ar)

	fmt.Println("Сортированный Двунаправленная")
	selectionSortDouble(ar)
	fmt.Println(ar)

}

func selectionSort(ar []int) {
	var minElementIndex int

	for i := 0; i < len(ar); i++ {

		minElementIndex = i

		for j := i + 1; j < len(ar); j++ {
			if ar[minElementIndex] > ar[j] {
				minElementIndex = j
			}
		}
		if ar[i] > ar[minElementIndex] {

			ar[i], ar[minElementIndex] = ar[minElementIndex], ar[i]
		}
	}
}

func selectionSortMax(ar []int) {
	var maxElementIndex int

	for i := len(ar) - 1; i > 0; i-- {

		maxElementIndex = i

		for j := 0; j < len(ar[:i]); j++ {
			if ar[maxElementIndex] < ar[j] {
				maxElementIndex = j
			}
		}
		if ar[i] < ar[maxElementIndex] {
			ar[i], ar[maxElementIndex] = ar[maxElementIndex], ar[i]
		}
	}
}

func selectionSortDouble(ar []int) {
	var minElementIndex int
	var maxElementIndex int

	for i := range ar {

		minElementIndex = i
		maxElementIndex = i
		i_eval := i + 1

		if i > len(ar)-i_eval {
			return
		}

		for j := i; j <= len(ar[i:len(ar)-i_eval]); j++ {

			if ar[j] < ar[minElementIndex] {
				minElementIndex = j
			}
			if ar[j] > ar[maxElementIndex] {
				maxElementIndex = j
			}

		}

		if ar[i] > ar[minElementIndex] {
			ar[i], ar[minElementIndex] = ar[minElementIndex], ar[i]
		}

		if ar[len(ar)-i_eval] < ar[maxElementIndex] {
			ar[len(ar)-i_eval], ar[maxElementIndex] = ar[maxElementIndex], ar[len(ar)-i_eval]
		}

	}
}
