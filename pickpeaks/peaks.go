package pickpeaks

type PosPeaks struct {
	Pos   []int
	Peaks []int
}

// Also, beware of plateaus !!! [1, 2, 2, 2, 1] has a peak while [1, 2, 2, 2, 3]
// and [1, 2, 2, 2, 2] do not. In case of a plateau-peak, please only return
// the position and value of the beginning of the plateau. For example:
// pickPeaks([1, 2, 2, 2, 1]) returns {pos: [1], peaks: [2]} (or equivalent
// in other languages)
func PickPeaks(array []int) PosPeaks {
	if len(array) <= 2 {
		return PosPeaks{}
	}
	newPick := true
	var plateau []int
	if array[0] == 1 {
		plateau = append(plateau, array[0])
	}
	var result [][2]int
	for i := 1; i < len(array)-1; i++ {
		if array[i] == 2 && len(plateau) > 0 {
			plateau = append(plateau, array[i])
		}
		if array[i] == 1 {

		}
		if array[i] < 3 && len(result) == i-1 {
			continue
		}
		if array[i] < 3 {
			newPick = true
			continue
		}
		if newPick {
			result = append(result, [2]int{i, array[i]})
			// j++
			newPick = false
			continue
		}
		if array[i] > result[len(result)-1][1] {
			result[len(result)-1] = [2]int{i, array[i]}
			continue
		}
	}
	var finalResult PosPeaks
	for _, v := range result {
		finalResult.Pos = append(finalResult.Pos, v[0])
		finalResult.Peaks = append(finalResult.Peaks, v[1])
	}
	return finalResult
}

func isPlateau() bool {
	return false
}
