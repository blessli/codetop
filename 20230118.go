package main

import "sort"

func permute(nums []int) [][]int {
	n := len(nums)
	vis := make([]int, n)
	ans := [][]int{}
	var dfs func([]int, int)
	dfs = func(curr []int, depth int) {
		if depth == n {
			temp := make([]int, n)
			for i := 0; i < n; i++ {
				temp[i] = curr[i]
			}
			ans = append(ans, temp)
			return
		}
		for i := 0; i < n; i++ {
			if vis[i] == 0 {
				vis[i] = 1
				curr[depth] = nums[i]
				dfs(curr, depth+1)
				curr[depth] = 0
				vis[i] = 0
			}
		}
	}
	dfs(make([]int, n), 0)
	return ans
}

// O(1)空间复杂度
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
