package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Ascii код 0 - 48, 9 - 57. byte - uint8
func isDigit(symbol byte) bool {
	return symbol-48 < 10
}

func trimZero(s string) string {
	for strings.ContainsRune(s, '.') == true && s[len(s)-1] == '0' || s[len(s)-1] == '.' {
		s = s[:len(s)-1]
	}

	return s
}

func calculate(expression string) {
	var notations Notations
	var stack StackFloat

	err := notations.New(expression)
	if err != nil {
		fmt.Printf("Ошибка: %s\n", err)
		os.Exit(2)
	}

	for _, lexeme := range notations.Reverse {
		if isDigit(lexeme[0]) {
			value, err := strconv.ParseFloat(lexeme, 64)
			if err != nil {
				fmt.Printf("Ошибка: %s\n", err)
				os.Exit(1)
			}
			stack.Push(value)

		} else {
			var v1, v2 float64 = stack.Pop(), stack.Pop()
			switch lexeme {
			case "+":
				stack.Push(v2 + v1)
			case "-":
				stack.Push(v2 - v1)
			case "*":
				stack.Push(v2 * v1)
			case "/":
				stack.Push(v2 / v1)
			default:
				stack.Push(math.Pow(v2, v1))
			}
		}
	}
	fmt.Printf(trimZero(fmt.Sprintf("%f", stack.Pop())))
}

func main() {
	if len(os.Args) > 1 {
		calculate(strings.Join(os.Args[1:], ""))
	} else {
		fmt.Println("Консольный калькулятор \"А нам всё равно\", версия 1.01 (c) telegram: @kampbell")
		fmt.Println("Поддерживаемые операции:")
		fmt.Println("    ()\t- скобки\n    +\t- сложение\n    -\t- вычитание")
		fmt.Println("    *\t- умножение\n    /\t- деление\n    ^\t- степень")
		fmt.Println("Пример использования: = 7 + 6 * (5.4 - 3)^2")
	}
}
