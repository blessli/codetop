package main

func rob2(root *TreeNode) int {
	var dfs func(*TreeNode) [2]int
	ans := 0
	dfs = func(root *TreeNode) [2]int {
		if root == nil {
			return [2]int{0, 0}
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		res := [2]int{}
		res[0] = root.Val + left[1] + right[1]
		res[1] = max(left[0], left[1]) + max(right[0], right[1])
		ans = max(ans, max(res[0], res[1]))
		return res
	}
	dfs(root)
	return ans
}
