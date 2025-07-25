package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	(rand.NewSource(time.Now().UnixNano())) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	ar := make([]int, 50)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}

	insertionSort(ar)

	fmt.Println(ar)
}

func insertionSort(ar []int) {
	for i := 1; i < len(ar); i++ {
		insertionValue := ar[i]
		j := i - 1
		for j >= 0 && ar[j] > insertionValue {
			ar[j+1] = ar[j]
			j--
		}
		ar[j+1] = insertionValue
	}
}
