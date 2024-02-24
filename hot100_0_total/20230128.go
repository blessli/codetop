package main

import "math"

// 最短无序连续子数组 https://leetcode.cn/problems/shortest-unsorted-continuous-subarray/description/
func findUnsortedSubarray(nums []int) int {
	n := len(nums)
	if n <= 1 {
		return 0
	}

	// 找到需要排序的子数组的左边界和右边界
	left, right := -1, -1
	maxVal, minVal := nums[0], nums[n-1]

	for i := 0; i < n; i++ {
		maxVal = max(maxVal, nums[i])
		if nums[i] < maxVal {
			right = i
		}

		minVal = min(minVal, nums[n-1-i])
		if nums[n-1-i] > minVal {
			left = n - 1 - i
		}
	}

	if left == -1 {
		return 0
	}

	return right - left + 1
}

// 找到字符串中所有字母异位词 https://leetcode.cn/problems/find-all-anagrams-in-a-string/description/
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

// 二叉树展开为链表 https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/description/
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

// 二叉树中的最大路径和 https://leetcode.cn/problems/binary-tree-maximum-path-sum/description/
func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32

	var maxPathSumHelper func(node *TreeNode) int
	maxPathSumHelper = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		left := max(0, maxPathSumHelper(node.Left))
		right := max(0, maxPathSumHelper(node.Right))

		maxSum = max(maxSum, node.Val+left+right)

		return node.Val + max(left, right)
	}

	maxPathSumHelper(root)

	return maxSum
}

// 实现 Trie (前缀树)
type TrieNode struct {
	children map[rune]*TrieNode
	isEnd    bool
}

type Trie struct {
	root *TrieNode
}

func Constructor() Trie {
	return Trie{
		root: &TrieNode{
			children: map[rune]*TrieNode{},
			isEnd:    false,
		},
	}
}

func (t *Trie) Insert(word string) {
	node := t.root
	for _, char := range word {
		if _, ok := node.children[char]; !ok {
			node.children[char] = &TrieNode{
				children: make(map[rune]*TrieNode),
				isEnd:    false,
			}
		}
		node = node.children[char]
	}
	node.isEnd = true
}

func (t *Trie) Search(word string) bool {
	node := t.root
	for _, char := range word {
		if _, ok := node.children[char]; !ok {
			return false
		}
		node = node.children[char]
	}
	return node.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	node := t.root
	for _, char := range prefix {
		if _, ok := node.children[char]; !ok {
			return false
		}
		node = node.children[char]
	}
	return true
}
