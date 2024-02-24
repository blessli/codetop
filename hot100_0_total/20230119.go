package main

// 移动零 https://leetcode.cn/problems/move-zeroes/description/
func moveZeroes(nums []int) {
	zeroIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[i], nums[zeroIndex] = nums[zeroIndex], nums[i]
			zeroIndex++
		}
	}
}

// 每日温度 https://leetcode.cn/problems/daily-temperatures/description/
func dailyTemperatures(temperatures []int) []int {
	n := len(temperatures)
	stack := []int{}
	answer := make([]int, n)
	for i := 0; i < n; i++ {
		for len(stack) > 0 && temperatures[i] > temperatures[stack[len(stack)-1]] {
			prevIndex := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			answer[prevIndex] = i - prevIndex
		}
		stack = append(stack, i)
	}
	return answer
}

// 寻找重复数 https://leetcode.cn/problems/find-the-duplicate-number/description/
func findDuplicate(nums []int) int {
	slow, fast := nums[0], nums[0]
	for {
		slow = nums[slow]
		fast = nums[nums[fast]]
		if slow == fast {
			break
		}
	}
	slow = nums[0]
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}
	return slow
}

// 盛最多水的容器 https://leetcode.cn/problems/container-with-most-water/description/
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0
	for left < right {
		h := min(height[left], height[right])
		w := right - left
		area := h * w
		maxArea = max(maxArea, area)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

// 和为 K 的子数组 https://leetcode.cn/problems/subarray-sum-equals-k/description/
func subarraySum(nums []int, k int) int {
	count := 0
	prefixSum := 0
	sumCount := make(map[int]int)
	sumCount[0] = 1
	for _, num := range nums {
		prefixSum += num
		count += sumCount[prefixSum-k]
		sumCount[prefixSum]++
	}
	return count
}
