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
// 两数相加 https://leetcode.cn/problems/add-two-numbers/description
func addTwoNumbers_(l1 *ListNode, l2 *ListNode, carry int) *ListNode {
	// 递归写法
	if l1 == nil && l2 == nil && carry == 0 {
		return nil
	}
	sum := carry
	if l1 != nil {
		sum += l1.Val
		l1 = l1.Next
	}
	if l2 != nil {
		sum += l2.Val
		l2 = l2.Next
	}
	newNode := &ListNode{Val: sum % 10}
	newNode.Next = addTwoNumbers_(l1, l2, sum/10)
	return newNode
}

type LRUCache struct {
	capacity int
	cache    map[int]*Node
	head     *Node
	tail     *Node
}

type Node struct {
	key   int
	value int
	prev  *Node
	next  *Node
}

func Constructor1(capacity int) LRUCache {
	head := &Node{}
	tail := &Node{}
	head.next = tail
	tail.prev = head
	return LRUCache{
		capacity: capacity,
		cache:    make(map[int]*Node),
		head:     head,
		tail:     tail,
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.cache[key]; ok {
		this.moveToHead(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.cache[key]; ok {
		node.value = value
		this.moveToHead(node)
	} else {
		newNode := &Node{key: key, value: value}
		this.cache[key] = newNode
		this.addToHead(newNode)

		if len(this.cache) > this.capacity {
			tailKey := this.tail.prev.key
			this.removeNode(this.tail.prev)
			delete(this.cache, tailKey)
		}
	}
}

func (this *LRUCache) moveToHead(node *Node) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeNode(node *Node) {
	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}
}

func (this *LRUCache) addToHead(node *Node) {
	if this.head.next != nil {
		this.head.next.prev = node
	}
	node.next = this.head.next
	node.prev = this.head
	this.head.next = node
}
