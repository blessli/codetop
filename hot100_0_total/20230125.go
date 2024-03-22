package main

// 最大矩形 https://leetcode.cn/problems/maximal-rectangle/description/
func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	rows, cols := len(matrix), len(matrix[0])
	heights := make([]int, cols)
	maxArea := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if matrix[i][j] == '1' {
				heights[j]++
			} else {
				heights[j] = 0
			}
		}
		area := largestRectangleArea(heights)
		if area > maxArea {
			maxArea = area
		}
	}

	return maxArea
}

// 柱状图中最大的矩形 https://leetcode.cn/problems/largest-rectangle-in-histogram/description/
func largestRectangleArea(heights []int) int {
	n := len(heights)
	stack := make([]int, 0)
	maxArea := 0

	for i := 0; i <= n; i++ {
		var h int
		if i < n {
			h = heights[i]
		} else {
			h = 0
		}

		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
			height := heights[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			width := i
			if len(stack) > 0 {
				width = i - stack[len(stack)-1] - 1
			}
			area := height * width
			if area > maxArea {
				maxArea = area
			}
		}

		stack = append(stack, i)
	}

	return maxArea
}

// 目标和 https://leetcode.cn/problems/target-sum/description/
func findTargetSumWays(nums []int, target int) int {
	ans := 0
	n := len(nums)
	var dfs func(int, int)
	dfs = func(index int, sum int) {
		if index == n {
			if sum == target {
				ans++
			}
			return
		}
		dfs(index+1, sum+nums[index])
		dfs(index+1, sum-nums[index])
	}
	dfs(0, 0)
	return ans
}

// 轮转数组 https://leetcode.cn/problems/rotate-array/description/
func rotate(matrix [][]int) {
	n := len(matrix)

	// 先对矩阵进行转置操作，即将矩阵沿着主对角线（左上-右下）进行镜像对称
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 再对每一行进行水平翻转，即将每一行的元素进行左右交换
	for i := 0; i < n; i++ {
		for j := 0; j < n/2; j++ {
			matrix[i][j], matrix[i][n-1-j] = matrix[i][n-1-j], matrix[i][j]
		}
	}
}


// 括号生成 https://leetcode.cn/problems/generate-parentheses/description/
func generateParenthesis(n int) []string {
	var result []string
	var backtrack func(s string, left, right int)

	backtrack = func(s string, left, right int) {
		if len(s) == 2*n {
			result = append(result, s)
			return
		}

		if left < n {
			backtrack(s+"(", left+1, right)
		}

		if right < left {
			backtrack(s+")", left, right+1)
		}
	}

	backtrack("", 0, 0)

	return result
}
