package main

func maximalRectangle(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])
	dp := make([]int, n)
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				dp[j] = int(matrix[i][j]) - '0'
			} else {
				if matrix[i][j] == '0' {
					dp[j] = 0
				} else {
					dp[j] += 1
				}
			}
			ans = max(ans, dp[j])
		}
		ans = max(ans, largestRectangleArea(dp))
	}
	return ans
}

func largestRectangleArea(heights []int) int {
	ans := 0
	nums := []int{}
	nums = append(nums, 0)
	nums = append(nums, heights...)
	nums = append(nums, 0)
	stack := []int{}
	for i := 0; i < len(nums); i++ {
		for len(stack) > 0 && nums[i] <= nums[stack[len(stack)-1]] {
			h := nums[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			ans = max(ans, h*(i-stack[len(stack)-1]-1))
		}
		stack = append(stack, i)
	}
	return ans
}

// todo:dp
func findTargetSumWays(nums []int, target int) int {
	ans := 0
	n := len(nums)
	var dfs func(int, int)
	dfs = func(pos int, val int) {
		if pos == n {
			if val == 0 {
				ans++
			}
			return
		}
		dfs(pos+1, val-nums[pos])
		dfs(pos+1, val+nums[pos])
	}
	dfs(0, target)
	return ans
}

func rotate(matrix [][]int) {
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[n-i-1][j]
			matrix[n-i-1][j] = temp
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			temp := matrix[i][j]
			matrix[i][j] = matrix[j][i]
			matrix[j][i] = temp
		}
	}
}

func generateParenthesis(n int) []string {
	ans := []string{}
	check := func(nums []int) bool {
		stack := []int{}
		for _, v := range nums {
			if v == 0 {
				stack = append(stack, v)
			} else {
				if len(stack) > 0 {
					stack = stack[:len(stack)-1]
				} else {
					return false
				}
			}
		}
		return len(stack) == 0
	}
	var dfs func(int, []int)
	dfs = func(pos int, stack []int) {
		if pos >= 2*n {
			if !check(stack) {
				return
			}
			str := ""
			for _, v := range stack {
				if v == 0 {
					str += "("
				} else {
					str += ")"
				}
			}
			ans = append(ans, str)
			return
		}
		dfs(pos+1, append(stack, 0))
		dfs(pos+1, append(stack, 1))
	}
	dfs(1, []int{0})
	return ans
}
