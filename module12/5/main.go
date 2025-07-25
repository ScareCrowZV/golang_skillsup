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

	ar = quickSort(ar)

	fmt.Println(ar)
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
