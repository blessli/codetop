package main

import "sort"

// 两数之和 https://leetcode.cn/problems/two-sum/description/
func twoSum(nums []int, target int) []int {
	m := map[int]int{}
	for i := 0; i < len(nums); i++ {
		key := target - nums[i]
		if v, ok := m[key]; ok {
			return []int{v, i}
		}
		m[nums[i]] = i
	}
	return nil
}

// 有效的括号 https://leetcode.cn/problems/valid-parentheses/description/
func isValid(s string) bool {
	mapping := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := []byte{}
	for i := 0; i < len(s); i++ {
		key := s[i]
		if key == '(' || key == '{' || key == '[' {
			stack = append(stack, key)
		} else {
			if len(stack) == 0 || mapping[key] != stack[len(stack)-1] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}
	return len(stack) == 0
}

// 三数之和 https://leetcode.cn/problems/3sum/description/
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j, k := i+1, len(nums)-1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				for j < k && nums[j] == nums[j+1] {
					j++
				}
				for j < k && nums[k] == nums[k-1] {
					k--
				}
				j++
				k--
			} else if sum > 0 {
				k--
			} else {
				j++
			}
		}

	}
	return ans
}

// 最大子数组和 https://leetcode.cn/problems/maximum-subarray/description/
func maxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	curNum := nums[0]
	maxNum := nums[0]
	for i := 1; i < len(nums); i++ {
		curNum = max(nums[i], curNum+nums[i])
		maxNum = max(maxNum, curNum)
	}
	return maxNum
}

// 无重复字符的最长子串 https://leetcode.cn/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	maxLen := 0
	start := 0
	for i := 0; i < len(s); i++ {
		// 表示当前字符重复出现在当前子串中，需要更新子串的起始位置。
		if v, ok := m[s[i]]; ok && v >= start {
			maxLen = max(maxLen, i-start)
			start = v + 1
		}
		m[s[i]] = i
	}
	return max(maxLen, len(s)-start)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
