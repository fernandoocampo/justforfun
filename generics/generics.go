package generics

import (
	"errors"
	"fmt"
	"io"
	"strings"
)

// Vector defines a vector that contains a specific type
// given at compile time.
// it restricts the constructions of the vector to a single
// data type.
type Vector[T any] []T

// Node prints values.
// * the T is a generic type to be determined at compile time.
// * any says there is no constraint on what type can be used.
// * field value can be of some type.
// * Next and Previous needs to point to a node of that same type.
type Node[T any] struct {
	value    T
	Next     *Node[T]
	Previous *Node[T]
}

func NewNode[T any](value T) *Node[T] {
	newNode := Node[T]{
		value: value,
	}

	return &newNode
}

// Print prints given values separating each item with a dash.
// important bullets.
// * functions can have additional type parameter list `any`
// * type paramaters can be used by parameters and in function body.
// * each type parameter has a type constraint. Here is `any`.
// * any is a type constraint that permits any type.
// * any is predeclared by the compiler.
// * the capital letter T to describe that a slice of some type T.
// * The compiler requires that all generic types have a well defined constraint.
func Print[T any](writer io.Writer, values []T) {
	var builder strings.Builder
	for idx, v := range values {
		builder.WriteString(fmt.Sprint(v))

		if idx < len(values)-1 {
			builder.WriteRune('-')
		}
	}

	fmt.Fprint(writer, builder.String())
}

// Last return the last element of the vector.
// * the receiver still specifies it contains a generic type.
// * the zero value can be specified declaring explicitly the variable of type T.
// * or we can return simply: *new(T).
// * be careful with using T{}, it does not work well with all types. e.g. Scalar.
func (v Vector[T]) Last() (T, error) {
	var zero T // get the possible zero value from the type T.

	if len(v) == 0 {
		return zero, errors.New("empty")
	}

	return v[len(v)-1], nil
}

func (n *Node[T]) Value() T {
	return n.value
}
