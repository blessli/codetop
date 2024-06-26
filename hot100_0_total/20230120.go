package main

// 组合总和 https://leetcode.cn/problems/combination-sum/description/
func combinationSum(candidates []int, target int) [][]int {
	n := len(candidates)
	ans := [][]int{}
	var dfs func(int, int, []int)
	dfs = func(start int, target int, path []int) {
		if target == 0 {
			tmp := make([]int, len(path))
			copy(tmp, path)
			ans = append(ans, tmp)
			return
		}
		for i := start; i < n; i++ {
			if candidates[i] > target {
				continue
			}
			path = append(path, candidates[i])
			dfs(i, target-candidates[i], path)
			path = path[:len(path)-1]
		}
	}
	dfs(0, target, []int{})
	return ans
}

// 乘积最大子数组 https://leetcode.cn/problems/maximum-product-subarray/description/
func maxProduct(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	maxProd, minProd, result := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] < 0 {
			maxProd, minProd = minProd, maxProd
		}
		maxProd = max(nums[i], maxProd*nums[i])
		minProd = min(nums[i], minProd*nums[i])
		result = max(result, maxProd)
	}
	return result
}

// 打家劫舍 https://leetcode.cn/problems/house-robber/description/
func rob(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}
	return dp[n-1]
}

// 打家劫舍 https://leetcode.cn/problems/house-robber/description/
func rob1(nums []int) int {
	n := len(nums)
	if n == 0 {
		return 0
	}
	if n == 1 {
		return nums[0]
	}
	prev1 := 0 // 存储到当前位置为止，包括当前位置的房屋能够偷到的最大金额。
	prev2 := 0 // 存储到前一个位置为止（不包括当前位置的房屋）能够偷到的最大金额。
	for _, num := range nums {
		temp := prev1
		prev1 = max(prev1, prev2+num)
		prev2 = temp
	}
	return prev1
}
// 合并两个有序链表 https://leetcode.cn/problems/merge-two-sorted-lists/description/
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			current.Next = list1
			list1 = list1.Next
		} else {
			current.Next = list2
			list2 = list2.Next
		}
		current = current.Next
	}
	if list1 != nil {
		current.Next = list1
	}
	if list2 != nil {
		current.Next = list2
	}
	return dummy.Next
}

type MinStack struct {
    stack []int
    minStack []int
}
func Constructor2() MinStack {
    return MinStack{}
}

func (this *MinStack) Push(val int) {
    this.stack = append(this.stack, val)
    if len(this.minStack) == 0 || val <= this.minStack[len(this.minStack)-1] {
        this.minStack = append(this.minStack, val)
    }
}

func (this *MinStack) Pop() {
    if this.stack[len(this.stack)-1] == this.minStack[len(this.minStack)-1] {
        this.minStack = this.minStack[:len(this.minStack)-1]
    }
    this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
    return this.stack[len(this.stack)-1]
}

func (this *MinStack) GetMin() int {
    return this.minStack[len(this.minStack)-1]
}