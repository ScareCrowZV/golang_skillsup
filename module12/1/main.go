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
	ar := make([]int, 10)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайное значение от [-100;100]
	}

	bubbleSort(ar)
	// bubbleSort(ar) // Запуск второй сортировки для отладки выхода из функции

	fmt.Println(ar)

	bubbleSortReversed(ar)
	fmt.Println(ar)
}

func bubbleSort(ar []int) {
	var is_sorting int = 0
	for i := 0; i < len(ar)-1; i++ {
		for j := 1; j < len(ar)-i; j++ {
			if ar[j-1] > ar[j] {
				ar[j-1], ar[j] = ar[j], ar[j-1]
				is_sorting = 1
			}

		}

		if is_sorting == 0 {
			// fmt.Printf("Сортировка не произошла %d\n", i)
			return
		}
	}

}

func bubbleSortReversed(ar []int) {
	var is_sorting int = 0
	for i := 0; i < len(ar)-1; i++ {
		for j := 1; j < len(ar)-i; j++ {
			if ar[j-1] < ar[j] {
				ar[j-1], ar[j] = ar[j], ar[j-1]
				is_sorting = 1
			}

		}

		if is_sorting == 0 {
			// fmt.Printf("Сортировка не произошла %d\n", i)
			return
		}
	}

}
