package main

import "fmt"

func addEdge(m [][]int, adjacencyMatrix [][]int, i, j int) {
	// Проставляем направления для неориентированного графа для матрицы
	m[i][j] = 1
	m[j][i] = 1

	// Устанавливаем значение для матрицы смежности
	adjacencyMatrix[i] = append(adjacencyMatrix[i], j)
	adjacencyMatrix[j] = append(adjacencyMatrix[j], i)

}

func printMatrix(m [][]int) {
	fmt.Println("  ", 0, 1, 2, 3, 4, 5)
	for i, j := range m {

		fmt.Println(i, j)

	}
}

func printAdjacencyMatrix(m [][]int) {
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

func main() {

	// Задаём размер квадратной матрицы
	var graphSize int = 6

	// Инициализируем двумерный массив
	var matrix = make([][]int, graphSize)

	// Инциализируем матрицу смежности
	var adjacencyMatrix = make([][]int, graphSize)

	// Заполняем элементы массива пустыми элементами(0 по умолчанию)
	for i := 0; i < graphSize; i++ {
		for j := 0; j < graphSize; j++ {
			matrix[i] = make([]int, graphSize)
		}
	}

	// Устанавливаем связи графа
	// Неориентированный связанный граф
	// addEdge(matrix, adjacencyMatrix, 0, 1)
	// addEdge(matrix, adjacencyMatrix, 1, 2)
	// addEdge(matrix, adjacencyMatrix, 2, 5)
	// addEdge(matrix, adjacencyMatrix, 5, 4)
	// addEdge(matrix, adjacencyMatrix, 4, 3)
	// addEdge(matrix, adjacencyMatrix, 3, 0)
	// addEdge(matrix, adjacencyMatrix, 1, 4)

	// Неориентированный несвязанный граф
	addEdge(matrix, adjacencyMatrix, 0, 1)
	addEdge(matrix, adjacencyMatrix, 1, 2)
	addEdge(matrix, adjacencyMatrix, 3, 4)
	addEdge(matrix, adjacencyMatrix, 5, 5)

	fmt.Println("Матрица")
	printMatrix(matrix)

	fmt.Println("Матрица смежности")
	printAdjacencyMatrix(adjacencyMatrix)

	v := prepareBfs(adjacencyMatrix)
	fmt.Println("Результат обхода в ширину:", v)

}
