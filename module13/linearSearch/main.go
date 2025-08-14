package main

import "fmt"

func main() {
	data := []int{1, 2, 3, 4, 5}
	fmt.Println(findMaxInSlice(data))
	fmt.Println(findInSlice(data, 5))
}

func findMaxInSlice(s []int) (int, int) { // возвращает индекс и значение максимального элемента
	var index, val int    // значения по умолчанию 0
	for k, v := range s { // проход по слайсу golang-style
		if v > val { // в случае, если текущий элемент больше запомненного
			index, val = k, v // обновляем запомненные элементы
		}
	}
	return index, val // возвращаем найденные значения
}

func findInSlice(s []int, value int) int { // возвращаем только индекс, значение уже известно
	for i := 0; i < len(s); i++ { // проход по слайсу C-style
		if s[i] == value { //
			return i // мы нашли искомый элемент и можем завершать работу функции
		}
	}
	return -1 // возвращаем -1, если элемент не найден
}
