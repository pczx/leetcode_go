package solution53

func search(nums []int, target int) int {
	return binarySearch(nums, target) - binarySearch(nums, target - 1)
}

func binarySearch(nums []int, target int) int {
	lo, hi := 0, len(nums) - 1
	for lo <= hi {
		mid := (lo + hi) >> 1
		if nums[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return lo
}