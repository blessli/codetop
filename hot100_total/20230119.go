package main

import "fmt"

func moveZeroes(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}
	left := 0
	right := 1
	for left < right && left < n {
		if nums[left] == 0 {
			for right < n && nums[right] == 0 {
				right++
			}
			if right == n {
				return
			}
			temp := nums[right]
			nums[right] = 0
			nums[left] = temp
			left++
		} else {
			left++
		}
		if left == right {
			right++
		}
	}
}

func dailyTemperatures(temperatures []int) []int {
	ans := make([]int, 0)
	res := make([]int, len(temperatures))
	ans = append(ans, -1)
	for i, v := range temperatures {
		top := ans[len(ans)-1]
		if top < 0 {
			ans = append(ans, i)
			continue
		}
		if v <= temperatures[top] {
			ans = append(ans, i)
		} else {
			for ans[len(ans)-1] >= 0 && temperatures[ans[len(ans)-1]] < v {
				curr := ans[len(ans)-1]
				res[curr] = i - curr
				ans = ans[:len(ans)-1]
			}
			ans = append(ans, i)
		}
	}
	return res
}

func findDuplicate(nums []int) int {
	for i := 0; i < len(nums); i++ {
		pos := i
		temp := nums[pos]
		if temp == nums[temp-1] {
			if pos+1 != temp {
				return temp
			}
			continue
		}
		for {
			nums[pos] = nums[temp-1]
			nums[temp-1] = temp
			temp = nums[pos]
			if temp == nums[temp-1] {
				if pos+1 != temp {
					return temp
				}
				break
			}
		}
	}
	fmt.Println(nums)
	return 0
}
// 盛最多水的容器
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	ans := 0
	for left < right {
		ans = max(ans, min(height[left], height[right])*(right-left))
		if height[left] <= height[right] {
			left++
		} else {
			right--
		}
	}
	return ans
}

func subarraySum(nums []int, k int) int {
	ans := 0
	pre := 0
	m := map[int]int{}
	m[0] = 1
	for i := 0; i < len(nums); i++ {
		pre += nums[i]
		if c, ok := m[pre-k]; ok {
			ans += c
		}
		m[pre]++
	}
	return ans
}
