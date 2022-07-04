package parentheses

func LongestValidParentheses(s string) int {
	if len(s) < 2 {
		return 0
	}
	var max int
	stack := []int{-1}

	for i, v := range s {
		switch v {
		case '(':
			stack = append(stack, i)
		case ')':
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
				continue
			}
			if (i - stack[len(stack)-1]) > max {
				max = i - stack[len(stack)-1]
			}
		}
	}
	return max
}
