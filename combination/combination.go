package combination

import (
	"fmt"
	"sort"
	"strconv"
)

func CombinationSum(candidates []int, target int) [][]int {
	var result [][]int
	for idx, v1 := range candidates {
		if v1 > target {
			continue
		}
		if v1 == target {
			result = append(result, []int{v1})
			continue
		}
		if target%v1 == 0 {
			result = append(result, fillValues(v1, target/v1))
			// continue
		}
		if idx == len(candidates)-1 {
			break
		}
		sum := v1
		sumarr := []int{v1}
		for jdx, v2 := range candidates {
			if idx == jdx {
				continue
			}
			if v2 == target || v1+v2 > target {
				continue
			}
			// if v1+v2 == target {
			// 	result = append(result, []int{v1, v2})
			// 	continue
			// }
			if (target-v2)%v1 == 0 && (target-v2)/v1 != 1 {
				result = append(result, fillValuesWithInitialValue(v1, v2, target-v2))
			}
			// if (target-v1)%v2 == 0 {
			// 	result = append(result, fillValuesWithInitialValue(v2, v1, target-v1))
			// }
			if sum+v2 > target {
				continue
			}
			if (sum + v2) == target {
				result = append(result, append(sumarr, v2))
				continue
			}
			// if (target-sum)%v2 == 0 {
			// 	entries := fillValues(v2, (target-sum)/v2)
			// 	result = append(result, append(sumarr, entries...))
			// }
			if (target-sum)%v1 == 0 {
				entries := fillValues(v1, (target-sum)/v1)
				result = append(result, append(sumarr, entries...))
			}

			sum += v2
			sumarr = append(sumarr, v2)
		}
		if sum == target {
			result = append(result, sumarr)
		}
	}
	return result
}

func fillValues(v int, size int) []int {
	entry := make([]int, size)
	for i := 0; i < size; i++ {
		entry[i] = v
	}
	return entry
}

func fillValuesWithInitialValue(v, initial, target int) []int {
	entry := make([]int, (target/v)+1)
	entry[0] = initial
	for i := 1; i < len(entry); i++ {
		entry[i] = v
	}
	return entry
}

func combinationSum(candidates []int, target int) [][]int {
	ans := make([][]int, 0)
	curr := make([]int, 0)
	var backtrack func(idx int, currSum int, curr []int)
	backtrack = func(idx int, currSum int, curr []int) {
		if currSum == target {
			ans = append(ans, append([]int{}, curr...))
			return
		}
		if currSum > target {
			return
		}
		for i := idx; i < len(candidates); i++ {
			curr = append(curr, candidates[i])
			backtrack(i, currSum+candidates[i], curr)
			curr = curr[:len(curr)-1]
		}

	}
	backtrack(0, 0, curr)
	return ans
}

type combinator struct {
	level  []int
	result [][]int // store results
	leaf   string
}

// CombinationSumInternet taken from someone who wrote an interesting solution
func CombinationSumInternet(candidates []int, target int) [][]int {
	var partialResult []int // partial results
	// sort candidates to avoid going to values greater than target
	sort.Ints(candidates)
	aCombinator := new(combinator)
	aCombinator.combinationSumInternet(candidates, partialResult, target)

	return aCombinator.result
}

func (c *combinator) combinationSumInternet(candidates, partialResult []int, target int) {
	c.addLevel()
	c.printLevel(candidates, partialResult, target)
	// check if target is equal to zero after substracting the candidates
	if target == 0 {
		fmt.Println("found one", partialResult)
		c.result = append(c.result, partialResult)
	}

	// we return if candidates length is zero after removing its elements one by one.
	// Or check if targets value is less than candidates first value
	if len(candidates) == 0 || target < candidates[0] {
		c.exitLevel()
		return
	}

	// creates a new partialResult
	partialResult = partialResult[:len(partialResult):len(partialResult)]

	c.leaf = "1"
	c.combinationSumInternet(candidates, append(partialResult, candidates[0]), target-candidates[0])

	c.leaf = "2"
	c.combinationSumInternet(candidates[1:], partialResult, target)
	c.leaf = ""
	c.exitLevel()
}

const levelLabel = "level: %s-%s, parent: %s : "

func (c *combinator) exitLevel() {
	fmt.Println("exiting", fmt.Sprintf("%d-%s", c.currentLevel(), c.leaf))
	c.popLevel()
}

func (c *combinator) printLevel(candidates, partialResult []int, target int) {
	var currentLevel, parent string
	if len(c.level) == 0 {
		parent = "none"
		currentLevel = "none"
	}

	currentLevel = strconv.Itoa(c.level[len(c.level)-1])

	if len(c.level) == 1 {
		parent = "none"
	}

	if len(c.level) > 1 {
		parent = strconv.Itoa(c.level[len(c.level)-2])
	}
	fmt.Println(fmt.Sprintf(levelLabel, currentLevel, c.leaf, parent), "candidates", candidates, "partial", partialResult, "target", target, "leaf", c.leaf)
}

func (c *combinator) currentLevel() int {
	return c.level[len(c.level)-1]
}

func (c *combinator) addLevel() {
	if len(c.level) == 0 {
		c.level = append(c.level, 0)
		return
	}
	c.level = append(c.level, c.level[len(c.level)-1]+1)
}

func (c *combinator) popLevel() {
	c.level = c.level[:len(c.level)-1]
}
