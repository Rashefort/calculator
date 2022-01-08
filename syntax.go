package main

import (
	"errors"
	"strings"
)

var MessageError = map[string]error{
	"WRONG_SYMBOL":         errors.New("Недопустимый символ в выражении."),
	"START_OPERATOR":       errors.New("Оператор в начале выражения."),
	"END_OPERATOR":         errors.New("Оператор в конце выражения."),
	"OPEN_BRACKET_BEFORE":  errors.New("Недопустимый символ перед открывающей скобкой."),
	"OPEN_BRACKET_AFTER":   errors.New("Недопустимый символ после открывающей скобки."),
	"CLOSE_BRACKET_BEFORE": errors.New("Недопустимый символ перед закрывающей скобкой."),
	"CLOSE_BRACKET_AFTER":  errors.New("Недопустимый символ после закрывающей скобки."),
	"DISPARITY_BRACKETS":   errors.New("Несоответствие открывающих и закрывающих скобок."),
	"BEFORE_DOT":           errors.New("Недопустимый символ перед точкой."),
	"AFTER_DOT":            errors.New("Недопустимый символ после точки."),
	"MULTI_DOT":            errors.New("Лишние точки в числе."),
	"AFTER_OPERATOR":       errors.New("Недопустимый символ после оператора."),
}

type Syntax struct {
	Expression string
	Length     int
}

// Проверка допустимых символов ("0123456789.+-*/^").
func (s *Syntax) checkSymbols() error {
	for _, symbol := range s.Expression {
		if !strings.ContainsRune("0123456789.()+-*/^", symbol) {
			return MessageError["WRONG_SYMBOL"]
		}
	}
	return nil
}

// Проверка начала и конца выражения.
func (s *Syntax) checkEnds() error {
	switch {
	case !isDigit(s.Expression[0]) && s.Expression[0] != '(':
		return MessageError["START_OPERATOR"]
	case !isDigit(s.Expression[s.Length-1]) && s.Expression[s.Length-1] != ')':
		return MessageError["END_OPERATOR"]
	}
	return nil
}

// Проверка точки в float числах.
func (s *Syntax) checkDots() error {
	for index, symbol := range s.Expression {
		if symbol == '.' {
			if !isDigit(s.Expression[index-1]) {
				return MessageError["BEFORE_DOT"]
			} else if !isDigit(s.Expression[index+1]) {
				return MessageError["AFTER_DOT"]
			}
			for i := index + 1; i < s.Length; i++ {
				if s.Expression[i] == '.' {
					return MessageError["MULTI_DOT"]
				} else if !isDigit(s.Expression[i]) {
					break
				}
			}
		}
	}
	return nil
}

// Проверка соответствия открывающих и закрывающих скобок. Также перед
// открывающей скобкой и после закрывающей не может быть цифра, а после
// открывающей и перед закрывающей арифметические операторы.
func (s *Syntax) checkBrackets() error {
	brackets := 0

	for i := 0; i < s.Length; i++ {
		switch s.Expression[i] {
		case '(':
			if i > 0 && strings.ContainsRune("0123456789.)", rune(s.Expression[i-1])) {
				return MessageError["OPEN_BRACKET_BEFORE"]
			} else if i < s.Length-1 && strings.ContainsRune("+*/^.)", rune(s.Expression[i+1])) {
				return MessageError["OPEN_BRACKET_AFTER"]
			}
			brackets++

		case ')':
			if i > 0 && strings.ContainsRune("+-*/^.(", rune(s.Expression[i-1])) {
				return MessageError["CLOSE_BRACKET_BEFORE"]
			} else if i < s.Length-1 && strings.ContainsRune("0123456789.(", rune(s.Expression[i+1])) {
				return MessageError["CLOSE_BRACKET_AFTER"]
			}
			brackets--
		}
		if brackets < 0 {
			return MessageError["DISPARITY_BRACKETS"]
		}
	}
	if brackets > 0 {
		return MessageError["DISPARITY_BRACKETS"]
	}
	return nil
}

// Проверка расстановки операторов
func (s *Syntax) checkOperators() error {
	for i := 0; i < s.Length-1; i++ {
		if strings.ContainsRune("+-*/^", rune(s.Expression[i])) &&
			strings.ContainsRune("+-*/^", rune(s.Expression[i-1])) {
			return MessageError["AFTER_OPERATOR"]
		}
	}
	return nil
}

// Перебор всех проверок
func (s *Syntax) CheckExpression(expression string) error {
	s.Expression = expression
	s.Length = len(expression)

	// Выражение вида "-2+(-2^2)" будет преобразовано в "0-2+(0-2^2)"
	for i := 0; i < s.Length; i++ {
		if s.Expression[i] == '-' && (i == 0 || s.Expression[i-1] == '(') {
			s.Expression = s.Expression[0:i] + "0" + s.Expression[i:]
			s.Length++
			i++
		}
	}

	// Перебор по всем методам проверок.
	funcs := []func() error{s.checkSymbols, s.checkEnds, s.checkDots, s.checkBrackets, s.checkOperators}
	for _, check := range funcs {
		err := check()
		if err != nil {
			return err
		}
	}
	return nil
}
