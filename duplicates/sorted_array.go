package duplicates

type remover struct {
	idx        int
	last       int
	duplicates map[int]bool
	nums       []int
}

func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	remover := remover{
		last:       len(nums) - 1,
		duplicates: make(map[int]bool),
		nums:       nums,
	}
	remover.remove()
	return remover.last + 1
}

func (r *remover) remove() {
	if r.idx > r.last {
		return
	}
	v := r.nums[r.idx]
	_, ok := r.duplicates[v]
	if !ok {
		r.duplicates[v] = true
		r.idx++
		r.remove()
		return
	}
	var i int
	if r.idx > 0 {
		i = r.idx - 1
	}
	r.nums = append(r.nums[:i], r.nums[i+1:]...)
	r.nums = append(r.nums, 0)
	r.last--
	r.remove()
}

func RemoveDuplicatesInternet(nums []int) int {
	left, right, size := 0, 1, len(nums)
	for ; right < size; right++ {
		if nums[left] == nums[right] {
			continue
		}
		left++
		nums[left], nums[right] = nums[right], nums[left]
	}
	return left + 1
}
