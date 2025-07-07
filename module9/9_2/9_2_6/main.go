package main

import "fmt"

func main() {
	var arr = [2][2][1][2]int{
		[2][1][2]int{
			[1][2]int{
				[2]int{1, 2},
			},
			[1][2]int{
				[2]int{3, 4},
			},
		},
		[2][1][2]int{
			[1][2]int{
				[2]int{5, 6},
			},
			[1][2]int{
				[2]int{7, 8},
			},
		},
	}

	fmt.Println(arr)
	fmt.Println(arr[1][1][0][0])
}
