package solution59

func maxSlidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || k == 0 {
		return []int{}
	}
	d := []int{}
	ans := make([]int, len(nums)-k+1)
	for i := 0; i < k; i++ {
		for len(d) > 0 && d[len(d)-1] < nums[i] {
			d = d[:len(d)-1]
		}
		d = append(d, nums[i])
	}
	ans[0] = d[0]
	for i := k; i < len(nums); i++ {
		if d[0] == nums[i-k] {
			d = d[1:]
		}
		for len(d) > 0 && d[len(d)-1] < nums[i] {
			d = d[:len(d)-1]
		}
		d = append(d, nums[i])
		ans[i-k+1] = d[0]
	}
	return ans
}
