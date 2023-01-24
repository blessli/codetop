package main

import (
	"sort"
)

func mergeKLists(lists []*ListNode) *ListNode {
	n := len(lists)
	if n == 0 {
		return nil
	}
	if n == 1 {
		return lists[0]
	}
	if n == 2 {
		return mergeTwoLists(lists[0], lists[1])
	}
	var merge func([]*ListNode, int, int) *ListNode
	merge = func(lists []*ListNode, low, high int) *ListNode {
		if low == high {
			return lists[low]
		}
		if low+1 == high {
			return mergeTwoLists(lists[low], lists[high])
		}
		mid := low + (high-low)/2
		return mergeTwoLists(merge(lists, low, mid), merge(lists, mid, high))
	}
	return merge(lists, 0, len(lists)-1)
}

func numTrees(n int) int {
	dp := make([]int, n+1)
	dp[0] = 1
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 0; j < i; j++ {
			dp[i] += dp[j] * dp[i-j-1]
		}
	}
	return dp[n]
}

func sortColors(nums []int) {
	dfs := func(nums []int, target int) int {
		cnt := 0
		for i := 0; i < len(nums); i++ {
			if nums[i] == target {
				nums[cnt], nums[i] = nums[i], nums[cnt]
				cnt++
			}
		}
		return cnt
	}
	cnt0 := dfs(nums, 0)
	dfs(nums[cnt0:], 1)
}

func groupAnagrams(strs []string) [][]string {
	m := map[string][]int{}
	dfs := func(s string) string {
		ss := []rune(s)
		sort.Slice(ss, func(i, j int) bool {
			return ss[i] <= ss[j]
		})
		return string(ss)
	}
	for i, v := range strs {
		nv := dfs(v)
		if _, ok := m[nv]; !ok {
			m[nv] = make([]int, 0)
		}
		m[nv] = append(m[nv], i)
	}
	ans := make([][]string, len(m))
	index := 0
	for _, v := range m {
		ans[index] = make([]string, 0)
		for i := 0; i < len(v); i++ {
			ans[index] = append(ans[index], strs[v[i]])
		}
		index++
	}
	return ans
}
// 类似：0-1背包问题。空间复杂度待优化
func canPartition(nums []int) bool {
	sum := 0
	for _, v := range nums {
		sum += v
	}
	if sum%2 == 1 {
		return false
	}
	n := len(nums)
	sum = sum / 2
	dp := make([][]bool, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]bool, sum+1)
	}
	for i := 0; i < n; i++ {
		for w := 1; w <= sum; w++ {
			if i == 0 {
				dp[i][w] = (w == nums[i])
				continue
			}
			if w < nums[i] {
				dp[i][w] = dp[i-1][w]
			} else {
				dp[i][w] = dp[i-1][w] || dp[i-1][w-nums[i]]
			}
		}
	}
	return dp[n-1][sum]
}
