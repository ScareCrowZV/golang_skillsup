package main

import "fmt"

func hashstr(val string) (res uint64) {

	for i, v := range val {
		i += 1 // Если указывать число, то первый элемент всегда 0 и в этом случае хэш всегда возвращается 0
		res += (uint64(v) * uint64(i))
	}

	return res % 1000
}

func requestMapSize() (sizeMap uint64, err error) {

	if _, err := fmt.Scanln(&sizeMap); err != nil {
		fmt.Println("error", err)
		return 0, fmt.Errorf("Ошибка при считывании ввода пользователя")
	}

	return sizeMap, nil
}

func requestMapValue() (valueMap string, err error) {

	if _, err := fmt.Scanln(&valueMap); err != nil {
		return "", fmt.Errorf("Ошибка при считывании ввода пользователя")
	}

	return valueMap, nil
}

func main() {

	var sizeMap uint64
	var size_map1 uint64
	var size_map2 uint64
	var valueMap string
	var hashedValue uint64
	var err error

	// --> Начинаем запрашивать размеры массивов
	fmt.Println("Введите размер первого массива")
	size_map1, err = requestMapSize()
	if err != nil {
		fmt.Println("Ошибка при вводе размера первого массива")
		return
	}
	sizeMap = size_map1
	if sizeMap <= 0 {
		fmt.Println("Размер массива не может быть равен 0 или иметь отрицательное значение")
		return
	}

	map1 := make(map[uint64]string, sizeMap)

	// на всякий случай сбрасываем sizeMap
	sizeMap = 0

	fmt.Println("Введите размер первого массива")
	size_map2, err = requestMapSize()

	if err != nil {
		fmt.Println("Ошибка при вводе размера второго массива")
		return
	}
	sizeMap = size_map2
	if sizeMap <= 0 {
		fmt.Println("Размер массива не может быть равен 0 или иметь отрицательное значение")
		return
	}

	map2 := make(map[uint64]string, sizeMap)
	// Закончили запрашивать элементы массивов <--

	// --> Начинаем запрашивать элементы массивов
	fmt.Println("Укажите элементы первого массива. Каждый элемент массива указывается на новой строке")
	for i := 0; i < int(size_map1); i++ {
		valueMap, err = requestMapValue()

		if err != nil {
			fmt.Println("Ошибка ввода значения массива")
			return
		}

		hashedValue = hashstr(valueMap)
		// println("vm:", valueMap, "hv:", hashedValue)

		map1[hashedValue] = valueMap

	}

	valueMap = ""
	hashedValue = 0

	fmt.Println("Укажите элементы второго массива. Каждый элемент массива указывается на новой строке")
	for i := 0; i < int(size_map2); i++ {
		valueMap, err = requestMapValue()

		if err != nil {
			fmt.Println("Ошибка ввода значения массива")
			return
		}

		hashedValue = hashstr(valueMap)
		// println("vm:", valueMap, "hv:", hashedValue)

		map2[hashedValue] = valueMap

	}

	var resultIntersection []string

	// Заранее создадим результирующий массив с подходящей ёмкостью. Ёмкость не может быть больше, чем минимальный размер из двух мап, ведь в массиве не может оказаться элементов больше чем в меньшей мапе
	if size_map1 <= size_map2 {
		resultIntersection = make([]string, 0, size_map1)
	} else {
		resultIntersection = make([]string, 0, size_map2)
	}

	// Для оптимизации определим какой массив меньше и по нему будем проходить, ведь в результате не может оказаться элементов больше чем в меньшей мапе - мы ищем пересечения.
	if size_map1 <= size_map2 {
		for k, v := range map1 {

			if map1[k] == map2[k] {
				resultIntersection = append(resultIntersection, v)
			}

		}
	} else {
		for k, v := range map2 {

			if map1[k] == map2[k] {
				resultIntersection = append(resultIntersection, v)
			}

		}
	}

	fmt.Println(resultIntersection)

}
