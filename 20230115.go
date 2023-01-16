package main

import (
	"math"
)

func diameterOfBinaryTree(root *TreeNode) int {
	var dfs func(*TreeNode) int
	ans := 0
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		l := dfs(root.Left)
		r := dfs(root.Right)
		ans = max(ans, l+r)
		return max(l, r) + 1
	}
	dfs(root)
	return ans
}

func longestConsecutive(nums []int) int {
	m := map[int]bool{}
	for _, v := range nums {
		m[v] = true
	}
	ans := 0
	for _, v := range nums {
		if m[v-1] {
			continue
		}
		temp := v + 1
		count := 1
		for {
			if _, ok := m[temp]; ok {
				temp++
				count++
			} else {
				break
			}
		}
		ans = max(ans, count)
	}
	return ans
}

func searchMatrix(matrix [][]int, target int) bool {
	m := len(matrix)
	n := len(matrix[0])
	x := m - 1
	y := 0
	for x >= 0 && y < n {
		key := matrix[x][y]
		if key == target {
			return true
		}
		if key > target {
			x--
		} else {
			y++
		}
	}
	return false
}

func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	search := func(target int) int {
		left := 0
		right := len(nums) - 1
		for left < right {
			mid := left + (right-left)/2
			v := nums[mid]
			if v == target {
				right = mid
			} else if v < target {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
		return right
	}
	left := search(target)
	if left < 0 || left > len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	right := search(target + 1)
	if left <= right && nums[right] == nums[left] {
		return []int{left, right}
	}
	return []int{left, right - 1}
}

func isValidBST(root *TreeNode) bool {
	var dfs func(*TreeNode, int, int) bool
	dfs = func(root *TreeNode, low, high int) bool {
		if root == nil {
			return true
		}
		if root.Val <= low || root.Val >= high {
			return false
		}
		return dfs(root.Left, low, root.Val) && dfs(root.Right, root.Val, high)
	}
	return dfs(root.Left, math.MinInt64, root.Val) && dfs(root.Right, root.Val, math.MaxInt64)
}