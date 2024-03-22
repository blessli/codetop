package main

// 相交链表 https://leetcode.cn/problems/intersection-of-two-linked-lists/description/
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

// 回文链表 https://leetcode.cn/problems/palindrome-linked-list/description/
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// 使用快慢指针找到链表中间节点
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	// 翻转后半部分链表
	var prev *ListNode
	current := slow
	for current != nil {
		next := current.Next
		current.Next = prev
		prev = current
		current = next
	}

	// 比较前半部分和翻转后的后半部分
	p1, p2 := head, prev
	for p2 != nil {
		if p1.Val != p2.Val {
			return false
		}
		p1 = p1.Next
		p2 = p2.Next
	}

	return true
}

// 反转链表 https://leetcode.cn/problems/reverse-linked-list/description/
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	reversedHead := reverseList(head.Next)
	head.Next.Next = head // 在递归的回溯过程中，将当前节点的下一个节点的Next指针指向当前节点，实现了反转的过程
	head.Next = nil       // 将当前节点的Next指针断开，避免形成环
	return reversedHead
}

// 反转链表 https://leetcode.cn/problems/reverse-linked-list/description/
func reverseList2(head *ListNode) *ListNode {
	var prev *ListNode
	current := head
	for current != nil {
		next := current.Next
		current.Next = prev// 当前节点的下一个节点指向前一个节点
		prev = current // 更新前一个节点和当前节点
		current = next
	}
	return prev
}

// 单词搜索 https://leetcode.cn/problems/word-search/description/
func exist(board [][]byte, word string) bool {
	rows, cols := len(board), len(board[0])
	var dfs func(int, int, int) bool
	dfs = func(x, y, k int) bool {
		if k >= len(word) {
			return true
		}
		if x < 0 || x >= rows || y < 0 || y >= cols || board[x][y] != word[k] {
			return false
		}
		temp := board[x][y]
		board[x][y] += '#' // 打上标记，降低空间复杂度
		if dfs(x-1, y, k+1) || dfs(x+1, y, k+1) || dfs(x, y-1, k+1) || dfs(x, y+1, k+1) {
			return true
		}
		board[x][y] = temp
		return false
	}
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}

// 单词拆分 https://leetcode.cn/problems/word-break/description/
func wordBreak3(s string, wordDict []string) bool {
	n := len(s)
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}
	visited := make(map[int]bool)
	var dfs func(start int) bool
	dfs = func(start int) bool {
		if start == n {
			return true
		}
		if _, ok := visited[start]; ok {
			return false
		}
		for end := start + 1; end <= len(s); end++ {
			if wordSet[s[start:end]] && dfs(end) {
				return true
			}
		}
		visited[start] = true
		return false
	}
	return dfs(0)
}

// 单词拆分 https://leetcode.cn/problems/word-break/description/
func wordBreak2(s string, wordDict []string) bool {
	n := len(s)
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}
	queue := []int{0}
	visited := make(map[int]bool)
	for len(queue) > 0 {
		start := queue[0]
		queue = queue[1:]
		if _, ok := visited[start]; ok {
			continue
		}
		for end := start + 1; end <= n; end++ {
			if wordSet[s[start:end]] {
				if end == n {
					return true
				}
				queue = append(queue, end)
			}
		}
		visited[start] = true
	}
	return false
}

// 单词拆分 https://leetcode.cn/problems/word-break/description/
func wordBreak(s string, wordDict []string) bool {
	wordSet := make(map[string]bool)
	for _, word := range wordDict {
		wordSet[word] = true
	}
	n := len(s)
	dp := make([]bool, n+1) // 前i个字符可被拆分
	dp[0] = true
	for i := 1; i <= n; i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[n]
}
