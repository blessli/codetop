package main

import (
	"strings"
)

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

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	n := 0
	temp := head
	for temp != nil {
		temp = temp.Next
		n++
	}
	pre := head
	last := head.Next.Next
	for last != nil && last.Next != nil {
		pre = pre.Next
		last = last.Next.Next
	}
	back := pre.Next
	if n%2 == 1 {
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

// 递归版
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

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	ans := false
	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}
	count := len(word)
	dist := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var dfs func(int, int, int)
	dfs = func(x, y, pos int) {
		if ans {
			return
		}
		if pos == count {
			ans = true
			return
		}
		for i := 0; i < 4; i++ {
			xx := x + dist[i][0]
			yy := y + dist[i][1]
			if xx >= 0 && xx < m && yy >= 0 && yy < n && pos < count && board[xx][yy] == word[pos] && !vis[xx][yy] {
				vis[xx][yy] = true
				dfs(xx, yy, pos+1)
				vis[xx][yy] = false
			}
		}
	}
	for i := 0; i < m && !ans; i++ {
		for j := 0; j < n && !ans; j++ {
			if board[i][j] == word[0] {
				vis[i][j] = true
				dfs(i, j, 1)
				vis[i][j] = false
			}
		}
	}
	return ans
}

func wordBreak(s string, wordDict []string) bool {
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
