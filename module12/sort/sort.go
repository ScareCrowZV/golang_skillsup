package sort

import (
	"math/rand"
	"time"
)

func init() {
	(rand.NewSource(time.Now().UnixNano())) // необходимо для того, чтобы рандом был похож на рандомный
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

func selectionSort(ar []int) {
	for i := 0; i < len(ar); i++ {
		var minIndex = i
		for j := i + 1; j < len(ar); j++ {
			if ar[j] < ar[minIndex] {
				minIndex = j
			}
		}

		ar[i], ar[minIndex] = ar[minIndex], ar[i]
	}
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

func mergeSort(ar []int) []int {
	if len(ar) < 2 {
		return ar
	}

	halfArray := len(ar) / 2

	sortingArray := make([]int, 0, len(ar))
	left, right := mergeSort(ar[:halfArray]), mergeSort(ar[halfArray:])

	var i, j = 0, 0
	for i < len(left) && j < len(right) {
		if left[i] > right[j] {
			sortingArray = append(sortingArray, right[j])
			j++
		} else {
			sortingArray = append(sortingArray, left[i])
			i++
		}
	}

	sortingArray = append(sortingArray, left[i:]...)
	sortingArray = append(sortingArray, right[j:]...)

	return sortingArray
}

func quickSort(ar []int) []int {
	if len(ar) < 2 {
		return ar
	}
	val := ar[0]
	var left, right []int
	for _, value := range ar[1:] {
		if value <= val {
			left = append(left, value)
		} else {
			right = append(right, value)
		}
	}

	var rs []int
	rs = append(quickSort(left), val)
	rs = append(rs, quickSort(right)...)
	return rs
}
