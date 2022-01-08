package main

// Инфиксная и обратная польская нотации
type Notations struct {
	Syntax
	Reverse []string
	Infix   []string
}

// Проверка синтаксических ошибок и создание нотаций
func (n *Notations) New(expression string) error {
	err := n.CheckExpression(expression)
	if err != nil {
		return err
	}

	n.infixNotations()
	n.reverseNotations()
	return nil
}

// Инфиксная нотация
func (n *Notations) infixNotations() {
	var start, end int
	for start = 0; start < n.Length; start++ {
		if isDigit(n.Expression[start]) {
			for end = start + 1; end < n.Length; end++ {
				if !isDigit(n.Expression[end]) && n.Expression[end] != '.' {
					break
				}
			}
			n.Infix = append(n.Infix, n.Expression[start:end])
			start = end - 1
		} else {
			n.Infix = append(n.Infix, n.Expression[start:start+1])
		}
	}
}

// Обратная польская нотация. Алгоритм сортировочной станции.
func (n *Notations) reverseNotations() {
	var weight = map[string]int{"^": 3, "*": 2, "/": 2, "+": 1, "-": 1, "(": 0}
	var stack StackString

	for _, lexeme := range n.Infix {
		switch {
		case isDigit(lexeme[0]):
			n.Reverse = append(n.Reverse, lexeme)
		case lexeme[0] == '(':
			stack.Push(lexeme)
		case lexeme[0] == ')':
			for st := stack.Pop(); st != "("; st = stack.Pop() {
				n.Reverse = append(n.Reverse, st)
			}
		default:
			for stack.deep > 0 && weight[stack.upper] >= weight[lexeme] {
				operator := stack.Pop()
				n.Reverse = append(n.Reverse, operator)
			}
			stack.Push(lexeme)
		}
	}
	for stack.deep > 0 {
		operator := stack.Pop()
		n.Reverse = append(n.Reverse, operator)
	}
}
