package main

import (
	"sort"
	"strings"
)

// 合并 K 个升序链表 https://leetcode.cn/problems/merge-k-sorted-lists/description/
func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	var merge func(lists []*ListNode, left, right int) *ListNode
	merge = func(lists []*ListNode, left, right int) *ListNode {
		if left == right {
			return lists[left]
		}
		if left < right {
			mid := left + (right-left)/2
			l1 := merge(lists, left, mid)
			l2 := merge(lists, mid+1, right)
			return mergeTwoLists(l1, l2)
		}
		return nil
	}
	return merge(lists, 0, len(lists)-1)
}

// 不同的二叉搜索树 https://leetcode.cn/problems/unique-binary-search-trees/description/
func numTrees(n int) int {
	if n <= 0 {
        return 0
    }
    
    dp := make([]int, n+1)
    dp[0], dp[1] = 1, 1
    
    for i := 2; i <= n; i++ {
        for j := 1; j <= i; j++ {
            dp[i] += dp[j-1] * dp[i-j]
        }
    }
    
    return dp[n]
}

// 颜色分类 https://leetcode.cn/problems/sort-colors/description/
func sortColors(nums []int) {
	count := make([]int, 3)
	for _, num := range nums {
		count[num]++
	}
	index := 0
	for color := 0; color < 3; color++ {
		for count[color] > 0 {
			nums[index] = color
			index++
			count[color]--
		}
	}
}

// 字母异位词分组 https://leetcode.cn/problems/group-anagrams/description/
func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string)
	
	for _, str := range strs {
		s := strings.Split(str, "")
		sort.Strings(s)
		sortedStr := strings.Join(s, "")
		
		anagrams[sortedStr] = append(anagrams[sortedStr], str)
	}
	
	result := make([][]string, 0, len(anagrams))
	for _, v := range anagrams {
		result = append(result, v)
	}
	
	return result
}


// 分割等和子集 https://leetcode.cn/problems/partition-equal-subset-sum/description/
func canPartition(nums []int) bool {
    total := 0
    for _, num := range nums {
        total += num
    }
    
    if total%2 != 0 {
        return false
    }
    
    target := total / 2
    dp := make([]bool, target+1)
    dp[0] = true
    
    for _, num := range nums {
        for i := target; i >= num; i-- {
            dp[i] = dp[i] || dp[i-num]
        }
    }
    
    return dp[target]
}
