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
// 翻转二叉树
func invertTree_(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	temp := root.Left
	root.Left = root.Right
	root.Right = temp
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}
// bfs 0ms
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	invert := func(root *TreeNode) {
		temp := root.Left
		root.Left = root.Right
		root.Right = temp
	}
	queue := []*TreeNode{}
	ans := root
	queue = append(queue, root)
	for len(queue) > 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			curr := queue[i]
			invert(curr)
			if curr.Left != nil {
				queue = append(queue, curr.Left)
			}
			if curr.Right != nil {
				queue = append(queue, curr.Right)
			}
		}
		queue = queue[size:]
	}
	return ans
}

func climbStairs(n int) int {
	if n <= 3 {
		return n
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dp[2] = 2
	dp[3] = 3
	for i := 3; i <= n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n]
}

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				dp[i][j] = grid[i][j]
			} else if i == 0 {
				dp[i][j] = dp[i][j-1] + grid[i][j]
			} else if j == 0 {
				dp[i][j] = dp[i-1][j] + grid[i][j]
			} else {
				dp[i][j] = min(dp[i][j-1], dp[i-1][j]) + grid[i][j]
			}
		}
	}
	return dp[m-1][n-1]
}

// 二叉树的最大深度
func maxDepth_(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth_(root.Left), maxDepth_(root.Right)) + 1
}

// bfs 0ms
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	ans := 0
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
		ans++
	}
	return ans
}

func numIslands(grid [][]byte) int {
	ans := 0
	m := len(grid)
	n := len(grid[0])
	dfs := func(x, y int) {}
	dfs = func(x, y int) {
		if x < 0 || x >= m || y < 0 || y >= n {
			return
		}
		if grid[x][y] == '0' {
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
			key := grid[i][j]
			if key == '1' {
				ans++
				dfs(i, j)
			}
		}
	}
	return ans
}
