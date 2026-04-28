/*
Домашняя работа 17.6.1 Попытка 1

Первая попытка через select не увенчалась успехом. "Случайный" выбор обработки потока внутри инструкции select не позволяет
на постоянной основе получать уверенный результат "в 2 раза больше". Бывают случаи запусков, когда близко к этому решению, но не всегда
и не со 100% точностью
*/
package main

import (
	"fmt"
	"time"
)

func main() {

	var cnt int = 1
	channel1 := make(chan int)
	channel2 := make(chan int)

	var channel1_count int
	var channel2_count int

	go func() {

		for i := range 10 {
			channel1 <- i
			channel1 <- i + 1
		}
		close(channel1)

	}()

	go func() {

		for i := range 10 {
			channel2 <- i
		}
		close(channel2)
	}()

	time.Sleep(1 * time.Second)

	fmt.Println("")
	for {

		fmt.Println(cnt, " попытка опроса каналов")
		cnt++

		if channel1 == nil && channel2 == nil {
			fmt.Println("Все каналы закрыты. Завершаем опрос каналов")
			break
		}

		select {
		case ch1value, is_open := <-channel1:
			fmt.Println("Получено сообщение из первого(1) канала. Значение: ", ch1value)
			channel1_count++
			if !is_open {
				fmt.Println("Закрыли канал 1")

				channel1 = nil
			}
			select {
			case ch1value_inner, is_open_inner := <-channel1:
				fmt.Println("Получено сообщение из первого(1) канала. Значение: ", ch1value_inner)
				channel1_count++
				if !is_open_inner {
					fmt.Println("Закрыли канал 1")

					channel1 = nil
				}
			default:
				fmt.Println("Нечего было получить во вложенном потоке канала 1")
			}
		case ch2value, is_open := <-channel2:
			fmt.Println("Получено сообщение из второго(2) канала. Значение: ", ch2value)
			channel2_count++
			if !is_open {
				fmt.Println("Закрыли канал 2")

				channel2 = nil
			}

		default:
			fmt.Println("Нечего было читать из каналов. Продолжим опрос. По какой-то причине произошла блокировка. Считаем это паникой - завершаем программу")
			return
		}

		fmt.Println("Количество сообшений из канала 1: ", channel1_count, "!!!!")
		fmt.Println("Количество сообшений из канала 2: ", channel2_count, "!!!!")

	}
	fmt.Println("== Количество сообшений из канала 1: ", channel1_count, "==")
	fmt.Println("== Количество сообшений из канала 2: ", channel2_count, "==")

}
