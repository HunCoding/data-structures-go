package main

import "fmt"

func isValid(s string) bool {
	stack := Stack{}

	for _, charPos := range s {
		char := fmt.Sprintf("%c", charPos)
		if char == "(" || char == "{" || char == "[" {
			stack.push(char)
		} else {
			if stack.isEmpty() {
				return false
			}

			top := stack.peek()
			if (char == ")" && top == "(") ||
				(char == "}" && top == "{") ||
				(char == "]" && top == "[") {
				stack.pop()
			} else {
				return false
			}
		}
	}

	return stack.isEmpty()
}
