package main

import "math"

// 最长回文子串
func longestPalindrome(s string) string {
	dp_ := func(s string) string {
		n := len(s)
		dp := make([][]int, n)
		for i := 0; i < n; i++ {
			dp[i] = make([]int, n)
			dp[i][i] = 1
		}
		ans := [2]int{0, 0}
		for i := 0; i < n; i++ {
			for j := i - 1; j >= 0; j-- {
				if s[i] == s[j] {
					if j+1 == i {
						dp[j][i] = max(dp[j][i], 2)
						if dp[j][i] > ans[1]-ans[0] {
							ans = [2]int{j, i}
						}
					} else if dp[j+1][i-1] > 0 {
						dp[j][i] = max(dp[j][i], dp[j+1][i-1]+2)
						if dp[j][i] > ans[1]-ans[0] {
							ans = [2]int{j, i}
						}
					}
				}
			}
		}

		return s[ans[0] : ans[1]+1]
	}
	return dp_(s)
}

// 打家劫舍 O(n)即可
func rob(nums []int) int {
	dp_ := func(nums []int) (ans int) {
		n := len(nums)
		if n == 1 {
			return nums[0]
		}
		if n == 2 {
			return max(nums[0], nums[1])
		}
		dp := make([]int, n)
		dp[0] = nums[0]
		dp[1] = max(nums[0], nums[1])
		for i := 2; i < n; i++ {
			dp[i] = max(dp[i-1], dp[i-2]+nums[i])
		}
		return dp[n-1]
	}
	return dp_(nums)
}

// 最大正方形
func maximalSquare(matrix [][]byte) int {
	dp_ := func(matrix [][]byte) (ans int) {
		m := len(matrix)
		n := len(matrix[0])
		dp := make([][]int, m)
		for i := 0; i < m; i++ {
			dp[i] = make([]int, n)
		}
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if i == 0 || j == 0 {
					dp[i][j] = int(matrix[i][j]) - '0'
				} else if matrix[i][j] == '1' {
					dp[i][j] = min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j])) + 1
				}
				ans = max(ans, dp[i][j]*dp[i][j])
			}
		}
		return ans
	}
	return dp_(matrix)
}

// 不同路径
func uniquePaths(m int, n int) int {
	dp_ := func(m, n int) int {
		dp := make([]int, n)
		for i := 0; i < n; i++ {
			dp[i] = 1
		}
		for i := 1; i < m; i++ {
			for j := 1; j < n; j++ {
				dp[j] += dp[j-1]
			}
		}
		return dp[n-1]
	}
	return dp_(m, n)
}

// 最小路径和
func minPathSum(grid [][]int) int {
	dp_ := func(grid [][]int) int {
		m := len(grid)
		n := len(grid[0])
		dp := make([][]int, m)
		for i := 0; i < m; i++ {
			dp[i] = make([]int, n)
		}
		for i := 0; i < m; i++ {
			for j := 0; j < n; j++ {
				if i == 0 && j == 0 {
					dp[i][j] = grid[i][j]
					continue
				}
				if i == 0 {
					dp[i][j] = dp[i][j-1] + grid[i][j]
					continue
				}
				if j == 0 {
					dp[i][j] = dp[i-1][j] + grid[i][j]
					continue
				}
				dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + grid[i][j]
			}
		}
		return dp[m-1][n-1]
	}
	return dp_(grid)
}

// 完全平方数
func numSquares(n int) int {
	dp_ := func(n int) int {
		dp := make([]int, n+1)
		for i := 1; i <= n; i++ {
			dp[i] = n
		}
		for i := 1; i <= n; i++ {
			for j := 1; j*j <= i; j++ {
				dp[i] = min(dp[i], dp[i-j*j]+1)
			}
		}
		return dp[n]
	}
	return dp_(n)
}

// 零钱兑换
func coinChange(coins []int, amount int) int {
	dp_ := func(coins []int, amount int) int {
		dp := make([]int, amount+1)
		for i := 1; i <= amount; i++ {
			dp[i] = math.MaxInt32
		}
		for i := 1; i <= amount; i++ {
			for _, v := range coins {
				if i >= v {
					dp[i] = min(dp[i], dp[i-v]+1)
				}
			}
		}
		if dp[amount] == math.MaxInt32 {
			return -1
		}
		return dp[amount]
	}
	return dp_(coins, amount)
}

// 单词拆分
func wordBreak(s string, wordDict []string) bool {
	m := map[string]bool{}
	n := len(s)
	for _, v := range wordDict {
		m[v] = true
	}
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	var dp_ func(int, int) bool
	dp_ = func(left, right int) bool {
		if m[s[left:right+1]] {
			return true
		}
		if left+1 > right {
			return false
		}
		if dp[left][right] != 0 {
			return dp[left][right] == 2
		}

		result := false
		for i := left; i < right; i++ {
			val := dp_(left, i) && dp_(i+1, right)
			result = result || val
		}
		if result {
			dp[left][right] = 2
		} else {
			dp[left][right] = 1
		}
		return result
	}
	return dp_(0, n-1)
}

// 最长递增子序列
func lengthOfLIS(nums []int) int {
	dpMethod := func(nums []int) int {
		n := len(nums)
		dp := make([]int, n)
		for i := 0; i < n; i++ {
			dp[i] = 1
		}
		ans := 1
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
	search := func(nums []int, target, n int) int {
		left, right := 0, n-1
		ans := 0
		for left <= right {
			mid := left + (right-left)/2
			if nums[mid] == target {
				return mid
			}
			if nums[mid] < target {
				left = mid + 1
				ans = mid + 1
			} else {
				right = mid - 1
			}
		}
		return ans
	}
	binaryMethod := func(nums []int) int {
		n := len(nums)
		dp := make([]int, n)
		count := 1
		dp[0] = nums[0]
		ans := 1
		for i := 1; i < n; i++ {
			idx := search(dp, nums[i], count)
			if idx >= count {
				dp[count] = nums[i]
				count++
				ans = max(ans, count)
			} else {
				dp[idx] = nums[i]
			}

		}
		return ans
	}
	dpMethod(nums)
	return binaryMethod(nums)
}

// 乘积最大子数组
func maxProduct(nums []int) int {
	n := len(nums)
	mindp := make([]int, 0)
	mindp = append(mindp, nums...)
	maxdp := make([]int, 0)
	maxdp = append(maxdp, nums...)
	ans := nums[0]
	for i := 1; i < n; i++ {
		maxdp[i] = max(maxdp[i-1]*nums[i], max(nums[i], mindp[i-1]*nums[i]))
		mindp[i] = min(mindp[i-1]*nums[i], min(nums[i], maxdp[i-1]*nums[i]))
		ans = max(ans, maxdp[i])
	}
	return ans
}

// 买卖股票的最佳时机
func maxProfit(prices []int) int {
	pre := prices[0]
	ans := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > pre {
			ans = max(ans, prices[i]-pre)
		}
		pre = min(pre, prices[i])
	}
	return ans
}

// 最大子数组和 todo: 分治法
func maxSubArray(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = nums[0]
	ans := nums[0]
	for i := 1; i < n; i++ {
		dp[i] = max(dp[i-1]+nums[i], nums[i])
		ans = max(ans, dp[i])
	}
	return ans
}

// 爬楼梯
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}
	dp := make([]int, n+1)
	dp[0] = 0
	dp[1] = 1
	dp[2] = 2
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
