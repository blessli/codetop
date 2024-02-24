package main

// 电话号码的字母组合 https://leetcode.cn/problems/letter-combinations-of-a-phone-number/description/
func letterCombinations(digits string) []string {
	var phoneMap = map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	n := len(digits)
	ans := []string{}
	if n == 0 {
		return ans
	}
	var dfs func(int, string)
	dfs = func(index int, path string) {
		if index == n {
			ans = append(ans, path)
			return
		}
		letters := phoneMap[digits[index]]
		for i := 0; i < len(letters); i++ {
			dfs(index+1, path+string(letters[i]))
		}
	}
	dfs(0, "")
	return ans
}

// 子集 https://leetcode.cn/problems/subsets/description/
func subsets(nums []int) [][]int {
	n := len(nums)
	ans := [][]int{}
	var dfs func(start int, path []int)
	dfs = func(start int, path []int) {
		temp := make([]int, len(path))
		copy(temp, path)
		ans = append(ans, temp)
		for i := start; i < n; i++ {
			path = append(path, nums[i])
			dfs(i+1, path)
			path = path[:len(path)-1]
		}
	}
	dfs(0, []int{})
	return ans
}

// 环形链表 II https://leetcode.cn/problems/linked-list-cycle-ii/description/
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	slow, fast := head, head
	hasCycle := false
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			hasCycle = true
			break
		}
	}
	if !hasCycle {
		return nil
	}
	slow = head
	for slow != fast {
		slow = slow.Next
		fast = fast.Next
	}
	return slow
}
// 两数相加 https://leetcode.cn/problems/add-two-numbers/description/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	carry := 0
	for l1 != nil || l2 != nil {
		sum := carry
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
		carry = sum / 10
	}
	if carry > 0 {
		current.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}

type LRUCache struct {
	cache    map[int]int
	capacity int
	count    int
	list     *Node
}

// 偷懒没用双向链表，险些超时
type Node struct {
	Val  int
	Next *Node
}

func Constructor1(capacity int) LRUCache {
	return LRUCache{cache: make(map[int]int), capacity: capacity, list: nil}
}

func (c *LRUCache) Get(key int) int {
	v, ok := c.cache[key]
	if !ok {
		return -1
	}
	if c.list == nil {
		panic("list nil")
	}
	if c.count > 1 && c.list.Val != key {
		c.refresh(key)
	}
	return v
}

func (c *LRUCache) refresh(key int) {
	pre := c.list
	curr := c.list.Next
	for curr != nil && curr.Val != key {
		pre = curr
		curr = curr.Next
	}
	pre.Next = pre.Next.Next
	c.list = &Node{Val: key, Next: c.list}
}

func (c *LRUCache) Put(key int, value int) {
	unHit := c.Get(key) == -1
	c.cache[key] = value
	if !unHit {
		return
	}
	if c.count < c.capacity {
		if c.count == 0 {
			c.list = &Node{Val: key}
		} else {
			c.list = &Node{Val: key, Next: c.list}
		}
		c.count++
	} else {
		if c.capacity == 1 {
			delete(c.cache, c.list.Val)
			c.list = &Node{Val: key}
			return
		}
		pre := c.list
		curr := c.list
		for curr.Next != nil {
			pre = curr
			curr = curr.Next
		}
		delete(c.cache, curr.Val)
		pre.Next = nil
		c.list = &Node{Val: key, Next: c.list}
	}
}
