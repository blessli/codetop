package main

// 1 https://leetcode.cn/problems/linked-list-cycle/description/
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next
	for fast != nil && fast.Next != nil {
		// 如果链表中存在环，那么快指针 fast 最终会追上慢指针 slow
		if slow == fast {
			return true
		}
		slow = slow.Next
		fast = fast.Next.Next
	}
	return false
}

// 1 https://leetcode.cn/problems/single-number/description/
func singleNumber(nums []int) int {
	ans := 0
	for _, v := range nums {
		ans ^= v
	}
	return ans
}

// 1 https://leetcode.cn/problems/binary-search/description/
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}
// 2 https://leetcode.cn/problems/binary-tree-level-order-traversal/description/
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var ans [][]int
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		levelSize := len(queue)
		levelVals := make([]int, levelSize)
		for i := 0; i < levelSize; i++ {
			node := queue[i]
			levelVals[i] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		ans = append(ans, levelVals)
		queue = queue[levelSize:]
	}
	return ans
}

// 2 https://leetcode.cn/problems/binary-tree-right-side-view/description/
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := []int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[i]
			if i == size-1 {
				ans = append(ans, node.Val)
			}
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		queue = queue[size:]
	}
	return ans
}
