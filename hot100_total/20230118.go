package main

import "sort"

func canJump(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return true
	}
	m := n - 1
	for i := n - 2; i >= 0; i-- {
		if nums[i]+i >= m {
			m = i
		}
	}
	return m == 0
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	pre := map[int]*TreeNode{}
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		if root.Left != nil {
			pre[root.Left.Val] = root
		}
		if root.Right != nil {
			pre[root.Right.Val] = root
		}
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	vis := map[int]bool{}
	for p != nil {
		vis[p.Val] = true
		p = pre[p.Val]
	}
	for q != nil {
		if vis[q.Val] {
			return q
		}
		q = pre[q.Val]
	}
	return nil
}

// 搜索旋转排序数组
func search1(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		v := nums[mid]
		if v == target {
			return mid
		}
		if v > target {
			if target < nums[left] && v > nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		} else {
			if target > nums[right] && v < nums[left] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		}
	}
	return -1
}

func permute(nums []int) [][]int {
	n := len(nums)
	vis := make([]int, n)
	ans := [][]int{}
	var dfs func([]int, int)
	dfs = func(curr []int, depth int) {
		if depth == n {
			temp := make([]int, n)
			for i := 0; i < n; i++ {
				temp[i] = curr[i]
			}
			ans = append(ans, temp)
			return
		}
		for i := 0; i < n; i++ {
			if vis[i] == 0 {
				vis[i] = 1
				curr[depth] = nums[i]
				dfs(curr, depth+1)
				curr[depth] = 0
				vis[i] = 0
			}
		}
	}
	dfs(make([]int, n), 0)
	return ans
}

// O(1)空间复杂度
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}
