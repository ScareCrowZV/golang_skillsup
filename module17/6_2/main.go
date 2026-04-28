package main

import (
	"fmt"
	"time"
)

func main() {
	var channel1, channel2 chan int // Неинициализированные каналы приведут к блокировке

	var ticker *time.Ticker = time.NewTicker(time.Second * 1)
	var t time.Time

	for {
		t = <-ticker.C
		select {
		case <-channel1:
			fmt.Println("Сообщение из 1 канала")
		case <-channel2:
			fmt.Println("Сообщение из 2 канала")
		default:
			fmt.Println("Текущее время: ", t.Format("15:04:05"))
		}

	}

}
