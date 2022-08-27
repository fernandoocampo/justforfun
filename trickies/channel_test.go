package trickies_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/trickies"
)

func TestGetStringFromChannel(t *testing.T) {
	givenStream := make(chan int, 1)
	expectedValue := "321"

	got := trickies.GetStringFromChannel(givenStream)

	if expectedValue != got {
		t.Errorf("want: %q, but got %q", expectedValue, got)
	}
}
