package main

import "fmt"

func main() {
	sl := []float64{1, 2, 3, 4, 5}
	fmt.Println(sl, cap(sl), len(sl))
	changeSlice(sl)
	fmt.Println(sl, cap(sl), len(sl))

}

func changeSlice(sl []float64) {
	sl[4] = 100
}
