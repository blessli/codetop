package main

import (
	"math"
)

// 二叉树的直径 https://leetcode.cn/problems/diameter-of-binary-tree/description/
func diameterOfBinaryTree(root *TreeNode) int {
	var dfs func(*TreeNode) int
	maxDiameter := 0
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		leftDepth := dfs(root.Left)
		rightDepth := dfs(root.Right)
		maxDiameter = max(maxDiameter, leftDepth+rightDepth) // 在计算深度的同时更新直径的最大值
		return max(leftDepth, rightDepth) + 1
	}
	dfs(root)
	return maxDiameter
}

// 最长连续序列 https://leetcode.cn/problems/longest-consecutive-sequence/description/
func longestConsecutive(nums []int) int {
	numSet := map[int]bool{}
	for _, v := range nums {
		numSet[v] = true
	}
	longestStreak := 0
	for _, num := range nums {
		if !numSet[num-1] { // 检查它是否为连续序列的起点
			currentNum := num
			currentStreak := 1
			for numSet[currentNum+1] {
				currentNum++
				currentStreak++
				delete(numSet, currentNum) // 删除已经访问过的数字，避免重复计算
			}
			longestStreak = max(longestStreak, currentStreak)
		}
	}
	return longestStreak
}

// 搜索二维矩阵 II https://leetcode.cn/problems/search-a-2d-matrix-ii/description/
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

// 在排序数组中查找元素的第一个和最后一个位置 https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/description/
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	search := func(target int) int {
		left, right := 0, len(nums)
		for left < right {
			mid := left + (right-left)/2
			if nums[mid] < target {
				left = mid + 1
			} else {
				right = mid
			}
		}
		return right
	}
	left := search(target)
	if left >= len(nums) || nums[left] != target {
		return []int{-1, -1}
	}
	right := search(target+1) - 1
	return []int{left, right}
}

// 验证二叉搜索树 https://leetcode.cn/problems/validate-binary-search-tree/description/
func isValidBST(root *TreeNode) bool {
	var dfs func(*TreeNode, int, int) bool
	dfs = func(node *TreeNode, min, max int) bool {
		if node == nil {
			return true
		}
		if node.Val <= min || node.Val >= max {
			return false
		}
		return dfs(node.Left, min, node.Val) && dfs(node.Right, node.Val, max)
	}
	return dfs(root, math.MinInt64, math.MaxInt64)
}

// 验证二叉搜索树 https://leetcode.cn/problems/validate-binary-search-tree/description/
func isValidBST_bfs(root *TreeNode) bool {
	if root == nil {
		return true
	}
	type NodeInfo struct {
		Node  *TreeNode
		Lower int
		Upper int
	}
	queue := []*NodeInfo{{root, math.MinInt64, math.MaxInt64}}
	for len(queue) > 0 {
		nodeInfo := queue[0]
		queue = queue[1:]
		node := nodeInfo.Node
		lower, upper := nodeInfo.Lower, nodeInfo.Upper
		if node.Val <= lower || node.Val >= upper {
			return false
		}
		if node.Left != nil {
			queue = append(queue, &NodeInfo{node.Left, lower, node.Val})
		}
		if node.Right != nil {
			queue = append(queue, &NodeInfo{node.Right, node.Val, upper})
		}
	}
	return true
}
