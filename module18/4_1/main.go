package main

import (
	"fmt"
	"time"
)

const (
	countOfSemaphors int = 10
)

// Semaphore — структура двоичного семафора
type Semaphore struct {
	// Семафор — абстрактный тип данных,
	// в нашем случае в основе его лежит канал
	sem chan int
	// Время ожидания основных операций с семафором, чтобы не
	// блокировать
	// операции с ним навечно (необязательное требование, зависит от
	// нужд программы)
	timeout time.Duration
}

// Acquire — метод захвата семафора
func (s *Semaphore) Acquire() error {
	select {
	case s.sem <- 0:
		return nil
	case <-time.After(s.timeout):
		return fmt.Errorf("не удалось захватить семафор")
	}
}

// Release — метод освобождения семафора
func (s *Semaphore) Release() error {
	select {
	case <-s.sem:
		return nil
	case <-time.After(s.timeout):
		return fmt.Errorf("не удалось освободить семафор")
	}
}

// NewSemaphore — функция создания семафора
func NewSemaphore(timeout time.Duration) *Semaphore {
	return &Semaphore{
		sem:     make(chan int, countOfSemaphors),
		timeout: timeout,
	}
}

func main() {

	// Создали новый семафор
	s := NewSemaphore(time.Second * 5) // Чем больше таймаут тем меньше ошибок захвата будет

	// Клиентский код, который обращается к общему ресурсу
	for i := range countOfSemaphors + 50 /* специальная константа, которая покажет, что превышение попыток обращение к семафору вызовет ошибку захвата или освобождения. Хорошо бы ещё иметь состояние семафора, что бы знать, можем сейчас захватить или будем ждать. Можно сделать так же через select + timeout*/ {

		go func(v int) {

			defer func() {
				if err := s.Release(); err != nil {
					fmt.Println(err, v)
					return
				}
			}()
			if err := s.Acquire(); err != nil {
				fmt.Println(err, v)
				return
			}
			fmt.Println("test", v)
			time.Sleep(time.Second * 1) // исскуственное замедление работы, для демонстрации ошибок при попытке захвата семафора другими
		}(i)
	}

	// Ждём пока горутины доработают
	time.Sleep(time.Second * 10)
}
