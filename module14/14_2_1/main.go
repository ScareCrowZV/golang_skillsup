package main

import "fmt"

func main() {
	fmt.Println(hashint64(5))
}

func hashint64(val int64) uint64 {
	return uint64(val % 1000)
}
