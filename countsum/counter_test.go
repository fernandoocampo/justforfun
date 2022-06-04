package countsum_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/countsum"
)

func TestCountPositivesSumNegatives(t *testing.T) {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, -11, -12, -13, -14, -15}
	res := []int{10, -65}

	got := countsum.CountPositivesSumNegatives(arr)

	if len(res) != len(got) {
		t.Errorf("want %v, but got %v", res, got)
		t.FailNow()
	}

	for i := 0; i < len(res); i++ {
		if res[i] != got[i] {
			t.Errorf("want %v, but got %v", res, got)
			t.FailNow()
		}
	}
}
