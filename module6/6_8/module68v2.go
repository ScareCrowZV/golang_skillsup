package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func calculate(firstOperand, secondOperand float64, operator byte) (float64, string) {

	var result float64
	var errorResult string

	switch string(operator) {
	case "+":
		result = firstOperand + secondOperand
	case "-":
		result = firstOperand - secondOperand
	case "*":
		result = firstOperand * secondOperand
	case "\\":
		if secondOperand == 0 {
			errorResult = "Деление на 0 невозможно"
			break
		}
		result = firstOperand / secondOperand
	default:
		errorResult = "Выражение введено не верно"
	}

	return result, errorResult
}

func main() {

	var expression string
	var err error
	var reader *bufio.Reader
	var firstOperand float64
	var firstOperandBuffer string
	var firstOperandBufferFull bool = false
	var secondOperand float64
	var secondOperandBuffer string
	var operator byte
	var result float64
	var errorResult string

	fmt.Print(
		"Вас приветствует простой калькулятор!\n",
		"Введите выражение и нажмите клавишу Enter\n",
		"Например: \"1 + 1\" Enter\n",
		"И даже вещественные числа: \"10.5 \\ 5.5\" Enter\n",
		"Учитываться будут только первые 2 операнда\n",
		"\nВозможные операции:\n",
		"+ сложение\n",
		"- вычитание\n",
		"* умножение\n",
		"\\ деление\n")

	reader = bufio.NewReader(os.Stdin)
	expression, err = reader.ReadString('\n')

	if err != nil {
		fmt.Printf("Ошибка ввода: %s", err)
		return
	}

	for i := 0; i < len(expression); i++ {

		switch {
		case (expression[i] >= 48 && expression[i] <= 57 || string(expression[i]) == "." || string(expression[i]) == ",") && !firstOperandBufferFull:
			firstOperandBuffer += string(expression[i])
		case (expression[i] >= 48 && expression[i] <= 57 || string(expression[i]) == "." || string(expression[i]) == ",") && firstOperandBufferFull && operator != 0:
			secondOperandBuffer += string(expression[i])
		case string(expression[i]) == "+":
			operator = expression[i]
			firstOperandBufferFull = true
		case string(expression[i]) == "-":
			operator = expression[i]
			firstOperandBufferFull = true
		case string(expression[i]) == "*":
			operator = expression[i]
			firstOperandBufferFull = true
		case string(expression[i]) == "\\":
			operator = expression[i]
			firstOperandBufferFull = true
		case expression[i] == 13: // Перенос строки игнорируем
		case expression[i] == 10: // Возврат каретки игнорируем
		case expression[i] == 32: // Пробел игнорируем
		default:
			fmt.Printf("Введён недопустимый символ %s, выражение введено не верно\n", string(expression[i]))
			return
		}
	}

	// Пробуем привести первый операнд
	firstOperand, _ = strconv.ParseFloat(firstOperandBuffer, 64)
	secondOperand, _ = strconv.ParseFloat(secondOperandBuffer, 64)

	result, errorResult = calculate(firstOperand, secondOperand, operator)

	if errorResult != "" {
		fmt.Printf("Ошибка при вычислении выражения: %s\n", errorResult)
		return
	}

	if operator > 0 {
		fmt.Printf("Ваше выражение: %f %s %f = %f\n", firstOperand, string(operator), secondOperand, result)
	}

	fmt.Println("Нажмите Enter для завершения")
	_, _ = reader.ReadString('\n')
}
