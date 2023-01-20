package main

import "math"

func combinationSum(candidates []int, target int) [][]int {
	n := len(candidates)
	ans := make([][]int, 0)
	var dfs func([]int, int, int)
	dfs = func(curr []int, sum int, pos int) {
		if sum == target {
			temp := make([]int, 0)
			temp = append(temp, curr...)
			ans = append(ans, temp)
		}
		for i := pos; i < n; i++ {
			v := candidates[i]
			if sum+v <= target {
				curr = append(curr, v)
				dfs(curr, sum+v, i)
				curr = curr[:len(curr)-1]
			}
		}
	}
	dfs([]int{}, 0, 0)
	return ans
}

func maxProduct(nums []int) int {
	ans := math.MinInt32
	n := len(nums)
	mindp := []int{}
	maxdp := []int{}
	mindp = append(mindp, nums...)
	maxdp = append(maxdp, nums...)
	ans = max(ans, maxdp[0])
	for i := 1; i < n; i++ {
		maxdp[i] = max(maxdp[i-1]*nums[i], max(nums[i], mindp[i-1]*nums[i]))
		mindp[i] = min(mindp[i-1]*nums[i], min(nums[i], maxdp[i-1]*nums[i]))
		ans = max(ans, maxdp[i])
	}
	return ans
}

func rob(nums []int) int {
	n := len(nums)
	dp := []int{}
	dp = append(dp, nums...)
	ans := 0
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if j+1 < i {
				dp[i] = max(dp[i], dp[j]+nums[i])
			}
		}
		ans = max(ans, dp[i])
	}
	return ans
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	curr := &ListNode{}
	ans := curr
	for list1 != nil || list2 != nil {
		if list1 != nil && list2 != nil {
			if list1.Val <= list2.Val {
				temp := &ListNode{Val: list1.Val}
				curr.Next = temp
				curr = curr.Next
				list1 = list1.Next
			} else {
				temp := &ListNode{Val: list2.Val}
				curr.Next = temp
				curr = curr.Next
				list2 = list2.Next
			}
		} else if list1 != nil {
			temp := &ListNode{Val: list1.Val}
			curr.Next = temp
			curr = curr.Next
			list1 = list1.Next
		} else {
			temp := &ListNode{Val: list2.Val}
			curr.Next = temp
			curr = curr.Next
			list2 = list2.Next
		}
	}
	return ans.Next
}

type MinStack struct {
	data    []int
	mindata []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
	if len(this.mindata) == 0 || val < this.mindata[len(this.mindata)-1] {
		this.mindata = append(this.mindata, val)
	}
}

func (this *MinStack) Pop() {
	top := this.Top()
	if top == this.mindata[len(this.mindata)-1] {
		this.mindata = this.mindata[:len(this.mindata)-1]
	}
	this.data = this.data[:len(this.data)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.mindata[len(this.mindata)-1]
}
