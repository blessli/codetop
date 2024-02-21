package main

// 最长回文子串 https://leetcode.cn/problems/longest-palindromic-substring/description/
func longestPalindrome(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, n)
		dp[i][i] = true
	}
	start, maxLen := 0, 1
	for l := 2; l <= n; l++ {
		for i := 0; i < n-l+1; i++ {
			j := i + l - 1
			if s[i] == s[j] && (l == 2 || dp[i+1][j-1]) {
				dp[i][j] = true
				if l > maxLen {
					start = i
					maxLen = l
				}
			}
		}
	}
	return s[start : start+maxLen]
}

// 最长回文子串 https://leetcode.cn/problems/longest-palindromic-substring/description/
func longestPalindrome2(s string) string {
	n := len(s)
	if n < 2 {
		return s
	}
	expandAroundCenter := func(left, right int) int {
		for left >= 0 && right < n && s[left] == s[right] {
			left--
			right++
		}
		return right - left - 1
	}
	start, end := 0, 0
	for i := 0; i < n; i++ {
		len1 := expandAroundCenter(i, i)
		len2 := expandAroundCenter(i, i+1)
		maxLen := max(len1, len2)
		if maxLen > end-start {
			start = i - (maxLen-1)/2
			end = i + maxLen/2
		}
	}
	return s[start : end+1]
}

// 不同路径 https://leetcode.cn/problems/unique-paths/description/
func uniquePaths(m int, n int) int {
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			dp[j] += dp[j-1]
		}
	}
	return dp[n-1]
}

// 不同路径 https://leetcode.cn/problems/unique-paths/description/
func uniquePaths2(m int, n int) int {
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

// 二叉树的中序遍历 https://leetcode.cn/problems/binary-tree-inorder-traversal/description/
func inorderTraversal(root *TreeNode) []int {
	var ans []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node != nil {
			dfs(node.Left)
			ans = append(ans, node.Val)
			dfs(node.Right)
		}
	}
	dfs(root)
	return ans
}

// 二叉树的中序遍历 https://leetcode.cn/problems/binary-tree-inorder-traversal/description/
func inorderTraversal2(root *TreeNode) []int {
	var ans []int
	stack := []*TreeNode{}
	curr := root
	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}
		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		ans = append(ans, curr.Val)
		curr = curr.Right
	}
	return ans
}

// 删除链表的倒数第 N 个结点 https://leetcode.cn/problems/remove-nth-node-from-end-of-list/description/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	fast, slow := dummy, dummy
	for i := 0; i < n+1; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dummy.Next
}

// 完全平方数 https://leetcode.cn/problems/perfect-squares/description/
func numSquares(n int) int {
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		dp[i] = i
		for j := 1; j*j <= i; j++ {
			dp[i] = min(dp[i], dp[i-j*j]+1)
		}
	}
	return dp[n]
}
