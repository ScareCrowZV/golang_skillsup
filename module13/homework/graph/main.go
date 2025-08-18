package main

import (
	"fmt"
)

func addEdgeUndirected(m [][]int, adjacencyArray [][]int, i, j int) {
	// Проставляем направления для неориентированного графа для матрицы
	m[i][j] = 1
	m[j][i] = 1

	// Устанавливаем значение для матрицы смежности
	adjacencyArray[i] = append(adjacencyArray[i], j)
	adjacencyArray[j] = append(adjacencyArray[j], i)
}

func addEdgeUndirectedWithDistance(m [][]int, adjacencyArray [][]int, adjacencyDistance [][]int, i, j, distance int) {
	// Проставляем направления для неориентированного графа для матрицы
	m[i][j] = 1
	m[j][i] = 1

	// Устанавливаем значение для матрицы смежности
	adjacencyArray[i] = append(adjacencyArray[i], j)
	adjacencyArray[j] = append(adjacencyArray[j], i)

	adjacencyDistance[i] = append(adjacencyDistance[i], distance)
	adjacencyDistance[j] = append(adjacencyDistance[j], distance)

}

func addEdgeDirected(m [][]int, adjacencyArray [][]int, i, j int) {
	// Проставляем направления для неориентированного графа для матрицы
	m[i][j] = 1

	// Устанавливаем значение для матрицы смежности
	adjacencyArray[i] = append(adjacencyArray[i], j)

}

func addEdgeDirectedWithDistance(m [][]int, adjacencyArray [][]int, adjacencyDistance [][]int, i, j, distance int) {
	// Проставляем направления для неориентированного графа для матрицы
	m[i][j] = 1

	// Устанавливаем значение для матрицы смежности
	adjacencyArray[i] = append(adjacencyArray[i], j)

	adjacencyDistance[i] = append(adjacencyDistance[i], distance)

}

func printMatrix(m [][]int) {
	fmt.Print("  ")
	for i := range m {
		fmt.Print(" ", i)
	}

	fmt.Println()

	for i, j := range m {

		fmt.Println(i, j)

	}
}

func printAdjacencyArray(m [][]int) {
	for i := range m {
		fmt.Println(i, m[i])
	}
}

func prepareBfs(m [][]int) []int {
	// Слайс с результатом обхода. Записываем узлы в том порядке, в котором прошли их
	var resultWalk []int
	// Слайс для отметки посещённых узлов
	visited := make([]bool, len(m))

	for i := 0; i < len(m); i++ {
		if !visited[i] {
			// Для каждого графа в матрице смежности выполняем поиск в ширину
			bfs(m, visited, &resultWalk, i)
		}
	}
	return resultWalk
}

func bfs(m [][]int, visited []bool, resultWalk *[]int, startElement int) {
	// Очередь для записи связанных узлов для последующего обхода
	queue := []int{}
	// Первый узел отмечаем посещённым, т.к. с него начинаем обход
	visited[startElement] = true
	// Записываем первый узел в очередь, т.к. с него начинаем обход и пойдём по его соседям
	queue = append(queue, startElement)

	for len(queue) > 0 {
		// Выбираем первый элемент в очереди
		// и... уменьшаем очередь на один взятый элемент
		currentElement := queue[0]
		queue = queue[1:]
		// Записываем в слайс пройденных текущий пройденный элемент
		*resultWalk = append(*resultWalk, currentElement)

		// fmt.Printf("Обходим %d\n", currentElement)

		// Начинаем обходить матрицу смежности для определения соседей. Проще говоря, обходим соседей
		for _, v := range m[currentElement] {
			// Если мы не посещали узел, то теперь считаем его пройденным и добавляем соседа в очередь для обхода, т.к. у него могут быть связи с другими соседями
			if !visited[v] {

				visited[v] = true
				queue = append(queue, v)
			}
		}
	}
}

func dijeikstraAlgorithm(m [][]int, distanceArray [][]int, startNode int) {
	var visited []bool = make([]bool, len(m))
	var distance []int = make([]int, len(m))
	queue := []int{}

	visited[startNode] = true
	distance[startNode] = 0
	queue = append(queue, startNode)

	for len(queue) > 0 {

		currentNode := queue[0]
		queue = queue[1:]

		for i, v := range m[currentNode] {

			if !visited[v] {
				visited[currentNode] = true
				if distance[v] >= distance[currentNode]+distanceArray[currentNode][i] || distance[v] == 0 {
					distance[v] = distance[currentNode] + distanceArray[currentNode][i]
				}
				queue = append(queue, v)
			}

		}
	}

	fmt.Println(distance)
	fmt.Println(visited)

}

func main() {

	// Задаём количество узлов графа. Это же и будет размером матрицы смежности.
	var graphSize int = 10

	// Инициализируем матрицу смежности
	var matrix = make([][]int, graphSize)

	// Инциализируем как список связанных графов
	var adjacencyArray = make([][]int, graphSize)

	// Инициализируем список дистанций для ребер графов в adjacencyArray
	var adjacencyDistance = make([][]int, graphSize)

	// Заполняем элементы матрицы пустыми элементами(0 по умолчанию)
	for i := 0; i < graphSize; i++ {
		for j := 0; j < graphSize; j++ {
			matrix[i] = make([]int, graphSize)
		}
	}

	// Устанавливаем связи графа

	// // Неориентированный связанный граф из 6 элементов. Не забудьте изменить переменную graphSize на 6
	// addEdgeUndirected(matrix, adjacencyArray, 0, 1)
	// addEdgeUndirected(matrix, adjacencyArray, 1, 2)
	// addEdgeUndirected(matrix, adjacencyArray, 2, 5)
	// addEdgeUndirected(matrix, adjacencyArray, 5, 4)
	// addEdgeUndirected(matrix, adjacencyArray, 4, 3)
	// addEdgeUndirected(matrix, adjacencyArray, 3, 0)
	// addEdgeUndirected(matrix, adjacencyArray, 1, 4)

	// // Неориентированный несвязанный граф из 6 элементов. Не забудьте изменить переменную graphSize на 6
	// addEdgeUndirected(matrix, adjacencyArray, 0, 1)
	// addEdgeUndirected(matrix, adjacencyArray, 1, 2)
	// addEdgeUndirected(matrix, adjacencyArray, 3, 4)

	// // Ориентированный граф из 4 элементов. Не забудьте изменить переменную graphSize на 4
	// addEdgeDirected(matrix, adjacencyArray, 0, 1)
	// addEdgeDirected(matrix, adjacencyArray, 0, 3)
	// addEdgeDirected(matrix, adjacencyArray, 1, 2)
	// addEdgeDirected(matrix, adjacencyArray, 2, 3)

	// // Неориентированный граф для проверки Дейкстры. Не забудьте изменить переменную graphSize на 7
	// addEdgeUndirectedWithDistance(matrix, adjacencyArray, adjacencyDistance, 0, 1, 2)
	// addEdgeUndirectedWithDistance(matrix, adjacencyArray, adjacencyDistance, 0, 2, 6)
	// addEdgeUndirectedWithDistance(matrix, adjacencyArray, adjacencyDistance, 1, 3, 5)
	// addEdgeUndirectedWithDistance(matrix, adjacencyArray, adjacencyDistance, 2, 3, 8)
	// addEdgeUndirectedWithDistance(matrix, adjacencyArray, adjacencyDistance, 3, 4, 10)
	// addEdgeUndirectedWithDistance(matrix, adjacencyArray, adjacencyDistance, 3, 5, 15)
	// addEdgeUndirectedWithDistance(matrix, adjacencyArray, adjacencyDistance, 4, 6, 2)
	// addEdgeUndirectedWithDistance(matrix, adjacencyArray, adjacencyDistance, 5, 6, 6)

	// // Ориентированный граф из 10 элементов для проверки Дейкстры. Не забудьте изменить переменную graphSize на 10
	addEdgeDirectedWithDistance(matrix, adjacencyArray, adjacencyDistance, 0, 1, 4)
	addEdgeDirectedWithDistance(matrix, adjacencyArray, adjacencyDistance, 0, 2, 7)
	addEdgeDirectedWithDistance(matrix, adjacencyArray, adjacencyDistance, 0, 3, 1)
	addEdgeDirectedWithDistance(matrix, adjacencyArray, adjacencyDistance, 1, 3, 5)
	addEdgeDirectedWithDistance(matrix, adjacencyArray, adjacencyDistance, 3, 6, 100)
	addEdgeDirectedWithDistance(matrix, adjacencyArray, adjacencyDistance, 2, 4, 5)
	addEdgeDirectedWithDistance(matrix, adjacencyArray, adjacencyDistance, 4, 5, 10)
	addEdgeDirectedWithDistance(matrix, adjacencyArray, adjacencyDistance, 6, 5, 2)

	fmt.Println("Матрица")
	printMatrix(matrix)

	fmt.Println("список связанных графов")
	printAdjacencyArray(adjacencyArray)

	fmt.Println("Список дистанций")
	fmt.Println(adjacencyDistance)

	v := prepareBfs(adjacencyArray)
	fmt.Println("Результат обхода в ширину:", v)

	fmt.Println("Результат работы алгоритма Дейкстры")
	dijeikstraAlgorithm(adjacencyArray, adjacencyDistance, 0)
}
