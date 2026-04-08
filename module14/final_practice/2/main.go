package main

import (
	"fmt"
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano()) // необходимо для того, чтобы рандом был похож на рандомный
}

func main() {
	// ar := make([]int, 1000)
	// for i := range ar {
	// 	ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	// }

	// ar = []int{-17, 34, -6, -28, -70, -44, -32, 85, -22, -8}
	// ar = []int{0, 1, 2, 3, 4, 5}
	// ar = []int{9, 7, 4, 1, 3, 5}
	// ar = []int{0}
	// ar = []int{}
	// ar = []int{1, 1}
	// ar = []int{3, 2, 1}
	// ar = []int{5, 15, 2, 13, 7, 16, 10, 2}
	ar := []int{1, 9, 7, 4, 6, 2, 1, 13, 22, -3, 12, 76}

	fmt.Println("До сортировки:", ar)

	quickSort(ar)

	fmt.Println("После сортировки:", ar)
}

func quickSort(ar []int) {
	if len(ar) < 2 {
		return
	}

	if len(ar) == 2 {
		fmt.Println("Два элемента")
		if ar[0] > ar[1] {
			ar[0], ar[1] = ar[1], ar[0]
		}
		return
	}

	var left int = 0
	var referenceElement int = len(ar) / 2
	var lastElement int = len(ar) - 1

	ar[referenceElement], ar[lastElement] = ar[lastElement], ar[referenceElement]

	for i := 0; i < len(ar); i++ {
		if ar[i] < ar[lastElement] {
			if ar[left] != ar[i] {
				ar[left], ar[i] = ar[i], ar[left]
			}
			left++
		}
	}
	ar[left], ar[lastElement] = ar[lastElement], ar[left]

	fmt.Println("Левый:", ar[left+1:])
	fmt.Println("Правый:", ar[:left])

	if len(ar[left+1:]) > 1 {
		quickSort(ar[left+1:])
	}
	if len(ar[:left]) > 1 {
		quickSort(ar[:left])
	}

}
