package main

import "fmt"

func main() {
	var x [5]int
	x[4] = 100
	fmt.Println(x)
	fmt.Println(x[4])

	var xn [5]float64
	xn[0] = 98
	xn[1] = 93
	xn[2] = 77
	xn[3] = 82
	xn[4] = 83

	var total float64 = 0
	for i := 0; i < 5; i++ {
		total += xn[i]
	}
	fmt.Println(total / 5)

	var xn2 [5]float64
	xn2[0] = 98
	xn2[1] = 93
	xn2[2] = 77
	xn2[3] = 82
	xn2[4] = 83

	var total2 float64 = 0
	for i := 0; i < 5; i++ {
		total2 += xn[i]
	}
	fmt.Println(total2 / float64(len(xn2)))
}
