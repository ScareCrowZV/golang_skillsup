// // N-ый элемент арифметической прогрессии. Прямая рекурсия.
// func ArithmeticProgressionElement(a1, d, n int) int{
// 	if n == 1 {
// 		return a1
// 	}
// 	return d + ArithmeticProgressionElement(a1, d, n-1)
// }

// // N-ый элемент геометрической прогрессии. Косвенная рекурсия.
// func GeometricProgressionElement(b1, q, n int) int{
// 	if n == 1 {
// 		return b1
// 	}
// 	return q * ProgressionElement(b1, q, n)
// }

// func ProgressionElement(b1, q, n int) int {
// 	return GeometricProgressionElement(b1, q, n-1)
// }