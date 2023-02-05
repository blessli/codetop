package main

import (
	"math"
)
// 二叉树中的最大路径和
func maxPathSum(root *TreeNode) int {
	ans := math.MinInt64
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left := max(0, dfs(root.Left))
		right := max(0, dfs(root.Right))
		ans = max(ans, root.Val+left+right)
		return root.Val + max(left, right)
	}
	dfs(root)
	return ans
}

// 从前序与中序遍历序列构造二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	var dfs func([]int, []int, int, int, int, int) *TreeNode
	dfs = func(preorder, inorder []int, pre1, pre2, in1, in2 int) *TreeNode {
		if pre1 > pre2 || in1 > in2 {
			return nil
		}
		root := &TreeNode{Val: preorder[pre1]}
		pos := in1
		for i := in1; i <= in2; i++ {
			if inorder[i] == preorder[pre1] {
				pos = i
				break
			}
		}
		root.Left = dfs(preorder, inorder, pre1+1, pre1+pos-in1, in1, pos-1)
		root.Right = dfs(preorder, inorder, pre1+pos-in1+1, pre2, pos+1, in2)
		return root
	}
	return dfs(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}

// 二叉树的中序遍历
func inorderTraversal(root *TreeNode) []int {
	ans := []int{}
	var dfs func(*TreeNode)
	var bfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root != nil {
			dfs(root.Left)
			ans = append(ans, root.Val)
			dfs(root.Right)
		}
	}
	bfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		stack := []*TreeNode{}
		stack = append(stack, root)
		for len(stack) > 0 || root != nil {
			for root != nil {
				if root.Left != nil {
					stack = append(stack, root.Left)
				}
				root = root.Left
			}
			root = stack[len(stack)-1]
			ans = append(ans, root.Val)
			stack = stack[:len(stack)-1]
			if root.Right != nil {
				stack = append(stack, root.Right)
			}
			root = root.Right
		}
	}

	dfs(nil)
	bfs(root)
	return ans
}

// 二叉树的最近公共祖先
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	var dfs func(*TreeNode, *TreeNode, *TreeNode)
	var bfs func(*TreeNode, *TreeNode, *TreeNode)
	var search func(*TreeNode, int) bool
	search = func(root *TreeNode, target int) bool {
		if root != nil {
			if root.Val == target {
				return true
			}
			ans := search(root.Left, target) || search(root.Right, target)
			if ans {
				return true
			}
		}
		return false
	}
	ans := root
	// 1196ms
	dfs = func(root, p, q *TreeNode) {
		if root == nil {
			return
		}
		left, right := root, root
		if search(left, p.Val) && search(right, q.Val) {
			ans = root
		}
		dfs(root.Left, p, q)
		dfs(root.Right, p, q)
	}
	// 16ms
	bfs = func(root, p, q *TreeNode) {
		if root == nil {
			return
		}
		m := map[int]*TreeNode{}
		queue := []*TreeNode{root}
		for len(queue) > 0 {
			size := len(queue)
			for i := 0; i < size; i++ {
				if queue[i].Left != nil {
					queue = append(queue, queue[i].Left)
					m[queue[i].Left.Val] = queue[i]
				}
				if queue[i].Right != nil {
					queue = append(queue, queue[i].Right)
					m[queue[i].Right.Val] = queue[i]
				}
			}
			queue = queue[size:]
		}
		vis := map[int]*TreeNode{}
		temp := p
		vis[temp.Val] = p
		for {
			if v, ok := m[temp.Val]; ok {
				vis[v.Val] = v
				temp = v
			} else {
				break
			}
		}
		temp = q
		for {
			if vis[temp.Val] != nil {
				ans = vis[temp.Val]
				break
			}
			if v, ok := m[temp.Val]; ok {
				temp = v
			} else {
				break
			}
		}
	}
	dfs(nil, p, q)
	bfs(root, p, q)
	return ans
}

// 验证二叉搜索树
func isValidBST(root *TreeNode) bool {
	var dfs func(*TreeNode, int, int) bool
	var bfs func(root *TreeNode) bool
	dfs = func(root *TreeNode, low, high int) bool {
		if root == nil {
			return true
		}
		if root.Val <= low || root.Val >= high {
			return false
		}
		return dfs(root.Left, low, root.Val) && dfs(root.Right, root.Val, high)
	}
	bfs = func(root *TreeNode) bool {
		if root == nil {
			return false
		}
		stack := []*TreeNode{}
		stack = append(stack, root)
		minv := math.MinInt64
		for len(stack) > 0 || root != nil {
			for root != nil {
				if root.Left != nil {
					stack = append(stack, root.Left)
					root = root.Left
				}
			}
			root = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if root.Val <= minv {
				return false
			}
			minv = root.Val
			if root.Right != nil {
				stack = append(stack, root.Right)
			}
			root = root.Right
		}
		return true
	}
	// bfs(nil)
	dfs(nil, math.MinInt64, math.MaxInt64)
	return bfs(root)
}

// 对称二叉树
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	var dfs func(*TreeNode, *TreeNode) bool
	var bfs func(root *TreeNode) bool
	// dfs解法
	dfs = func(left, right *TreeNode) bool {
		if left == nil && right == nil {
			return true
		}
		if left == nil || right == nil || left.Val != right.Val {
			return false
		}
		return dfs(left.Left, right.Right) && dfs(left.Right, right.Left)
	}
	// bfs解法
	bfs = func(root *TreeNode) bool {
		ans := 0
		queue := []*TreeNode{root}
		for len(queue) > 0 {
			size := len(queue)
			ans++
			if ans > 1 && size%2 == 1 {
				return false
			}
			if ans > 1 {
				for i := 0; i < size/2; i++ {
					left, right := queue[i], queue[size-i-1]
					if left.Val != right.Val {
						return false
					}
					if (left.Left == nil && right.Right != nil) || (left.Left != nil && right.Right == nil) {
						return false
					}
					if (left.Right == nil && right.Left != nil) || (left.Right != nil && right.Left == nil) {
						return false
					}
				}
			}
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
		return true
	}
	return bfs(root)
	// return dfs(root.Left, root.Right)
}

// 二叉树的层序遍历
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	ans := [][]int{}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		size := len(queue)
		curr := []int{}
		for i := 0; i < size; i++ {
			curr = append(curr, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[size:]
		ans = append(ans, curr)
	}
	return ans
}

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
