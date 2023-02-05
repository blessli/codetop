package main

// todo:dp
func countBits(n int) []int {
	data := []int{}
	if n == 0 {
		return data
	}
	m := map[int]bool{}
	for i := 2; i <= n; i = i * 2 {
		m[i] = true
		data = append(data, i)
	}
	dp := make([]int, n+1)
	dp[1] = 1
	dfs := func(curr int) int {
		cnt := 0
		for i := len(data) - 1; i >= 0 && curr > 0; i-- {
			if curr >= data[i] {
				cnt += dp[data[i]]
				curr -= data[i]
			}
		}
		return cnt
	}

	for i := 2; i <= n; i++ {
		if i%2 == 1 {
			dp[i] = dp[i-1] + 1
		} else {
			if m[i] {
				dp[i] = 1
			} else {
				dp[i] = dfs(i)
			}
		}
	}
	return dp
}

func canFinish(numCourses int, prerequisites [][]int) bool {
	m := make([][]int, numCourses)
	for _, edge := range prerequisites {
		a, b := edge[0], edge[1]
		if len(m[b]) == 0 {
			m[b] = make([]int, 0)
		}
		m[b] = append(m[b], a)
	}
	hasCycle := false
	vis := make([]int, numCourses)
	f := map[int]bool{}
	var dfs func(int)
	dfs = func(pos int) {
		vis[pos] = 1
		if hasCycle {
			return
		}
		f[pos] = true
		for _, v := range m[pos] {
			if vis[v] == 1 {
				hasCycle = true
				return
			}
			if vis[v] == 0 {
				dfs(v)
			}
		}
		vis[pos] = 2
	}
	for i, _ := range m {
		if len(m[i]) == 0 {
			continue
		}
		dfs(i)
		if hasCycle {
			return false
		}
	}
	return true
}

func topKFrequent(nums []int, k int) []int {
	m := map[int]int{}
	data := []int{}
	for _, v := range nums {
		if _, ok := m[v]; !ok {
			data = append(data, v)
		}
		m[v]++
	}
	var dfs func(int, int)
	dfs1 := func(low, high int) int {
		start := data[low]
		left, right := low, high
		for left < right {
			for left < right && m[data[right]] <= m[start] {
				right--
			}
			data[left] = data[right]
			for left < right && m[data[left]] >= m[start] {
				left++
			}
			data[right] = data[left]
		}
		data[left] = start
		return left
	}
	dfs = func(low, high int) {
		if low >= high || low > k {
			return
		}
		mid := dfs1(low, high)
		dfs(low, mid-1)
		dfs(mid+1, high)
	}
	dfs(0, len(data)-1)
	return data[:k]
}

// 最大正方形 dp
func maximalSquare(matrix [][]byte) int {
	m := len(matrix)
	n := len(matrix[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}
	ans := 0
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == '0' {
				continue
			}
			if i*j == 0 {
				dp[i][j] = 1
				ans = max(ans, 1)
				continue
			}
			dp[i][j] = min(dp[i-1][j-1], min(dp[i][j-1], dp[i-1][j])) + 1
			ans = max(ans, dp[i][j])
		}
	}
	return ans * ans
}

// 数组中的第K个最大元素 快排
func findKthLargest_(nums []int, k int) int {
	var dfs func(int, int)
	dfs1 := func(low, high int) int {
		start := nums[low]
		left, right := low, high
		for left < right {
			for left < right && nums[right] <= start {
				right--
			}
			nums[left] = nums[right]
			for left < right && nums[left] >= start {
				left++
			}
			nums[right] = nums[left]
		}
		nums[left] = start
		return left
	}
	dfs = func(low, high int) {
		if low >= high || low >= k {
			return
		}
		mid := dfs1(low, high)
		dfs(low, mid-1)
		dfs(mid+1, high)
	}
	dfs(0, len(nums)-1)
	return nums[k-1]
}
// 大顶堆
func findKthLargest(nums []int, k int) int {
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
