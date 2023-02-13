package eval

import (
	"container/list"
	"fmt"
	"unicode"
)

// InfixToPostFix converts math expression to an arithmetic expression in a Reverse Polish notation
func InfixToPostFix(expr string) (*list.List, error) {
	if len(expr) == 0 {
		return nil, fmt.Errorf("error empty input")
	}

	queue := list.New()
	stack := list.New()

	for i := 0; i < len(expr); i++ {
		if expr[i] == ' ' {
			continue
		}
		val := string(expr[i])

		if i > 0 && unicode.IsDigit(rune(expr[i])) && unicode.IsDigit(rune(expr[i-1])) {
			return nil, fmt.Errorf("error invalid input, numbers can be between 0-9")
		}

		if unicode.IsDigit(rune(expr[i])) {
			queue.PushBack(val)
		} else if expr[i] == ')' {
			for stack.Len() != 0 && stack.Back().Value.(string) != "(" {
				elem := stack.Back()
				data := elem.Value
				stack.Remove(elem)
				queue.PushBack(data)
			}

			stack.Remove(stack.Back())
		} else if expr[i] == '(' {
			stack.PushBack("(")
		} else {
			if stack.Len() != 0 && precedence(val) > precedence(stack.Back().Value.(string)) {
				stack.PushBack(val)
			} else {
				for stack.Len() != 0 && precedence(val) <= precedence(stack.Back().Value.(string)) {
					queue.PushBack(stack.Back().Value.(string))
					stack.Remove(stack.Back())
				}
				stack.PushBack(val)
			}
		}
	}

	for stack.Len() != 0 {
		elem := stack.Back()
		queue.PushBack(elem.Value)
		stack.Remove(elem)
	}

	return queue, nil
}
