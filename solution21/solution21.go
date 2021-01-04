package solution21

func exchange(nums []int) []int {
	lo, hi := 0, len(nums)-1
	for lo < hi {
		for lo < hi && nums[lo]&1 != 0 {
			lo++
		}

		for lo < hi && nums[hi]&1 == 0 {
			hi--
		}

		if lo < hi {
			nums[lo], nums[hi] = nums[hi], nums[lo]
		}
	}
	return nums
}
