package parentheses

import (
	"log"
)

type holder struct {
	text   []rune
	result []string
}

func GenerateParenthesis(n int) []string {
	switch {
	case n == 0:
		return nil
	case n == 1:
		return []string{"()"}
	case n == 2:
		return []string{"(())", "()()"}
	}
	generator := holder{
		text: make([]rune, n*2),
	}
	generator.generate(n, n, 0)
	return generator.result
}

func (h *holder) generate(open, close, idx int) {
	if open == 0 && close == 0 {
		h.result = append(h.result, string(h.text))
		return
	}

	if open > 0 {
		h.text[idx] = '('
		h.generate(open-1, close, idx+1)
	}
	if close > 0 && open < close {
		h.text[idx] = ')'
		h.generate(open, close-1, idx+1)
	}
}

func GenerateParenthesisInternet(n int) []string {
	res := make([]string, 0, n*n)
	bytes := make([]byte, n*2)
	dfs(n, n, 0, bytes, &res)
	return res
}

var stack int

func dfs(left, right, idx int, bytes []byte, res *[]string) {
	locally := stack
	stack++
	if left == 0 && right == 0 {
		log.Println("stack", locally, "left", left, "right", right, "idx", idx, "bytes", string(bytes))
		*res = append(*res, string(bytes))
		log.Println("res", res)
		return
	}

	if left > 0 {
		bytes[idx] = '('
		log.Println("stackL", locally, "left", left, "right", right, "idx", idx, "bytes", string(bytes))
		dfs(left-1, right, idx+1, bytes, res)
	}

	if right > 0 && left < right {
		bytes[idx] = ')'
		log.Println("stackR", locally, "left", left, "right", right, "idx", idx, "bytes", string(bytes))
		dfs(left, right-1, idx+1, bytes, res)
	}
}
