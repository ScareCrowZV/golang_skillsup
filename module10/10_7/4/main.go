package main

import "fmt"

func main() {
	// fmt.Println(findMostOftenRepeated([]int{1, 2, 2, 2, 3, 3, 3}))
	// fmt.Println(findMostOftenRepeatedOptimized([]int{1, 2, 2, 2, 3, 3, 3}))
	fmt.Println(findMostOftenRepeatedWithMap([]int{1, 1, 1, 1, 1, 2, 2, 2, 3, 3, 3, 4, 4, 4, 4, 4, 4, 4, 4, 4, 4}))
}

func findMostOftenRepeated(array []int) (mostOften int, err error) {

	if len(array) == 0 {
		return 0, fmt.Errorf("could not found repeated numbers in empty slice")
	}

	var maxIndex, maxCount = 0, 0
	for i, number := range array {
		currentCount := 0
		for _, numberToCompare := range array {
			if number == numberToCompare {
				currentCount++
			}
		}

		if currentCount > maxCount {
			maxIndex = i
			maxCount = currentCount
		}
	}

	return array[maxIndex], nil
}

func findMostOftenRepeatedOptimized(array []int) (mostOften int, err error) {

	if len(array) == 0 {
		return 0, fmt.Errorf("could not found repeated numbers in empty slice")
	}

	var maxIndex, maxCount = 0, 0
	for i, number := range array {
		currentCount := 0
		for _, numberToCompare := range array[i:] {
			if number == numberToCompare {
				currentCount++
			}
		}

		if currentCount > maxCount {
			maxIndex = i
			maxCount = currentCount
		}
	}

	return array[maxIndex], nil
}

func findMostOftenRepeatedWithMap(array []int) (mostOften int, err error) {

	r := make(map[int]int)

	for _, value := range array {
		r[value]++
	}

	var count int
	for key, value := range r {
		if count < value {
			mostOften = key
			count = value
		}
	}

	return mostOften, nil
}
