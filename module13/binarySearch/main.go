package main

func main() {

}

func findBinary(s []int, val int) int { // двоичный поиск в отсортированном слайсе
	start, end := 0, len(s) // задаём начальные значения
	for {
		middle := (start + end) / 2 // находим середину массива
		if s[middle] == val {       // элемент найден
			return middle
		}
		if middle == start { // мы не нашли элемент
			return -1
		}
		if s[middle] < val { // искомый элемент справа от середины
			start = middle
		}
		if s[middle] > val { // искомый элемент слева от середины
			end = middle
		}
	}
}

func findBinaryR(s []int, val, start, end int) int { // рекурсивная версия
	middle := (start + end) / 2 // находим середину массива
	if s[middle] == val {       // элемент найден
		return middle
	}
	if middle == start { // мы не нашли элемент
		return -1
	}
	if s[middle] < val { // искомый элемент справа от середины
		return findBinaryR(s, val, middle, end)
	}
	if s[middle] > val { // искомый элемент слева от середины
		return findBinaryR(s, val, start, middle)
	}
	return -1
}
