package generics_test

import (
	"bytes"
	"testing"

	"github.com/fernandoocampo/justforfun/generics"
	"github.com/stretchr/testify/assert"
)

func TestPrintIntSlices(t *testing.T) {
	t.Parallel()
	// Given
	slice := []int{1, 2, 3}
	expectedResult := "1-2-3"
	buffer := bytes.NewBufferString("")

	// When
	// notice that you can show the type of T explicitly.
	// but it is not longer required, since the compiler
	// knows your intentions.
	generics.Print[int](buffer, slice)
	result := buffer.String()

	// Then
	if result != expectedResult {
		t.Errorf("want: %s, but got: %s", expectedResult, result)
	}
}

func TestPrintSlices(t *testing.T) {
	t.Parallel()

	// Given
	cases := map[string]struct {
		input []any
		want  string
	}{
		"integers": {
			input: []any{1, 2, 3},
			want:  "1-2-3",
		},
		"runes": {
			input: []any{'a', 'b', 'c'},
			want:  "97-98-99",
		},
		"strings": {
			input: []any{"a", "b", "c"},
			want:  "a-b-c",
		},
	}

	for name, data := range cases {
		name, data := name, data
		t.Run(name, func(st *testing.T) {
			writer := bytes.NewBufferString("")

			generics.Print(writer, data.input)
			result := writer.String()
			if result != data.want {
				st.Errorf("want: %s, but got: %s", data.want, result)
			}
		})
	}
}

func TestVectorLastElement(t *testing.T) {
	t.Parallel()
	// Given
	aVector := generics.Vector[int]{1, 2, 3, 4, 5}
	expectedLastValue := 5
	// When
	got, err := aVector.Last()
	// Then
	if err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if expectedLastValue != got {
		t.Errorf("want: %d, but got: %d", expectedLastValue, got)
	}
}

func TestBidirectionalLinkedList(t *testing.T) {
	t.Parallel()
	// Given
	expectedPreviousValue := "a"
	expectedMiddleValue := "b"
	expectedNextValue := "c"

	previousNode := generics.NewNode("a")
	middleNode := generics.NewNode("b")
	nextNode := generics.NewNode("c")
	// When
	previousNode.Next = middleNode
	middleNode.Previous = previousNode
	middleNode.Next = nextNode
	nextNode.Previous = middleNode
	// Then
	assert.Equal(t, expectedPreviousValue, middleNode.Previous.Value())
	assert.Equal(t, expectedMiddleValue, middleNode.Value())
	assert.Equal(t, expectedNextValue, middleNode.Next.Value())
}
