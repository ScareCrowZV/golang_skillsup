package main

import "fmt"

func main() {
	fmt.Println(trimLessAverage([]int{1, 2, 3, 1, 4, 6, 8, 3, 5, 7, 2, 5, 3, 1, 3, 7, 9, 2}))
}

func trimLessAverage(array []int) (resultArray []int) {

	var sum int

	var avg int
	for _, value := range array {
		sum += value
	}

	avg = sum / len(array)

	for _, value := range array {
		if value > avg {
			resultArray = append(resultArray, value)
		}
	}

	return resultArray

}
