package main

import "sort"

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

func isValid(s string) bool {
	mapping := map[byte]byte{
		')': '(',
		'}': '{',
		']': '[',
	}
	ans := true
	stack := []byte{}
	stack = append(stack, '0')
	for i := 0; i < len(s); i++ {
		key := s[i]
		if key == '(' || key == '{' || key == '[' {
			stack = append(stack, key)
		} else {
			if v, ok := mapping[key]; !ok || v != stack[len(stack)-1] {
				ans = false
				break
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	if len(stack) != 1 {
		ans = false
	}
	return ans
}

func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}
	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && i < len(nums) && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for j < k {
			sum := nums[i] + nums[j] + nums[k]
			if sum == 0 {
				ans = append(ans, []int{nums[i], nums[j], nums[k]})
				for j < k && j < len(nums)-1 && nums[j] == nums[j+1] {
					j++
				}
				for j < k && j < len(nums)-1 && nums[k] == nums[k-1] {
					k--
				}
			}
			if sum > 0 {
				k--
			} else {
				j++
			}
		}

	}
	return ans
}

func maxSubArray(nums []int) int {
	size := len(nums)
	if size == 1 {
		return nums[0]
	}
	dp := make([]int, len(nums))
	maxi := nums[0]
	dp[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		if dp[i-1] > 0 {
			dp[i] = dp[i-1] + nums[i]
		} else {
			dp[i] = nums[i]
		}
		maxi = max(maxi, dp[i])
	}
	return maxi
}

func lengthOfLongestSubstring(s string) int {
	m := map[byte]int{}
	maxi := 0
	pre := 0
	for i := 0; i < len(s); i++ {
		key := s[i]
		if v, ok := m[key]; ok && v >= pre {
			maxi = max(maxi, i-pre)
			pre = v + 1
		}
		m[key] = i
	}
	return max(maxi, len(s)-pre)
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
