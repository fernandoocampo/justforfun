package parentheses

func IsValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	if s[0] == '}' || s[0] == ']' || s[0] == ')' {
		return false
	}
	return isBalanced(s, []rune{})
}

func IsValidLoop(s string) bool {
	if len(s) == 0 {
		return true
	}
	if s[0] == '}' || s[0] == ']' || s[0] == ')' {
		return false
	}
	var stack []rune
	for _, v := range s {
		if isOpening(v) {
			stack = append(stack, v)
			continue
		}

		if len(stack) > 0 && isClosing(v, stack[len(stack)-1]) {
			stack = stack[:len(stack)-1]
			continue
		}
		return false
	}
	return len(stack) == 0
}

func isBalanced(s string, stack []rune) bool {
	if len(s) == 0 {
		return len(stack) == 0
	}
	if isOpening(rune(s[0])) {
		return isBalanced(s[1:], append(stack, rune(s[0])))
	}

	if len(stack) > 0 && isClosing(rune(s[0]), stack[len(stack)-1]) {
		return isBalanced(s[1:], stack[:len(stack)-1])
	}
	return false
}

func isOpening(ch rune) bool {
	return ch == '{' || ch == '(' || ch == '['
}

func isClosing(new, old rune) bool {
	return new == '}' && old == '{' ||
		new == ']' && old == '[' ||
		new == ')' && old == '('
}
