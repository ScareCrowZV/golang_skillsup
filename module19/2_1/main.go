package main

import (
	"fmt"
	"sync"
	"time"
)

// Число сообщений от источника
const messagesAmountPerGoroutine int = 500000

type ProcessingChannels struct {
	finish_processing chan int
}

func (pc *ProcessingChannels) startProcessing() {
	pc.finish_processing = make(chan int)
}

func (pc *ProcessingChannels) finishProcessing() {
	close(pc.finish_processing)
}

// Функция разуплотнения каналов
func (pc *ProcessingChannels) demultiplexingFunc(dataSourceChan chan int, amount int) []chan int {
	var output = make([]chan int, amount)
	// Обратите внимание: вышеприведённая команда инициализирует слайс,
	// но не инициализирует каждый его элемент. Каждый элемент будет
	// представлять так называемое нулевое значение
	// для данного типа.
	// Так как тип у нас канала — ссылочный, то все элементы
	// будут
	// равны nil
	for i := range output {
		output[i] = make(chan int)
	}
	go func() {
		var wg sync.WaitGroup
		wg.Add(1)
		go func() {
			defer wg.Done()
			// При поступлении сообщения в канал-источник
			// отправляем его в каждый из каналов-потребителей
			for v := range dataSourceChan {
				for _, c := range output {
					c <- v
				}
			}
		}()
		wg.Wait()
		/*
			После завершения посылки сообщений в основной канал-источник данных
			Дожидаемся одну секунду, что бы горутины функции multiplexingFunc() успели считать из каналалов
			затем закрываем сигнальный канал вызовом pc.finishProcessing()
		*/

	}()
	return output
}

// Функция уплотнения каналов
func (pc *ProcessingChannels) multiplexingFunc(channels ...chan int) <-chan int {
	var wg sync.WaitGroup
	// Общий канал, в который будут попадать сообщения от всех
	// источников
	// Именно его мы и вернём из этой функции для употребления
	// внешним кодом
	multiplexedChan := make(chan int)
	multiplex := func(c <-chan int) {
		defer wg.Done()
		for i := range c {
			/*
				Если поступило сообщение из одного из
				каналов-источников,
				перенаправляем его в общий канал
				Так же добавляем проверку, что если не успеем считать до завершения закрытия сигнального канала
			*/
			select {
			case <-pc.finish_processing:
				return
			case multiplexedChan <- i:
			}
		}

	}
	wg.Add(len(channels))
	for _, c := range channels {
		go multiplex(c)
	}

	// Запускаем горутину, которая закроет канал после того,
	// как завершится работа всех горутин-отправителей в мультиплексированный канал
	go func() {
		wg.Wait()
		close(multiplexedChan)

	}()
	return multiplexedChan
}

func main() {

	pc := new(ProcessingChannels)
	pc.startProcessing()

	// Горутина — источник данных
	// Функция создаёт свой собственный канал
	// и посылает в него пять сообщений
	startDataSource := func() chan int {
		c := make(chan int)
		go func() {
			var wg sync.WaitGroup
			wg.Add(1)
			go func() {
				defer wg.Done()
				for i := 1; i <= messagesAmountPerGoroutine; i++ {
					c <- i
				}
			}()
			wg.Wait()
			close(c)
		}()
		return c
	}
	// Запускаем источник данных и уплотняем каналы
	var consumers []chan int = pc.demultiplexingFunc(startDataSource(), 5)
	c := pc.multiplexingFunc(consumers...)

	// Централизованно получаем сообщения от всех нужных нам
	// источников
	// данных
	for {
		select {
		case data := <-c:
			fmt.Println(data)
		case <-pc.finish_processing:
			return
		case <-time.Tick(1 * time.Second):
			pc.finishProcessing()
		}
	}

}
