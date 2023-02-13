package eval

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

// Evaluates a Reverse Polish Notation
func EvalRpn(tokens *list.List) (int, error) {
	if tokens.Len() == 0 {
		return 0, fmt.Errorf("error invalid input reverse polish notation")
	}

	if tokens.Len() == 2 {
		return 0, fmt.Errorf("error invalid input either negative literal or unary operation")
	}

	var res int
	stack := list.New()

	for iter := tokens.Front(); iter != nil; iter = iter.Next() {
		if stack.Len() >= 2 && strings.Contains("*/-+", iter.Value.(string)) {
			a := stack.Back().Value.(string)
			stack.Remove(stack.Back())
			b := stack.Back().Value.(string)
			stack.Remove(stack.Back())

			aInt, err := strconv.Atoi(a)

			if err != nil {
				return 0, err
			}

			bInt, err := strconv.Atoi(b)

			if err != nil {
				return 0, err
			}

			switch operator := iter.Value.(string); operator {
			case "*":
				stack.PushBack(fmt.Sprintf("%d", aInt*bInt))
			case "-":
				stack.PushBack(fmt.Sprintf("%d", bInt-aInt))
			case "+":
				stack.PushBack(fmt.Sprintf("%d", aInt+bInt))
			case "/":
				stack.PushBack(fmt.Sprintf("%d", bInt/aInt))
			}

			continue
		}

		stack.PushBack(iter.Value)
	}

	if stack.Len() == 0 {
		return res, nil
	}

	return strconv.Atoi(stack.Back().Value.(string))
}
