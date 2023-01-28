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

type Trie struct {
	Next [26]*Trie
    End bool
}

func Constructor() Trie {
	return Trie{}
}
// 循环better
func (t *Trie) Insert(word string) {
	n := len(word)
	var dfs func(*Trie, int)
	dfs = func(root *Trie, pos int) {
		if pos >= n {
            root.End=true
			return
		}
		if root.Next[word[pos]-'a'] == nil {
			root.Next[word[pos]-'a'] = &Trie{}
		}
		dfs(root.Next[word[pos]-'a'], pos+1)
	}
	node := t
	dfs(node, 0)
}

func (t *Trie) Search(word string) bool {
	ans := t.search(word)
	if ans == nil || !ans.End{
		return false
	}
	return true
}

func (t *Trie) search(word string) (ans *Trie) {
	n := len(word)
	vis := false
	var dfs func(*Trie, int)
	dfs = func(root *Trie, pos int) {
		if vis {
			return
		}
		if root==nil {
			return
		}
		if pos == n {
			ans = root
			vis = true
			return
		}
		dfs(root.Next[word[pos]-'a'], pos+1)
	}
	node := t
	dfs(node, 0)
	if !vis {
		ans = nil
	}
	return
}

func (t *Trie) StartsWith(prefix string) bool {
	if t.search(prefix) == nil {
		return false
	}
	return true
}