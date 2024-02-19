package main

import (
	"sort"
)

// 1 https://leetcode.cn/problems/symmetric-tree/description/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var dfs func(*TreeNode, *TreeNode) bool
	dfs = func(node1, node2 *TreeNode) bool {
		if node1 == nil && node2 == nil {
			return true
		}
		if node1 == nil || node2 == nil || node1.Val != node2.Val {
			return false
		}
		return dfs(node1.Left, node2.Right) && dfs(node1.Right, node2.Left)
	}
	return dfs(root.Left, root.Right)
}

// 1 https://leetcode.cn/problems/symmetric-tree/description/
func isSymmetric_bfs(root *TreeNode) bool {
	if root == nil {
		return true
	}
	queue := []*TreeNode{root.Left, root.Right}
	for len(queue) > 0 {
		node1, node2 := queue[0], queue[1]
		queue = queue[2:]
		if node1 == nil && node2 == nil {
			continue
		}
		if node1 == nil || node2 == nil || node1.Val != node2.Val {
			return false
		}
		queue = append(queue, node1.Left, node2.Right, node1.Right, node2.Left)
	}
	return true
}

// 2 https://leetcode.cn/problems/merge-intervals/description/
func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	merged := make([][]int, 0)
	for _, interval := range intervals {
		if len(merged) == 0 || merged[len(merged)-1][1] < interval[0] {
			merged = append(merged, interval)
		} else {
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1],interval[1])
		}
	}
	return merged
}

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = 1e7
	}
	for _, v := range coins {
		if v <= amount {
			dp[v] = 1
		}
	}
	for i := 0; i <= amount; i++ {
		for _, v := range coins {
			if v < i {
				dp[i] = min(dp[i], dp[i-v]+1)
			}
		}
	}
	if dp[amount] == 1e7 {
		return -1
	}
	return dp[amount]
}

// O(n^2) dp
func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}

// O(nlogn)
func lengthOfLIS2(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	cnt := 0
	search := func(target int) int {
		left := 0
		right := cnt - 1
		for left <= right {
			mid := left + (right-left)/2
			v := dp[mid]
			if v == target {
				left = mid
				break
			} else if v < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		return left
	}
	for i := 0; i < n; i++ {
		index := search(nums[i])
		dp[index] = nums[i]
		if index >= cnt {
			cnt++
		}
	}
	return cnt
}

func maxProfit(prices []int) int {
	mini := prices[0]
	ans := 0
	for _, v := range prices {
		if v > mini {
			ans = max(ans, v-mini)
		}
		if v < mini {
			mini = v
		}
	}
	return ans
}
