package balance

import (
	"errors"
)

var ErrInvalid error = errors.New("invalid")

func Validate(parentheses string) error {
	if len(parentheses) == 0 {
		return nil
	}
	if len(parentheses) < 2 || parentheses[0] == ')' {
		return ErrInvalid
	}
	stack := make([]rune, 0)
	for _, v := range parentheses {
		if v == '(' {
			stack = append(stack, v)
			continue
		}
		if len(stack) == 0 && v == ')' {
			return ErrInvalid
		}
		if v == ')' {
			stack = stack[:len(stack)-1]
		}
	}
	if len(stack) > 0 {
		return ErrInvalid
	}
	return nil
}

func ValidateRec(parentheses string) error {
	if isBalanced(parentheses, []rune{}) {
		return nil
	}
	return ErrInvalid
}

func isBalanced(parentheses string, stack []rune) bool {
	if len(parentheses) == 0 {
		return len(stack) < 1
	}
	if parentheses[0] == '(' {
		stack = append(stack, rune(parentheses[0]))
	}
	if len(stack) == 0 && parentheses[0] == ')' {
		return false
	}
	if parentheses[0] == ')' {
		stack = stack[:len(stack)-1]
	}
	return isBalanced(parentheses[1:], stack)
}
