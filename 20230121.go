package main

import (
	"strconv"
)

func letterCombinations(digits string) []string {
	grid := [][]string{{"a", "b", "c"}, {"d", "e", "f"}, {"g", "h", "i"}, {"j", "k", "l"}, {"m", "n", "o"}, {"p", "q", "r", "s"}, {"t", "u", "v"}, {"w", "x", "y", "z"}}
	n := len(digits)
	if n==0{
		return []string{}
	}
	nums := make([]int, n)
	ans := []string{}
	for i := 0; i < n; i++ {
		nums[i] = int(digits[i]) - '0' - 2
	}
	var dfs func(string, int)
	dfs = func(curr string, pos int) {
		if pos == n {
			ans = append(ans, curr)
			return
		}
		for i := 0; i < len(grid[nums[pos]]); i++ {
			dfs(curr+grid[nums[pos]][i], pos+1)
		}
	}
	dfs("", 0)
	return ans
}
func subsets(nums []int) [][]int {
	n := len(nums)
	ans := make([][]int, 0)
	vis := map[string]bool{}
	m := map[int]bool{}
	var dfs func([]int, int)
	dfs = func(curr []int, pos int) {
		val := ""
		for _, v := range curr {
			val += strconv.Itoa(v)
		}
		if vis[val] {
			return
		}
		vis[val] = true
		temp := make([]int, 0)
		temp = append(temp, curr...)
		ans = append(ans, temp)
		for i := pos; i < n; i++ {
			if len(curr) < n && !m[i] {
				v := nums[i]
				curr = append(curr, v)
				m[i] = true
				dfs(curr, i)
				m[i] = false
				curr = curr[:len(curr)-1]
			}
		}
	}
	dfs([]int{}, 0)
	return ans
}
func detectCycle(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return nil
	}
	temp := head
	pre := head
	last := head
	vis := false
	for last.Next != nil && last.Next.Next != nil {
		pre = pre.Next
		last = last.Next.Next
		if pre == last {
			vis = true
			break
		}
	}
	if !vis {
		return nil
	}
	for pre != temp {
		pre = pre.Next
		temp = temp.Next
	}
	return pre
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	curr := &ListNode{}
	ans := curr
	mod := 0
	for l1 != nil || l2 != nil {
		if l1 == nil {
			temp := &ListNode{Val: (l2.Val + mod) % 10}
			mod = (l2.Val + mod) / 10
			curr.Next = temp
			curr = curr.Next
			l2 = l2.Next
		} else if l2 == nil {
			temp := &ListNode{Val: (l1.Val + mod) % 10}
			mod = (l1.Val + mod) / 10
			curr.Next = temp
			curr = curr.Next
			l1 = l1.Next
		} else {
			temp := &ListNode{Val: (l1.Val + l2.Val + mod) % 10}
			mod = (l1.Val + l2.Val + mod) / 10
			curr.Next = temp
			curr = curr.Next
			l2 = l2.Next
			l1 = l1.Next
		}
	}
	if mod > 0 {
		temp := &ListNode{Val: mod}
		curr.Next = temp
		curr = curr.Next
	}
	return ans.Next
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
