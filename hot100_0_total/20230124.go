package main

import (
	"sort"
)

// 下一个排列 https://leetcode.cn/problems/next-permutation/description/
func nextPermutation(nums []int) {
	pos := -1
	n := len(nums)
	for i := n - 1; i >= 0 && pos < 0; i-- {
		for j := n - 1; j > i && pos < 0; j-- {
			if nums[i] < nums[j] {
				temp := nums[i]
				nums[i] = nums[j]
				nums[j] = temp
				pos = i
				break
			}
		}
	}
	if pos < 0 {
		for i := 0; i < n/2; i++ {
			temp := nums[i]
			nums[i] = nums[n-i-1]
			nums[n-i-1] = temp
		}
		return
	}
	sort.Ints(nums[pos+1:])
}

// 除自身以外数组的乘积 https://leetcode.cn/problems/product-of-array-except-self/description/
func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)
	answer[0] = 1
	for i := 1; i < n; i++ {
		answer[i] = answer[i-1] * nums[i-1]
	}
	right := 1
	for i := n - 1; i >= 0; i-- {
		answer[i] *= right
		right *= nums[i]
	}
	return answer
}

// 把二叉搜索树转换为累加树 https://leetcode.cn/problems/convert-bst-to-greater-tree/description/
func convertBST(root *TreeNode) *TreeNode {
	sum := 0
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Right)
		sum += node.Val
		node.Val = sum
		dfs(node.Left)
	}
	dfs(root)
	return root
}

// 滑动窗口最大值 https://leetcode.cn/problems/sliding-window-maximum/description/
func maxSlidingWindow(nums []int, k int) []int {
	n := len(nums)
	if n == 0 {
		return []int{}
	}
	result := make([]int, 0)
	deque := make([]int, 0)
	for i := 0; i < n; i++ {
		if len(deque) > 0 && deque[0] < i-k+1 {
			deque = deque[1:]
		}
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[i] {
			deque = deque[:len(deque)-1]
		}
		deque = append(deque, i)
		if i >= k-1 {
			result = append(result, nums[deque[0]])
		}
	}
	return result
}

// 接雨水 https://leetcode.cn/problems/trapping-rain-water/description/
func trap(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}
	leftMax := make([]int, n)
	rightMax := make([]int, n)
	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}
	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}
	result := 0
	for i := 0; i < n; i++ {
		result += min(leftMax[i], rightMax[i]) - height[i]
	}
	return result
}
