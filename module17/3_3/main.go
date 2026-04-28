// Перепишите приведённый выше пример со счётчиком из основного текста, но вместо примитивов
// из пакета atomic используйте условную переменную и попробуйте реализовать динамическую
// проверку достижения конечного значения счётчиком.

package main

import (
	"fmt"
	"sync"
)

const step int = 1
const iterationAmount int = 1000

func main() {

	var counter int = 0

	c := sync.NewCond(&sync.Mutex{})

	increment := func() {
		c.L.Lock()

		counter += step

		if counter == iterationAmount*step { // Я отдаю себе отчёт, что если какая-либо горутина потеряется, сломается или что-то ещё с ней произойдёт, и она не увеличит значение, то программа просто повиснет навсегда. Пока я не знаю что с этим делать. Можно поставить таймаут, можно обрабатывать панику в горутинах
			c.Signal()
		}

		c.L.Unlock()
	}

	for range iterationAmount {
		go increment()
	}

	c.L.Lock()
	c.Wait()
	c.L.Unlock()

	fmt.Println(counter)
}
