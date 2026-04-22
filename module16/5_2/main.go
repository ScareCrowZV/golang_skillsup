// Напишите программу, которая будет каждую секунду увеличивать значение переменной на единицу
// и каждую ⅕ секунды выводить в консоль значение этой переменной. Программа должна выполняться n секунд и после этого закрываться.
// n вводит пользователь после запуска программы.

package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	fmt.Println("Введите количество секунд, которое будет работать программа")

	var durationSeconds int
	mutexControlValue := sync.Mutex{}
	var controlValue int

	if _, err := fmt.Scanln(&durationSeconds); err != nil {
		fmt.Println("error", err)
		return
	}

	go func() {

		for durationSeconds != 0 {

			mutexControlValue.Lock()
			controlValue++
			mutexControlValue.Unlock()
			time.Sleep(time.Second)

		}
	}()

	go func() {

		for durationSeconds != 0 {
			mutexControlValue.Lock()
			fmt.Println("Контрольное значение:", controlValue)
			mutexControlValue.Unlock()
			time.Sleep(time.Millisecond * 200)

		}

	}()

	for durationSeconds != 0 {

		fmt.Println(controlValue)
		durationSeconds--
		time.Sleep(time.Second)

	}

}
