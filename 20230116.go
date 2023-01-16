package main

import "math"

func inorderTraversal(root *TreeNode) []int {
	ans := []int{}
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root != nil {
			dfs(root.Left)
			ans = append(ans, root.Val)
			dfs(root.Right)
		}
	}
	dfs(root)
	return ans
}

// 迭代法
func inorderTraversal2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := []int{}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)
	for len(stack) > 0 && root != nil {
		for root != nil {
			if root.Left != nil {
				stack = append(stack, root.Left)
			}
			root = root.Left
		}
		root = stack[len(stack)-1]
		ans = append(ans, root.Val)
		stack = stack[:len(stack)-1]
		if root != nil {
			if root.Right != nil {
				stack = append(stack, root.Right)
			}
			root = root.Right
		}
	}
	return ans
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	temp := head
	pre := head
	last := head
	count := 0
	for head != nil {
		head = head.Next
		count++
	}
	if count == 1 {
		return nil
	}
	if count == n {
		return temp.Next
	}
	for i := 0; i < n; i++ {
		pre = pre.Next
	}
	for pre != nil && pre.Next != nil {
		pre = pre.Next
		last = last.Next
	}
	if last.Next != nil {
		last.Next = last.Next.Next
	}
	return temp
}

func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 0; i <= n; i++ {
		dp[i] = math.MaxInt32
	}
	dp[0] = 0
	for i := 1; i <= n; i++ {
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}
