package parentheses_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/parentheses"
	"github.com/stretchr/testify/assert"
)

func TestGenerateParenthesis(t *testing.T) {
	cases := map[string]struct {
		n    int
		want []string
	}{
		"1": {
			n:    1,
			want: []string{"()"},
		},
		"2": {
			n:    2,
			want: []string{"(())", "()()"},
		},
		"3": {
			n:    3,
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		"4": {
			n:    4,
			want: []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"},
		},
		"5": {
			n:    5,
			want: []string{"((((()))))", "(((()())))", "(((())()))", "(((()))())", "(((())))()", "((()(())))", "((()()()))", "((()())())", "((()()))()", "((())(()))", "((())()())", "((())())()", "((()))(())", "((()))()()", "(()((())))", "(()(()()))", "(()(())())", "(()(()))()", "(()()(()))", "(()()()())", "(()()())()", "(()())(())", "(()())()()", "(())((()))", "(())(()())", "(())(())()", "(())()(())", "(())()()()", "()(((())))", "()((()()))", "()((())())", "()((()))()", "()(()(()))", "()(()()())", "()(()())()", "()(())(())", "()(())()()", "()()((()))", "()()(()())", "()()(())()", "()()()(())", "()()()()()"},
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := parentheses.GenerateParenthesis(data.n)
			assert.Equal(st, data.want, got)
		})
	}
}

func TestGenerateParenthesisInternet(t *testing.T) {
	cases := map[string]struct {
		n    int
		want []string
	}{
		"1": {
			n:    1,
			want: []string{"()"},
		},
		"2": {
			n:    2,
			want: []string{"(())", "()()"},
		},
		"3": {
			n:    3,
			want: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
		"4": {
			n:    4,
			want: []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"},
		},
	}

	for name, data := range cases {
		t.Run(name, func(st *testing.T) {
			got := parentheses.GenerateParenthesisInternet(data.n)
			assert.Equal(st, data.want, got)
		})
	}
}
