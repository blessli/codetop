package main

import (
	"sort"
)

// 对称二叉树 https://leetcode.cn/problems/symmetric-tree/description/
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

// 对称二叉树 https://leetcode.cn/problems/symmetric-tree/description/
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

// 合并区间 https://leetcode.cn/problems/merge-intervals/description/
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
			merged[len(merged)-1][1] = max(merged[len(merged)-1][1], interval[1])
		}
	}
	return merged
}

// 零钱兑换 https://leetcode.cn/problems/coin-change/description/
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	dp := make([]int, amount+1)
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}
	for _, coin := range coins {
		for i := coin; i <= amount; i++ {
			dp[i] = min(dp[i], dp[i-coin]+1)
		}
	}
	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

// 最长上升子序列 https://leetcode.cn/problems/longest-increasing-subsequence/description/
func lengthOfLIS_dp(nums []int) int {
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

// 最长上升子序列 https://leetcode.cn/problems/longest-increasing-subsequence/description/
func lengthOfLIS(nums []int) int {
	n := len(nums)
	tails := make([]int, n)
	size := 0
	for _, num := range nums {
		l, r := 0, size
		for l < r {
			mid := l + (r-l)/2
			if tails[mid] < num {
				l = mid + 1
			} else {
				r = mid
			}
		}
		tails[l] = num
		if l == size {
			size++
		}
	}
	return size
}
// 最长公共子序列 https://leetcode.cn/problems/longest-common-subsequence/description/
func longestCommonSubsequence(text1 string, text2 string) int {
    m, n := len(text1), len(text2)
    dp := make([]int, n+1)
    var prev int

    for i := 1; i <= m; i++ {
        prev = 0
        for j := 1; j <= n; j++ {
            tmp := dp[j]
            if text1[i-1] == text2[j-1] {
                dp[j] = prev + 1
            } else {
                dp[j] = max(dp[j], dp[j-1])
            }
            prev = tmp
        }
    }
    return dp[n]
}

// 买卖股票的最佳时机 https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/description/
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	mini := prices[0]
	ans := 0
	for _, price := range prices {
		ans = max(ans, price-mini)
		mini = min(mini, price)
	}
	return ans
}
