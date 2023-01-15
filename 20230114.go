package main

import (
	"fmt"
	"sort"
)

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var dfs func(*TreeNode, *TreeNode) bool
	dfs = func(root1, root2 *TreeNode) bool {
		if root1 == nil && root2 == nil {
			return true
		}
		if root1 == nil || root2 == nil {
			return false
		}
		if root1.Val != root2.Val {
			return false
		}
		return dfs(root1.Right, root2.Left) && dfs(root1.Left, root2.Right)
	}
	return dfs(root.Left, root.Right)
}

func merge(intervals [][]int) [][]int {
	if len(intervals) == 1 {
		return intervals
	}
	sort.Slice(intervals, func(i, j int) bool {
		if intervals[i][0] < intervals[j][0] {
			return true
		} else if intervals[i][0] == intervals[j][0] && intervals[i][1] < intervals[j][1] {
			return true
		}
		return false
	})
	fmt.Println(intervals)
	ans := make([][]int, 0)
	pre := intervals[0]
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= pre[1] {
			pre = []int{pre[0], max(intervals[i][1], pre[1])}
		} else {
			ans = append(ans, pre)
			pre = []int{intervals[i][0], intervals[i][1]}
		}
	}
	ans = append(ans, pre)
	return ans
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
