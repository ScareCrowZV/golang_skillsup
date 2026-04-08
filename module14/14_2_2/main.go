package main

import "fmt"

func main() {
	fmt.Println(hashstr("test"))
	fmt.Println(hashstr("tset"))
	fmt.Println(hashstr("booombaaar"))
}

func hashstr(val string) (res uint64) {

	for i, v := range val {
		res += (uint64(v) * uint64(i))
	}

	return res % 1000
}
