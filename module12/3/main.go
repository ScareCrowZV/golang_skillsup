package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ar := make([]int, 50)

	// ar[0] = 2
	// ar[1] = 1
	// ar[2] = 3
	// ar[3] = 4
	// ar[4] = 5

	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}

	insertionSort(ar)

	fmt.Println(ar)
}

func insertionSort(ar []int) {

	for i := 1; i < len(ar); i++ {
		for j := range ar[:i] {
			if ar[j] > ar[i] {
				insertValue := ar[i]
				ar[i] = ar[j]
				ar[j] = insertValue
			}
		}
	}
}
