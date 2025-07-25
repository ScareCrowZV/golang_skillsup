package main

import "fmt"

func main() {

	var counter int = 0
	for i := 2; counter < 20; i++ {
		var countModOfNumbers int = 0
		for j := 2; j < i; j++ {
			if i%j == 0 {
				countModOfNumbers++
			}
		}
		if countModOfNumbers == 0 {
			counter++
			fmt.Println(i)
		}
	}

}
