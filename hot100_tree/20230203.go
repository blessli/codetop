package main

// 翻转二叉树
func invertTree(root *TreeNode) *TreeNode {
	var dfs func(*TreeNode)
	// dfs解法
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		root.Left, root.Right = root.Right, root.Left
		dfs(root.Left)
		dfs(root.Right)
	}
	// bfs解法
	bfs := func(root *TreeNode) {
		if root == nil {
			return
		}
		ans := 0
		queue := []*TreeNode{root}
		for len(queue) > 0 {
			size := len(queue)
			ans++
			for i := 0; i < size; i++ {
				queue[i].Left, queue[i].Right = queue[i].Right, queue[i].Left
				if queue[i].Left != nil {
					queue = append(queue, queue[i].Left)
				}
				if queue[i].Right != nil {
					queue = append(queue, queue[i].Right)
				}
			}
			queue = queue[size:]
		}
	}
	temp := root
	dfs(temp)
	bfs(temp)
	return root
}

// 二叉树的最大深度
func maxDepth(root *TreeNode) int {
	var dfs func(*TreeNode) int
	// dfs解法
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		return max(dfs(root.Left), dfs(root.Right)) + 1
	}
	// bfs解法
	bfs := func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		ans := 0
		queue := []*TreeNode{root}
		for len(queue) > 0 {
			size := len(queue)
			ans++
			for i := 0; i < size; i++ {
				if queue[i].Left != nil {
					queue = append(queue, queue[i].Left)
				}
				if queue[i].Right != nil {
					queue = append(queue, queue[i].Right)
				}
			}
			queue = queue[size:]
		}
		return ans
	}
	// return dfs(root)
	return bfs(root)
}

// 二叉树展开为链表
// 写的有点长，貌似不是最优解
// todo：如何优雅
func flatten(root *TreeNode) {
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		dfs(root.Left)
		temp1 := root.Left
		temp2 := root.Left
		for temp2 != nil && temp2.Right != nil {
			temp2 = temp2.Right
		}
		root.Left = nil
		temp3 := root.Right
		if temp1 != nil {
			root.Right = temp1
			if temp2 != nil {
				temp2.Right = temp3
			}
		}
		dfs(temp3)
	}
	temp := root
	dfs(temp)
}

// 二叉树的直径
func diameterOfBinaryTree(root *TreeNode) int {
	ans := 0
	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		ans = max(ans, left+right+1)
		return max(left, right) + 1
	}
	dfs(root)
	return ans - 1
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
