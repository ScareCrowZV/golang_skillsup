package main

import "fmt"

func main() {

	var a1, b1 float64 = 3, 5
	var a2, b2 float64 = 1, 7
	var x, y float64

	if a1 == a2 {
		fmt.Println("Никогда не пересекаются!")
		return
	}

	x = (b2 - b1) / (a1 - a2)
	y = a1*x + b1
	fmt.Printf("x: %f, y: %f\n", x, y)

}

// y=ax+b
// y=3x+5
// y=x+7

// y-y=3x+5-(x+7)
// 0=2x-2
// 2x=-2
// x=-1
