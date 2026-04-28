/*
Домашняя работа 17.6.1 Попытка 2

Вторая попытка будет реализована с помощью контроля выборки из потоков
*/
package main

import (
	"fmt"
)

func main() {

	var cnt int = 1
	channel1 := make(chan int)
	channel2 := make(chan int)

	var ch1value int
	var ch2value int

	var channel1_is_open bool
	var channel2_is_open bool

	var channel1_count int
	var channel2_count int

	go func() {

		for i := range 20 {
			channel1 <- i
		}
		close(channel1)

	}()

	go func() {

		for i := range 20 {
			channel2 <- i
		}
		close(channel2)
	}()

	for channel1 != nil || channel2 != nil {

		fmt.Println(cnt, " попытка опроса каналов")
		cnt++

		if channel1 != nil {
			ch1value, channel1_is_open = <-channel1
			if !channel1_is_open {
				fmt.Println("Закрыли канал 1")
				channel1 = nil
			} else {
				fmt.Println("Получено сообщение из первого(1) канала. Значение: ", ch1value)
				channel1_count++
			}
		}
		if channel2 != nil {
			if channel2_count < channel1_count/2 || channel1 == nil {
				ch2value, channel2_is_open = <-channel2
				if !channel2_is_open {
					fmt.Println("Закрыли канал 2")
					channel2 = nil
				} else {
					fmt.Println("Получено сообщение из второго(2) канала. Значение: ", ch2value)
					channel2_count++
				}
			}
		}

		fmt.Println("Количество сообшений из канала 1: ", channel1_count)
		fmt.Println("Количество сообшений из канала 2: ", channel2_count)

	}

	fmt.Println("== Количество сообшений из канала 1: ", channel1_count, "==")
	fmt.Println("== Количество сообшений из канала 2: ", channel2_count, "==")

}
