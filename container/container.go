package container

func MaxArea(height []int) int {
	maxarea := 0
	for j := 0; j < len(height)-1; j++ {
		maxCurrentArea := height[j] * (len(height) - j - 1)
		if maxCurrentArea <= maxarea {
			continue
		}
		for i := j + 1; i < len(height); i++ {
			dist := i - j
			area := dist * height[i]
			if height[j] <= height[i] {
				area = dist * height[j]
			}
			if area > maxarea {
				maxarea = area
			}
		}
	}
	return maxarea
}

func MaxAreaRec(height []int, max int) int {
	if len(height) == 0 {
		return max
	}
	maxCurrentArea := height[0] * (len(height) - 1)
	if maxCurrentArea <= max {
		return MaxAreaRec(height[1:], max)
	}
	max = calculateAreas(height[1:], height[0], 1, max)
	return MaxAreaRec(height[1:], max)
}

func calculateAreas(height []int, val, dist, max int) int {
	if len(height) == 0 {
		return max
	}
	area := dist * height[0]
	if val <= height[0] {
		area = dist * val
	}
	if area > max {
		max = area
	}
	return calculateAreas(height[1:], val, dist+1, max)
}

func getHigher() [4]int {
	return [4]int{}
}
