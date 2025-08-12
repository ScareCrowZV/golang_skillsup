package main

import (
	"fmt"
	"math/rand"
)

func main() {
	ar := make([]int, 1000)
	for i := range ar {
		ar[i] = rand.Intn(200) - 100 // ограничиваем случайно значение от [-100;100]
	}

	fmt.Println(ar)
	ar = mergeSort(ar)

	fmt.Println(ar)
}

func mergeSort(ar []int) (resultArr []int) {
	var temporaryArr1 []int
	var temporaryArr2 []int

	if len(ar) > 2 {
		middleElement := len(ar) / 2
		temporaryArr1 = mergeSort(ar[middleElement:])
		temporaryArr2 = mergeSort(ar[:middleElement])

		for len(temporaryArr1) > 0 && len(temporaryArr2) > 0 {
			if temporaryArr1[0] > temporaryArr2[0] {
				resultArr = append(resultArr, temporaryArr2[0])
				temporaryArr2 = temporaryArr2[1:]
			} else {
				resultArr = append(resultArr, temporaryArr1[0])
				temporaryArr1 = temporaryArr1[1:]
			}

		}

		resultArr = append(resultArr, temporaryArr1...)
		resultArr = append(resultArr, temporaryArr2...)
	}

	if len(ar) == 2 {
		resultArr = append(resultArr, ar...)
		if resultArr[0] > resultArr[1] {
			resultArr[0], resultArr[1] = resultArr[1], resultArr[0]
		}
	}

	if len(ar) == 1 {
		resultArr = append(resultArr, ar[0])
	}

	fmt.Println("Результат прохода", resultArr)
	return resultArr
}
