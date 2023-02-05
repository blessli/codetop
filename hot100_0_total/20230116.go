package main

import "math"

// 最长子串 dp 100ms
func longestPalindrome(s string) string {
	n := len(s)
	dp := make([][]int, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		dp[i][i] = 1
	}
	ans := []int{0, 0}
	maxi := 0
	for i := 1; i < n; i++ {
		for j := i - 1; j >= 0; j-- {
			if s[i] == s[j] {
				if j+1 == i {
					dp[j][i] = 2
				}
				if dp[j+1][i-1] > 0 {
					dp[j][i] = max(dp[j][i], dp[j+1][i-1]+2)
				}
				if dp[j][i] > maxi {
					maxi = dp[j][i]
					ans = []int{j, i}
				}
			}
		}
	}
	return s[ans[0] : ans[1]+1]
}

// 中心扩展算法 4ms
func longestPalindrome_(s string) string {
	n := len(s)
	dfs := func(pos int) [2]int {
		ans := [2]int{pos, pos}
		left, right, c := pos-1, pos+1, 1
		for left >= 0 && right < n && s[left] == s[right] {
			left--
			right++
			c += 2
		}
		if c > 1 {
			ans = [2]int{left + 1, right - 1}
		}
		left, right, c = pos-1, pos, 0
		for left >= 0 && right < n && s[left] == s[right] {
			left--
			right++
			c += 2
		}
		if c > ans[1]-ans[0]+1 {
			ans = [2]int{left + 1, right - 1}
		}
		return ans
	}
	ans := [2]int{0, 0}
	for i := 0; i < n; i++ {
		temp := dfs(i)
		if (temp[1] - temp[0] + 1) > ans[1]-ans[0]+1 {
			ans = [2]int{temp[0], temp[1]}
		}
	}
	return s[ans[0] : ans[1]+1]
}

// 不同路径 二维dp
func uniquePaths(m int, n int) int {
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[i][j] = 1
			} else {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}
	return dp[m-1][n-1]
}

// 一维dp 降低空间复杂度
func uniquePaths_(m int, n int) int {
	dp := make([]int, n)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 || j == 0 {
				dp[j] = 1
			} else {
				dp[j] += dp[j-1]
			}
		}
	}
	return dp[n-1]
}

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
func inorderTraversal_(root *TreeNode) []int {
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
