package solution57

func findContinuousSequence(target int) [][]int {
	ans := [][]int{}
	sum, limit := 0, (target-1)/2
	for i := 1; i <= limit; i++ {
		for j := i; ; j++ {
			sum += j
			if sum > target {
				sum = 0
				break
			} else if sum == target {
				list := []int{}
				for k := i; k <= j; k++ {
					list = append(list, k)
				}
				sum = 0
				ans = append(ans, list)
				break
			}
		}
	}
	return ans
}
