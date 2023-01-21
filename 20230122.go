package main


func decodeString(s string) string {
	ans := ""
	n := len(s)
	find := func(pos int) int {
		c := 0
		for i := pos; i < n; i++ {
			if s[i] == '[' {
				c++
			} else if s[i] == ']' {
				c--
				if c == 0 {
					return i
				}
			}
		}
		return 0
	}
	var dfs func(int, int, int) string
	dfs = func(m, start, end int) string {
		if start > end {
			return ""
		}
		if s[start] >= 'a' && s[start] <= 'z' {
			return string(s[start]) + dfs(0, start+1, end)
		}
		if s[start] >= '0' && s[start] <= '9' {
			if s[start-1] < '0' || s[start-1] > '9' {
				m = 0
			}
			m = m*10 + (int(s[start]) - '0')
			return dfs(m, start+1, end)
		}
		if s[start] == '[' {
			curr := start
			res := ""
			for curr <= end {
				if s[curr] >= 'a' && s[curr] <= 'z' {
					res += string(s[curr])
					curr++
					continue
				}
				if s[curr] >= '0' && s[curr] <= '9' {
					m = m*10 + (int(s[curr]) - '0')
					curr++
					continue
				}
				if s[curr] == '[' {
					subs := dfs(m, curr+1, find(curr)-1)
					for j := 0; j < m; j++ {
						res += subs
					}
					curr = find(curr) + 1
					m = 0
				}
			}
			return res
		}
		return dfs(m, start+1, end)
	}

	m := 0
	i := 0
	for {
		if i >= n {
			break
		}
		if s[i] >= 'a' && s[i] <= 'z' {
			ans += string(s[i])
			i++
			continue
		}
		if s[i] >= '0' && s[i] <= '9' {
			m = m*10 + (int(s[i]) - '0')
			i++
			continue
		}
		if s[i] == '[' {
			subs := dfs(m, i+1, find(i)-1)
			for j := 0; j < m; j++ {
				ans += subs
			}
			i = find(i) + 1
			m = 0
		}
	}
	return ans
}
func buildTree(preorder []int, inorder []int) *TreeNode {
	var dfs func(int, int, int, int) *TreeNode
	dfs = func(pre1, pre2, in1, in2 int) *TreeNode {
		if pre1 > pre2 || in1 > in2 {
			return nil
		}
		root := &TreeNode{Val: preorder[pre1]}
		pos := in1
		for i := 0; i < len(inorder); i++ {
			if inorder[i] == preorder[pre1] {
				pos = i
				break
			}
		}
		root.Left = dfs(pre1+1, pre1+pos-in1, in1, pos-1)
		root.Right = dfs(pre1+pos-in1+1, pre2, pos+1, in2)
		return root
	}
	return dfs(0, len(preorder)-1, 0, len(inorder)-1)
}

func hammingDistance(x int, y int) int {
	ans := 0
	xx, yy := 0, 0
	for x != 0 || y != 0 {
		xx = x % 2
		x /= 2
		yy = y % 2
		y /= 2
		if xx != yy {
			ans++
		}
	}
	return ans
}

func findDisappearedNumbers(nums []int) []int {
	for i := 0; i < len(nums); i++ {
		pos := i
		temp := nums[pos]
		if temp == nums[temp-1] {
			continue
		}
		for {
			nums[pos] = nums[temp-1]
			nums[temp-1] = temp
			temp = nums[pos]
			if temp == nums[temp-1] {
				break
			}
		}
	}
	ans := []int{}
	for i := 0; i < len(nums); i++ {
		if nums[i] != i+1 {
			ans = append(ans, i+1)
		}
	}
	return ans
}
func pathSum(root *TreeNode, targetSum int) int {
	var dfs func(*TreeNode) int
	var dfsPath func(*TreeNode, int) int
	dfsPath = func(tn *TreeNode, sum int) int {
		if tn == nil {
			return 0
		}
		res := 0
		if sum == tn.Val {
			res++
		}
		res += dfsPath(tn.Left, sum-tn.Val)
		res += dfsPath(tn.Right, sum-tn.Val)
		return res
	}
	dfs = func(tn *TreeNode) int {
		if tn == nil {
			return 0
		}
		ret := dfsPath(tn, targetSum)
		ret += dfs(tn.Left)
		ret += dfs(tn.Right)
		return ret
	}
	return dfs(root)
}
