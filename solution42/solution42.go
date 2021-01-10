package solution42

import (
	"math"
)

func maxSubArray(nums []int) int {
	curSum, maxSum := 0, math.MinInt64

	for _, v := range nums {
		if curSum <= 0 {
			curSum = v
		} else {
			curSum += v
		}

		if curSum > maxSum {
			maxSum = curSum
		}
	}
	return maxSum
}
