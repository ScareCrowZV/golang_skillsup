package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup

	wg.Add(5)
	for i := range 5 {
		go func(i int) {
			for range 10 {
				fmt.Println("Мой порядковый номер: ", i)
			}
			wg.Done()
		}(i)
	}

	wg.Wait()
}
