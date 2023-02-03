package main

import (
	"strings"
)

// 相交链表
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	tempA := headA
	tempB := headB
	for headA != nil || headB != nil {
		if headA == headB {
			return headA
		}
		if headA == nil {
			headA = tempB
		} else {
			headA = headA.Next
		}
		if headB == nil {
			headB = tempA
		} else {
			headB = headB.Next
		}
	}
	return nil
}

// 回文链表
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	length := func(head *ListNode) int {
		n := 0
		temp := head
		for temp != nil {
			temp = temp.Next
			n++
		}
		return n
	}
	midNode := func(head *ListNode) *ListNode {
		pre := head
		last := head.Next.Next
		for last != nil && last.Next != nil {
			pre = pre.Next
			last = last.Next.Next
		}
		return pre
	}
	pre := midNode(head)
	back := pre.Next
	if length(head)%2 == 1 {
		back = back.Next
	}
	pre.Next = nil
	front := reverseList(head)
	for front != nil && back != nil {
		if front.Val != back.Val {
			return false
		}
		front = front.Next
		back = back.Next
	}
	return true
}

// 反转链表 递归版
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	node := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return node
}

// 迭代版
func reverseList1(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ans := &ListNode{}
	pre := ans.Next
	for head != nil {
		temp := head.Next
		head.Next = pre
		pre = head
		head = temp
	}
	return pre
}

// 单词搜索 递归
func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	var dfs func(int, int, int) bool
	dfs = func(x, y, k int) bool {
		if k >= len(word) {
			return true
		}
		if x < 0 || x >= m || y < 0 || y >= n || board[x][y] != word[k] {
			return false
		}
		board[x][y] += 'a' // 打上标记，降低空间复杂度
		result := dfs(x-1, y, k+1) || dfs(x+1, y, k+1) || dfs(x, y-1, k+1) || dfs(x, y+1, k+1)
		board[x][y] -= 'a'
		return result
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

// 单词拆分 记忆化搜索 4ms
func wordBreak_(s string, wordDict []string) bool {
	count := len(wordDict)
	ans := false
	vis := map[string]bool{}
	var dfs func(string)
	dfs = func(curr string) {
		if ans {
			return
		}
		if curr == s {
			ans = true
			return
		}
		for i := 0; i < count; i++ {
			now := curr + wordDict[i]
			if len(now) <= len(s) && strings.HasPrefix(s, now) && !vis[now] {
				vis[now] = true
				dfs(now)
			}
		}
	}
	dfs("")
	return ans
}
// dp 0ms
func wordBreak(s string, wordDict []string) bool {
	dict := map[string]bool{}
	for _, word := range wordDict {
		dict[word] = true
	}
	n := len(s)
	dp := make([]bool, n+1)// 前i个字符可被拆分
	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && dict[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}
