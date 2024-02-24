package main

import "sort"

// 打家劫舍 III https://leetcode.cn/problems/house-robber-iii/description/
func rob3(root *TreeNode) int {
	var robSub func(node *TreeNode) [2]int
	robSub = func(node *TreeNode) [2]int {
		if node == nil {
			return [2]int{0, 0}
		}

		left := robSub(node.Left)
		right := robSub(node.Right)

		result := [2]int{0, 0}
		result[0] = max(left[0], left[1]) + max(right[0], right[1])
		result[1] = node.Val + left[0] + right[0]

		return result
	}
	result := robSub(root)
	return max(result[0], result[1])
}

// 根据身高重建队列 https://leetcode.cn/problems/queue-reconstruction-by-height/description/
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

// 戳气球 https://leetcode.cn/problems/burst-balloons/description/
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
