package main

// 比特位计数 https://leetcode.cn/problems/counting-bits/description/
func countBits(n int) []int {
	ans := make([]int, n+1)
	countOnes := func(num int) int {
		count := 0
		for num > 0 {
			count++
			num = num & (num - 1)
		}
		return count
	}
	for i := 0; i <= n; i++ {
		ans[i] = countOnes(i)
	}

	return ans
}

// 课程表 https://leetcode.cn/problems/course-schedule/description/
func canFinish(numCourses int, prerequisites [][]int) bool {
	edges := make([][]int, numCourses)
	inDegree := make([]int, numCourses)
	for _, pre := range prerequisites {
		edges[pre[1]] = append(edges[pre[1]], pre[0])
		inDegree[pre[0]]++
	}
	queue := make([]int, 0)
	for i := 0; i < numCourses; i++ {
		if inDegree[i] == 0 {
			queue = append(queue, i)
		}
	}
	learned := 0
	for len(queue) > 0 {
		course := queue[0]
		queue = queue[1:]
		learned++
		for _, nextCourse := range edges[course] {
			inDegree[nextCourse]--
			if inDegree[nextCourse] == 0 {
				queue = append(queue, nextCourse)
			}
		}
	}
	return learned == numCourses
}

// 前 K 个高频元素 https://leetcode.cn/problems/top-k-frequent-elements/description/
func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, num := range nums {
		freqMap[num]++
	}

	var buckets = make([][]int, len(nums)+1)
	for num, freq := range freqMap {
		buckets[freq] = append(buckets[freq], num)
	}

	var result []int
	for i := len(buckets) - 1; i >= 0 && len(result) < k; i-- {
		if len(buckets[i]) > 0 {
			result = append(result, buckets[i]...)
		}
	}

	return result
}

// 最大正方形 https://leetcode.cn/problems/maximal-square/description/
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}
	rows := len(matrix)
	cols := len(matrix[0])
	maxSide := 0
	prev := 0
	dp := make([]int, cols)
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			temp := dp[j]
			if matrix[i][j] == '1' {
				if i == 0 || j == 0 {
					dp[j] = 1
				} else {
					dp[j] = min(min(dp[j-1], prev), dp[j]) + 1
				}
				maxSide = max(maxSide, dp[j])
			} else {
				dp[j] = 0
			}
			prev = temp
		}
	}
	return maxSide * maxSide
}

// 数组中的第K个最大元素 https://leetcode.cn/problems/kth-largest-element-in-an-array/description/
func findKthLargest(nums []int, k int) int {
    var quickSelect func(nums []int, left, right, k int) int
	quickSelect = func(nums []int, left, right, k int) int {
		pivot := nums[right]
		i := left
		for j := left; j < right; j++ {
			if nums[j] < pivot {
				nums[i], nums[j] = nums[j], nums[i]
				i++
			}
		}
		nums[i], nums[right] = nums[right], nums[i]

		if i == k {
			return nums[i]
		} else if i < k {
			return quickSelect(nums, i+1, right, k)
		} else {
			return quickSelect(nums, left, i-1, k)
		}
	}

	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}

// 数组中的第K个最大元素 https://leetcode.cn/problems/kth-largest-element-in-an-array/description/
func findKthLargest1(nums []int, k int) int {
	n := len(nums)
	var dfs func(int, int)
	dfs = func(i, n int) {
		l, r, largest := i*2+1, i*2+2, i
		if l < n && nums[l] > nums[largest] {
			largest = l
		}
		if r < n && nums[r] > nums[largest] {
			largest = r
		}
		if largest != i {
			nums[i], nums[largest] = nums[largest], nums[i]
			dfs(largest, n)
		}
	}
	// 建堆
	for i := n / 2; i >= 0; i-- {
		dfs(i, n)
	}
	for i := len(nums) - 1; i >= len(nums)-k+1; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		n--
		dfs(0, n)
	}
	return nums[0]
}
