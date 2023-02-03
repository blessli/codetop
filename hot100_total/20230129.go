package main

import "sort"

func rob2(root *TreeNode) int {
	var dfs func(*TreeNode) [2]int
	ans := 0
	dfs = func(root *TreeNode) [2]int {
		if root == nil {
			return [2]int{0, 0}
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		res := [2]int{}
		res[0] = root.Val + left[1] + right[1]
		res[1] = max(left[0], left[1]) + max(right[0], right[1])
		ans = max(ans, max(res[0], res[1]))
		return res
	}
	dfs(root)
	return ans
}

func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || (a[0] == b[0] && a[1] < b[1])
	})
	ans := people
	for _, v := range people {
		idx := v[1]
		ans = append(ans[:idx], append([][]int{v}, ans[idx:]...)...)
	}
	return ans
}

func maxCoins(nums []int) int {
	n := len(nums)
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	dp := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		dp[i] = make([]int, n+2)
	}
	for i := 0; i < n+2; i++ {
		for j := 0; j < n+2; j++ {
			dp[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(left, right int) int {
		if left >= right-1 {
			return 0
		}
		if dp[left][right] != -1 {
			return dp[left][right]
		}
		for i := left + 1; i < right; i++ {
			v := nums[left] * nums[i] * nums[right]
			v += dfs(left, i) + dfs(i, right)
			dp[left][right] = max(dp[left][right], v)
		}
		return dp[left][right]
	}
	return dfs(0, n+1)
}
