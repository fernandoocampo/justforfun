package pickpeaks_test

import (
	"testing"

	"github.com/fernandoocampo/justforfun/pickpeaks"
)

func TestPickPeaks(t *testing.T) {
	testcases := map[string]struct {
		input []int
		want  pickpeaks.PosPeaks
	}{
		// "one_peak": {
		// 	input: []int{0, 1, 2, 5, 1, 0},
		// 	want: pickpeaks.PosPeaks{
		// 		Pos:   []int{3},
		// 		Peaks: []int{5},
		// 	},
		// },
		// "two_peaks": {
		// 	input: []int{0, 1, 2, 5, 1, 6, 0},
		// 	want: pickpeaks.PosPeaks{
		// 		Pos:   []int{3, 5},
		// 		Peaks: []int{5, 6},
		// 	},
		// },
		// "three_peaks": {
		// 	input: []int{0, 1, 2, 5, 1, 6, 2, 4, 1, 8, 7, 0},
		// 	want: pickpeaks.PosPeaks{
		// 		Pos:   []int{3, 5, 7, 9},
		// 		Peaks: []int{5, 6, 4, 8},
		// 	},
		// },
		// "two_items": {
		// 	input: []int{5, 6},
		// 	want:  pickpeaks.PosPeaks{},
		// },
		// "no_one": {
		// 	input: []int{0, 1, 1, 1, 1, 0},
		// 	want:  pickpeaks.PosPeaks{},
		// },
		// "1, 2, 3, 6, 4, 1, 2, 3, 2, 1": {
		// 	input: []int{1, 2, 3, 6, 4, 1, 2, 3, 2, 1},
		// 	want: pickpeaks.PosPeaks{
		// 		Pos:   []int{3, 7},
		// 		Peaks: []int{6, 3},
		// 	},
		// },
		"3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 2, 2, 1": {
			input: []int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 2, 2, 1},
			want: pickpeaks.PosPeaks{
				Pos:   []int{3, 7, 10},
				Peaks: []int{6, 3, 2},
			},
		},
		// "3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3": {
		// 	input: []int{3, 2, 3, 6, 4, 1, 2, 3, 2, 1, 2, 3},
		// 	want: pickpeaks.PosPeaks{
		// 		Pos:   []int{3, 7},
		// 		Peaks: []int{6, 3},
		// 	},
		// },
	}
	for test, data := range testcases {
		got := pickpeaks.PickPeaks(data.input)
		pickPeaksDataValidator(t, test, got, data.want)
	}
}

func pickPeaksDataValidator(t *testing.T, test string, got, want pickpeaks.PosPeaks) {
	t.Helper()

	if len(got.Peaks) != len(want.Peaks) {
		t.Log("got", got)
		t.Errorf("%s - expected Peaks = %d, but got %d", test, len(want.Peaks), len(got.Peaks))
		return
	}

	if len(got.Pos) != len(want.Pos) {
		t.Errorf("%s - expected Peak Positions = %d, but got %d", test, len(want.Pos), len(got.Pos))
		return
	}

	for i, v := range want.Pos {
		if v != got.Pos[i] {
			t.Errorf("%s - expected Pos[%d] = %d, but got %d", test, i, v, got.Pos[i])
			break
		}
	}
	for i, v := range want.Peaks {
		if v != got.Peaks[i] {
			t.Errorf("%s - expected Peaks[%d] = %d, but got %d", test, i, v, got.Peaks[i])
			break
		}
	}
}
