package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMaxNegative([]int{2, 1, 1}))

}

func findMaxNegative(array []int) (searchVal int, err error) {

	searchVal = array[0]
	for _, val := range array[1:] {
		if val < searchVal {
			searchVal = val
		}
	}

	if searchVal >= 0 {
		return 0, fmt.Errorf("could not found negative number")
	}

	return searchVal, nil
}
