package main

import "fmt"

func checkSliceIsSorted(a []int) bool {
	for i := 0; i < len(a); i++ {
		if i == 0 {
			continue
		}

		if a[i] < a[i-1] {
			return false
		}
	}

	return true
}

func main() {
	fmt.Println("Проверяем, отсортирован ли слайс по возрастанию")
	sl := []int{2, 1}
	fmt.Println(checkSliceIsSorted(sl))
}
