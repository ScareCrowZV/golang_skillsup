package main

import "fmt"

func main() {

	channel1 := make(chan int)

	go func() {
		for i := 1; i < 101; i++ {
			channel1 <- i

		}
		close(channel1)
	}()

	for channel1 != nil {
		v, is_open := <-channel1
		if is_open {
			fmt.Println(v)
		} else {
			channel1 = nil
		}
	}

}
