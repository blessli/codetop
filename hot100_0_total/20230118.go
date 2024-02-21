package main

// 跳跃游戏 https://leetcode.cn/problems/jump-game/description/
func canJump(nums []int) bool {
	n := len(nums)
	if n == 1 {
		return true
	}
	target := n - 1
	for i := n - 2; i >= 0; i-- {
		if nums[i]+i >= target { // 贪心算法
			target = i
		}
	}
	return target == 0
}

// 二叉树的最近公共祖先 https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/description/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	parent := map[*TreeNode]*TreeNode{}
	stack := []*TreeNode{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Left != nil {
			parent[node.Left] = node
			stack = append(stack, node.Left)
		}
		if node.Right != nil {
			parent[node.Right] = node
			stack = append(stack, node.Right)
		}
	}
	ancestors := map[*TreeNode]bool{}
	for p != nil {
		ancestors[p] = true
		p = parent[p]
	}
	for q != nil {
		if ancestors[q] {
			return q
		}
		q = parent[q]
	}
	return nil
}

// 搜索旋转排序数组 https://leetcode.cn/problems/search-in-rotated-sorted-array/description/
func search1(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] { // 左半部分是否有序
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] { // 右半部分是否有序
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
// 全排列 https://leetcode.cn/problems/permutations/description/
func permute(nums []int) [][]int {
	var res [][]int
	n := len(nums)
	var dfs func(start int)
	dfs = func(start int) {
		if start == n {
			perm := make([]int, n)
			copy(perm, nums)
			res = append(res, perm)
			return
		}
		for i := start; i < n; i++ {
			nums[start], nums[i] = nums[i], nums[start]
			dfs(start + 1)
			nums[start], nums[i] = nums[i], nums[start]
		}
	}
	dfs(0)
	return res
}

// 多数元素 https://leetcode.cn/problems/majority-element/description/
func majorityElement(nums []int) int {
	candidate := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if count == 0 {
			candidate = nums[i]
			count = 1
		} else if candidate == nums[i] {
			count++
		} else {
			count--
		}
	}
	return candidate
}
