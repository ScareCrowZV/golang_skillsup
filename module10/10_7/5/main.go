package main

import "fmt"

func main() {
	fmt.Println(trimNegative([]int{1, -1, 2, -0, 3, -1, 4, 5, 6, -2}))
}

func trimNegative(array []int) (cleanArray []int) {
	for _, value := range array {
		if value >= 0 {
			cleanArray = append(cleanArray, value)
		}
	}

	return cleanArray
}
