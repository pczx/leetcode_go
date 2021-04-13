package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	ans := [][]int{}
	cnt := []int{}
	sort.Ints(nums)
	var dfs func(index int)
	dfs = func(index int) {
		ans = append(ans, cnt)
		for i := index; i < len(nums); i++ {
			if i > index && nums[i] == nums[i-1] {
				continue
			}
			cnt = append(cnt, nums[i])
			dfs(i + 1)
			cnt = cnt[:len(cnt)-1]
		}
	}
	dfs(0)
	return ans
}
