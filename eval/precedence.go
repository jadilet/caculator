package eval

// Precedence of math operator
func precedence(operator string) int {
	if operator == "+" || operator == "-" {
		return 1
	} else if operator == "*" || operator == "/" {
		return 2
	} 
	
	return 0
}