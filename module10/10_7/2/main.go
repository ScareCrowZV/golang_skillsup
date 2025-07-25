package main

import (
	"fmt"
)

func main() {
	fmt.Println(findMax([]int{2, 1, 4}))

}

func findMax(array []int) (max int, err error) {
	if len(array) == 0 {
		return 0, fmt.Errorf("could not found max in empty slice")
	}

	max = array[0]
	for _, val := range array[1:] {
		if val > max {
			max = val
		}
	}

	return max, nil
}
