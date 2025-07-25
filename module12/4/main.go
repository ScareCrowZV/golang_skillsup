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

	ar = mergeSort(ar)

	fmt.Println(ar)
}

func mergeSort(ar []int) []int {

	divideAr := len(ar) / 2
	firstAr := mergeSort(ar[:divideAr])
	secondAr := mergeSort(ar[divideAr:])

	var resultAr []int
	for len(firstAr) > 0 && len(secondAr) > 0 {
		if firstAr[0] <= secondAr[0] {
			resultAr = append(resultAr, firstAr[0])
			firstAr = firstAr[1:]
		} else {
			resultAr = append(resultAr, secondAr[0])
			secondAr = secondAr[1:]
		}
	}
	resultAr = append(resultAr, firstAr...)
	resultAr = append(resultAr, secondAr...)

	return resultAr
}
