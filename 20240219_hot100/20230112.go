package main

type ListNode struct {
	Val  int
	Next *ListNode
}
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 1 https://leetcode.cn/problems/invert-binary-tree/description/
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	root.Left, root.Right = invertTree(root.Right), invertTree(root.Left)
	return root
}

// 1 https://leetcode.cn/problems/invert-binary-tree/description/
func invertTree_bfs(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		curr.Left, curr.Right = curr.Right, curr.Left
		if curr.Left != nil {
			queue = append(queue, curr.Left)
		}
		if curr.Right != nil {
			queue = append(queue, curr.Right)
		}
	}
	return root
}

// 1 https://leetcode.cn/problems/climbing-stairs/description/
func climbStairs(n int) int {
	if n <= 3 {
		return n
	}
	a, b := 1, 2
	for i := 3; i <= n; i++ {
		a, b = b, a+b
	}
	return b
}

// 2 https://leetcode.cn/problems/minimum-path-sum/description/
func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j > 0 {
				grid[i][j] += grid[i][j-1]
			} else if i > 0 && j == 0 {
				grid[i][j] += grid[i-1][j]
			} else if i > 0 && j > 0 {
				grid[i][j] += min(grid[i][j-1], grid[i-1][j])
			}
		}
	}
	return grid[m-1][n-1]
}

// 1 https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	leftDepth := maxDepth(root.Left)
	rightDepth := maxDepth(root.Right)
	return max(leftDepth, rightDepth) + 1
}

// 1 https://leetcode.cn/problems/maximum-depth-of-binary-tree/description/
func maxDepth_bfs(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	depth := 0
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[i]
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
		queue = queue[size:]
		depth++
	}
	return depth
}
// 2 https://leetcode.cn/problems/number-of-islands/description/
func numIslands(grid [][]byte) int {
	m := len(grid)
	n := len(grid[0])
	count := 0
	dfs := func(x, y int) {}
	dfs = func(x, y int) {
		if x < 0 || x >= m || y < 0 || y >= n || grid[x][y] == '0' {
			return
		}
		grid[x][y] = '0'
		dfs(x+1, y)
		dfs(x, y-1)
		dfs(x, y+1)
		dfs(x-1, y)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == '1' {
				count++
				dfs(i, j)
			}
		}
	}
	return count
}
