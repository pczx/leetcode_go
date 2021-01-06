package solution3

func findRepeatNumber(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}

	for _, value := range nums {
		if value < 0 || value > len(nums)-1 {
			return 0
		}
	}

	for i := 0; i < len(nums); i++ {
		for ; nums[i] != i; nums[i], nums[nums[i]] = nums[nums[i]], nums[i] {
			if nums[i] == nums[nums[i]] {
				return nums[i]
			}
		}
	}
	return 0
}
