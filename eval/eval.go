package eval

import "fmt"

// Eval evaluates arithmetic expression
// returns result as `expr = result` or `expr = error message`
func Eval(expr string) string {
	tokens, err := InfixToPostFix(expr)

	if err != nil {
		return fmt.Sprintf("%s = %s", expr, err.Error())
	}

	res, err := EvalRpn(tokens)

	if err != nil {
		return fmt.Sprintf("%s = %s", expr, err.Error())
	}

	return fmt.Sprintf("%s = %d", expr, res)
}
