package main

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	pre := head
	last := head.Next
	for last != nil && last.Next != nil {
		if pre == last {
			return true
		}
		pre = pre.Next
		last = last.Next.Next
	}
	return false
}

func singleNumber(nums []int) int {
	ans := 0
	for _, v := range nums {
		ans ^= v
	}
	return ans
}

func search(nums []int, target int) int {
	left := 0
	right := len(nums) - 1
	for left <= right {
		mid := left + (right-left)/2
		v := nums[mid]
		if v == target {
			return mid
		} else if v < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := make([][]int, 0)
	data := []*TreeNode{}
	data = append(data, root)
	for len(data) > 0 {
		curr := []*TreeNode{}
		layer := []int{}
		for _, node := range data {
			layer = append(layer, node.Val)
			if node.Left != nil {
				curr = append(curr, node.Left)
			}
			if node.Right != nil {
				curr = append(curr, node.Right)
			}
		}
		data = curr
		ans = append(ans, layer)
	}
	return ans
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	ans := make([]int, 0)
	data := []*TreeNode{}
	data = append(data, root)
	for len(data) > 0 {
		curr := []*TreeNode{}
		layer := []int{}
		ans = append(ans, data[len(data)-1].Val)
		for _, node := range data {
			layer = append(layer, node.Val)
			if node.Left != nil {
				curr = append(curr, node.Left)
			}
			if node.Right != nil {
				curr = append(curr, node.Right)
			}
		}
		data = curr

	}
	return ans
}
