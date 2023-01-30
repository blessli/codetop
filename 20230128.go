package main

import "math"

func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	maxv := math.MinInt32
	right := -1
	minv := math.MaxInt32
	left := -1
	for i := 0; i < n; i++ {
		if nums[i] >= maxv {
			maxv = nums[i]
		} else {
			right = i
		}
		if nums[n-i-1] > minv {
			left = n - i - 1
		} else {
			minv = nums[n-i-1]
		}
	}
	if right == -1 {
		return 0
	}
	return right - left + 1
}

func findAnagrams(s string, p string) []int {
	ans := []int{}
	mp := map[byte]int{}
	ms := map[byte]int{}
	lenp := len(p)
	lens := len(s)
	if lens < lenp {
		return ans
	}
	for i := 0; i < lenp; i++ {
		mp[p[i]]++
	}
	check := func() bool {
		for k, v := range mp {
			if ms[k] != v {
				return false
			}
		}
		return true
	}
	for i := 0; i < lenp; i++ {
		ms[s[i]]++
	}
	if check() {
		ans = append(ans, 0)
	}
	left := 1
	for i := lenp; i < lens; i++ {
		ms[s[left-1]]--
		ms[s[i]]++
		if check() {
			ans = append(ans, left)
		}
		left++
	}
	return ans
}

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	flatten(root.Left)
	flatten(root.Right)
	left := root.Left
	right := root.Right
	if left != nil {
		for left.Right != nil {
			left = left.Right
		}
		left.Right = right
		root.Right = root.Left
		root.Left = nil
	}
}
// 二叉树中的最大路径和
func maxPathSum(root *TreeNode) int {
	ans := math.MinInt32
	var dfs func(*TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		if root.Left == nil && root.Right == nil {
			return root.Val
		}
		left := dfs(root.Left)
		right := dfs(root.Right)
		ans = max(ans, root.Val+left+right)
		return root.Val + max(left, right)
	}
	dfs(root)
	return ans
}

// 实现 Trie (前缀树)
type Trie struct {
	Next [26]*Trie
	End  bool
}

func Constructor() Trie {
	return Trie{}
}

func (t *Trie) Insert(word string) {
	root := t
	for i := 0; i < len(word); i++ {
		v := word[i] - 'a'
		if root.Next[v] == nil {
			root.Next[v] = &Trie{}
		}
		root = root.Next[v]
	}
	root.End = true
}

func (t *Trie) Search(word string) bool {
	ans := t.search(word)
	if ans == nil || !ans.End {
		return false
	}
	return true
}

func (t *Trie) search(word string) (ans *Trie) {
	root := t
	for i := 0; i < len(word); i++ {
		v := word[i] - 'a'
		if root.Next[v] == nil {
			return nil
		}
		root = root.Next[v]
	}
	return root
}

func (t *Trie) StartsWith(prefix string) bool {
	return t.search(prefix) != nil
}
