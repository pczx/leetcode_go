package main

import "math"

func canJump(nums []int) bool {
	n := len(nums)
	for i := 0; i < n; {
		cnt := nums[i]
		maxStep := math.MinInt32
		index := 0
		for j := 1; j <= cnt; j++ {
			if i+j < n && nums[i+j]+j > maxStep {
				maxStep = nums[i+j] + j
				index = j
			}
		}
		i += index
		if i == n-1 {
			return true
		}
		if nums[i] == 0 {
			return false
		}
	}
	return true
}
