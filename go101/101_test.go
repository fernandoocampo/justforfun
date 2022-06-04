package go101_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/go101"
)

func TestLongTimeRequest(t *testing.T) {
	done := make(chan interface{})
	defer close(done)
	want := int32(32)
	got := go101.LongTimeRequest(done, callAndReturn32)
	if want != got {
		t.Errorf("want %d, but got %d", want, got)
	}
}

func callAndReturn32() int32 {
	return 32
}
