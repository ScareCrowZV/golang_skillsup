// Напишите код, в котором несколько горутин увеличивают значение целочисленного счётчика и синхронизируют свою работу через канал.
// Нужно предусмотреть возможность настройки количества используемых горутин и конечного значения счётчика, до которого его следует увеличивать.

// Попробуйте реализовать счётчик с элементами ООП (в виде структуры и методов структуры).
// Попробуйте реализовать динамическую проверку достижения счётчиком нужного значения.

package main

import (
	"fmt"
	"sync"
)

const goroutinesAmount = 1000
const maxValueAmount = 100000

type CounterS struct {
	channel chan int
	value   int
}

func (c *CounterS) addValue(v int, wg *sync.WaitGroup) {
	defer wg.Done()
	c.channel <- v
}

/*
Запускает горутину, которая обрабатывает сообщения из канала и добавляет к значению счётчика
Вызывать нужно только один раз. По хорошему нужно делать проверку, что горутина занимающаяся записью уже существует,
либо создавать её в конструкторе
*/
func (c *CounterS) startValueApplyer(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for newValueFromChannel := range c.channel {
			if c.value+newValueFromChannel <= maxValueAmount {
				c.value += newValueFromChannel
			} else {
				//Тут хочется послать сообщение основной горутине, что бы она перестала генерить горутины, которые увеличивают счётчик
				// или послать им сообщение напрямую через ещё один канал, что бы они завершили свою работу
				// Мы просто дождёмся пока они доделают своё дело, хотя это и избыточно, т.к. не можем просто закрыть канал. Это приведёт к ошибке.
				// Кстати интересный вопрос, а может ли горутина-писатель в канал, проверить, открыт ли он ?
				continue
			}

		}
	}()
}

func NewCounterS(initValue int) *CounterS {

	return &CounterS{
		value:   initValue,
		channel: make(chan int),
	}

}

func main() {

	var wgMain sync.WaitGroup
	var wgAdder sync.WaitGroup
	var cntr CounterS = *NewCounterS(0)                    // Создаём "объёкт" счётчика. Инициализируем стартовое значение и канал синхронизации
	var addedValue int = maxValueAmount / goroutinesAmount // Вычисляем какое значение нужно добавлять, что бы достигнуть порога

	cntr.startValueApplyer(&wgMain) // Запускаем горутину, которая будет отвечать за суммирование счётчика

	wgAdder.Add(goroutinesAmount) // Добавляем горутины, которые будут приходить и пробовать увеличивать счётчик
	for range goroutinesAmount {
		go cntr.addValue(addedValue, &wgAdder)
	}
	wgAdder.Wait()      // Пришлось пойти на хитрость и сделать второй WaitGroup, вначале дождусь пока доработают горутины, которые знаимаются увеличением счётчика
	close(cntr.channel) // Потом могу безопасно закрыть канал
	wgMain.Wait()       // И уже дождаться пока доработает основная горутина

	fmt.Println(cntr.value)

}
