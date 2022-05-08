package problems

import "strconv"

func evalRPN(tokens []string) int {
	stack := make([]string, 0)
	operationsMap := map[string]bool {
		"+": true,
		"-": true,
		"*": true,
		"/": true,
	}

	for _, token := range tokens {
		_, isOperator := operationsMap[token]

		if !isOperator {
			stack = append(stack, token)
			continue
		}

		result := 0
		operandOne, _ := strconv.Atoi(stack[len(stack) - 2])
		operandTwo, _ := strconv.Atoi(stack[len(stack) - 1])

		if token == "+" {
			result = operandOne + operandTwo
		} else if token == "-" {
			result = operandOne - operandTwo
		} else if token == "*" {
			result = operandOne * operandTwo
		} else if token == "/" {
			result = operandOne / operandTwo
		}

		stack = stack[:len(stack) - 2]
		stack = append(stack, strconv.Itoa(result))
	}

	result, _ := strconv.Atoi(stack[0])
	return result
}