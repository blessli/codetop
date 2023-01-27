package main

func mergeTrees(root1 *TreeNode, root2 *TreeNode) *TreeNode {
	if root1 == nil && root2 == nil {
		return nil
	}
	if root1 == nil {
		return root2
	}
	if root2 == nil {
		return root1
	}
	root := &TreeNode{}
	if root1 != nil {
		root.Val += root1.Val
	}
	if root2 != nil {
		root.Val += root2.Val
	}
	root.Left = mergeTrees(root1.Left, root2.Left)
	root.Right = mergeTrees(root1.Right, root2.Right)
	return root
}

func sortList(head *ListNode) *ListNode {
	merge := func(l1, l2 *ListNode) *ListNode {
		ans := &ListNode{}
		res, left, right := ans, l1, l2
		for left != nil && right != nil {
			if left.Val <= right.Val {
				ans.Next = left
				left = left.Next
			} else {
				ans.Next = right
				right = right.Next
			}
			ans = ans.Next
		}
		if left != nil {
			ans.Next = left
		}
		if right != nil {
			ans.Next = right
		}
		return res.Next
	}
	findMid := func(head *ListNode) *ListNode {
		left := head
		right := head.Next.Next
		for right != nil && right.Next != nil {
			left = left.Next
			right = right.Next.Next
		}
		return left
	}
	var dfs func(*ListNode) *ListNode
	dfs = func(head *ListNode) *ListNode {
		if head == nil || head.Next == nil {
			return head
		}
		mid := findMid(head)
		left := head
		right := mid.Next
		mid.Next = nil
		return merge(dfs(left), dfs(right))
	}
	return dfs(head)
}

func longestValidParentheses(s string) int {
	n := len(s)
	stack := []int{}
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		if s[i] == '(' {
			stack = append(stack, i)
		}
		if s[i] == ')' {
			if len(stack) == 0 {
				dp[i] = 1
				continue
			}
			if len(stack) > 0 && s[stack[len(stack)-1]] == '(' {
				stack = stack[:len(stack)-1]
			}
		}
	}
	for len(stack) > 0 {
		dp[stack[len(stack)-1]] = 1
		stack = stack[:len(stack)-1]
	}
	ans := 0
	pc := 0
	for i := 0; i < n; i++ {
		if dp[i] == 1 {
			ans = max(ans, pc)
			pc = 0
		} else {
			pc++
		}
	}
	ans = max(ans, pc)
	return ans
}

func countSubstrings(s string) int {
	n := len(s)
	dfs := func(pos int) int {
		ans := 1
		left, right := pos-1, pos+1
		for left >= 0 && right < n && s[left] == s[right] {
			ans++
			left--
			right++
		}
		left, right = pos-1, pos
		for left >= 0 && right < n && s[left] == s[right] {
			ans++
			left--
			right++
		}
		return ans
	}
	ans := 0
	for i := 0; i < n; i++ {
		ans += dfs(i)
	}
	return ans
}

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
	return max(dp[n-1][1],dp[n-1][2])
}
