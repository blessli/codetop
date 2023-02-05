package main

import (
	"sort"
)

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

func productExceptSelf(nums []int) []int {
	n := len(nums)
	ans := make([]int, n)
	for i := 0; i < n; i++ {
		ans[i] = 1
	}
	left, right := 1, 1
	for i := 0; i < n; i++ {
		ans[i] *= left
		left *= nums[i]
		ans[n-i-1] *= right
		right *= nums[n-i-1]
	}
	return ans
}

func convertBST(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	num := 0
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Right)
		root.Val += num
		num = root.Val
		dfs(root.Left)
	}
	dfs(root)
	return root
}

func maxSlidingWindow(nums []int, k int) []int {
	ans := []int{}
	stack := []int{}
	for i := 0; i < k; i++ {
		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}
		for len(stack) > 0 && nums[i] >= nums[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	ans = append(ans, nums[stack[0]])
	for i := k; i < len(nums); i++ {
		if stack[0] <= i-k {
			stack = stack[1:]
		}
		for len(stack) > 0 && nums[i] >= nums[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
		ans = append(ans, nums[stack[0]])
	}
	return ans
}

func trap(height []int) int {
	ans := 0
	stack := []int{}
	for i := 0; i < len(height); i++ {
		if len(stack) == 0 {
			stack = append(stack, i)
			continue
		}
		for len(stack) > 0 && height[i] >= height[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			ans += (min(height[left], height[i]) - height[top]) * (i - left - 1)
		}
		stack = append(stack, i)
	}
	return ans
}
