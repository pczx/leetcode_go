package solution39

func majorityElement(nums []int) int {
	start, mid, end := 0, len(nums)>>1, len(nums)-1
	index := partition(nums, start, end)
	for index != mid {
		if index > mid {
			end = index - 1

		} else {
			start = index + 1

		}

		index = partition(nums, start, end)
	}

	return nums[mid]
}

func partition(nums []int, lo, hi int) int {
	temp := nums[lo]
	for lo < hi {
		for lo < hi && nums[hi] > temp {
			hi--
		}

		if lo < hi {
			nums[lo] = nums[hi]
		}

		for lo < hi && nums[lo] <= temp {
			lo++
		}

		if lo < hi {
			nums[hi] = nums[lo]
		}
	}

	nums[lo] = temp
	return lo
}
