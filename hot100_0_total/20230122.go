package main

// 字符串解码 https://leetcode.cn/problems/decode-string/description/
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

// 从前序与中序遍历序列构造二叉树 https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/description/
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	rootVal := preorder[0]
	root := &TreeNode{Val: rootVal}
	var rootIndex int
	for i, val := range inorder {
		if val == rootVal {
			rootIndex = i
			break
		}
	}
	root.Left = buildTree(preorder[1:1+rootIndex], inorder[:rootIndex])
	root.Right = buildTree(preorder[1+rootIndex:], inorder[rootIndex+1:])
	return root
}

// 汉明距离 https://leetcode.cn/problems/hamming-distance/description/
func hammingDistance(x int, y int) int {
	xor := x ^ y
	count := 0
	for xor != 0 {
		count += xor & 1
		xor >>= 1
	}
	return count
}

// 找到所有数组中消失的数字 https://leetcode.cn/problems/find-all-numbers-disappeared-in-an-array/description/
func findDisappearedNumbers(nums []int) []int {
	result := []int{}
	for _, num := range nums {
		index := abs(num) - 1
		if nums[index] > 0 {
			nums[index] = -nums[index]
		}
	}
	for i, num := range nums {
		if num > 0 {
			result = append(result, i+1)
		}
	}
	return result
}
func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// 路径总和 II https://leetcode.cn/problems/path-sum-ii/description/
func pathSum(root *TreeNode, targetSum int) [][]int {
	var result [][]int
	var path []int
	var dfs func(*TreeNode, int)

	dfs = func(node *TreeNode, remaining int) {
		if node == nil {
			return
		}

		remaining -= node.Val
		path = append(path, node.Val)

		if node.Left == nil && node.Right == nil && remaining == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			result = append(result, tmp)
		}

		dfs(node.Left, remaining)
		dfs(node.Right, remaining)

		path = path[:len(path)-1]
	}

	dfs(root, targetSum)

	return result
}