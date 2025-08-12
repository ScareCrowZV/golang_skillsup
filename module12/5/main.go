package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ar := make([]int, 5)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}
	fmt.Println(ar)
	quickSort(ar)
	fmt.Println(ar)
}

func quickSort(ar []int) {

	if len(ar) < 2 {
		return
	}

	middleElement := len(ar) / 2
	lastElement := len(ar) - 1
	left := 0

	ar[middleElement], ar[lastElement] = ar[lastElement], ar[middleElement]

	for i, v := range ar {
		if v < ar[lastElement] {
			ar[left], ar[i] = ar[i], ar[left]
			left++
		}
	}
	ar[left], ar[lastElement] = ar[lastElement], ar[left]

	quickSort(ar[left+1:])
	quickSort(ar[:left])
}
