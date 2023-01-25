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
