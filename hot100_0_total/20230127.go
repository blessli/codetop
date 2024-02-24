package main

// 合并二叉树 https://leetcode.cn/problems/merge-two-binary-trees/description/
func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	merged := &TreeNode{Val: root1.Val + root2.Val}
	merged.Left = mergeTrees(root1.Left, root2.Left)
	merged.Right = mergeTrees(root1.Right, root2.Right)
	return merged
}

// 排序链表 https://leetcode.cn/problems/sort-list/description/
func sortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var prev *ListNode
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		prev = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	prev.Next = nil
	left := sortList(head)
	right := sortList(slow)
	return mergeTwoLists(left, right)
}

// 最长有效括号 https://leetcode.cn/problems/longest-valid-parentheses/description/
func longestValidParentheses(s string) int {
	if len(s) == 0 {
		return 0
	}
	stack := []int{-1}
	maxLen := 0

	for i := 0; i < len(s); i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		} else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack, i)
			} else {
				maxLen = max(maxLen, i-stack[len(stack)-1])
			}
		}
	}

	return maxLen
}

// 回文子串 https://leetcode.cn/problems/palindromic-substrings/description/
func countSubstrings(s string) int {
	extendPalindrome := func(s string, left, right int) int {
		count := 0
		for left >= 0 && right < len(s) && s[left] == s[right] {
			count++
			left--
			right++
		}
		return count
	}
	count := 0
	for i := 0; i < len(s); i++ {
		count += extendPalindrome(s, i, i)
		count += extendPalindrome(s, i, i+1)
	}
	return count
}

// 买卖股票的最佳时机 II https://leetcode.cn/problems/best-time-to-buy-and-sell-stock-ii/
func maxProfit2(prices []int) int {
	n := len(prices)
	dp := make([][3]int, n)
	// f[i][0]: 手上持有股票的最大收益
	// f[i][1]: 手上不持有股票，并且处于冷冻期中的累计最大收益
	// f[i][2]: 手上不持有股票，并且不在冷冻期中的累计最大收益
	dp[0][0] = -prices[0]
	dp[0][1] = 0
	dp[0][2] = 0
	for i := 1; i < n; i++ {
		dp[i][0] = max(dp[i-1][0], dp[i-1][2]-prices[i])
		dp[i][1] = prices[i] + dp[i-1][0]
		dp[i][2] = max(dp[i-1][1], dp[i-1][2])
	}
	return max(dp[n-1][1], dp[n-1][2])
}
