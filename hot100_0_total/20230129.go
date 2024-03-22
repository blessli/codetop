package main

import (
	"sort"
	"strings"
)

// 打家劫舍 III https://leetcode.cn/problems/house-robber-iii/description/
func rob3(root *TreeNode) int {
	var robSub func(node *TreeNode) [2]int
	robSub = func(node *TreeNode) [2]int {
		if node == nil {
			return [2]int{0, 0}
		}

		left := robSub(node.Left)
		right := robSub(node.Right)

		result := [2]int{0, 0}
		result[0] = max(left[0], left[1]) + max(right[0], right[1])
		result[1] = node.Val + left[0] + right[0]

		return result
	}
	result := robSub(root)
	return max(result[0], result[1])
}

// 根据身高重建队列 https://leetcode.cn/problems/queue-reconstruction-by-height/description/
func reconstructQueue(people [][]int) [][]int {
	sort.Slice(people, func(i, j int) bool {
		a, b := people[i], people[j]
		return a[0] > b[0] || (a[0] == b[0] && a[1] < b[1])
	})
	ans := people
	for _, v := range people {
		idx := v[1]
		ans = append(ans[:idx], append([][]int{v}, ans[idx:]...)...)
	}
	return ans
}

// 戳气球 https://leetcode.cn/problems/burst-balloons/description/
func maxCoins(nums []int) int {
	n := len(nums)
	nums = append([]int{1}, nums...)
	nums = append(nums, 1)
	dp := make([][]int, n+2)
	for i := 0; i < n+2; i++ {
		dp[i] = make([]int, n+2)
	}
	for i := 0; i < n+2; i++ {
		for j := 0; j < n+2; j++ {
			dp[i][j] = -1
		}
	}
	var dfs func(int, int) int
	dfs = func(left, right int) int {
		if left >= right-1 {
			return 0
		}
		if dp[left][right] != -1 {
			return dp[left][right]
		}
		for i := left + 1; i < right; i++ {
			v := nums[left] * nums[i] * nums[right]
			v += dfs(left, i) + dfs(i, right)
			dp[left][right] = max(dp[left][right], v)
		}
		return dp[left][right]
	}
	return dfs(0, n+1)
}

// 寻找两个正序数组的中位数 https://leetcode.cn/problems/median-of-two-sorted-arrays/description/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	totalLen := len(nums1) + len(nums2)
	if totalLen%2 == 1 {
		return float64(findKth(nums1, nums2, totalLen/2+1))
	} else {
		return float64(findKth(nums1, nums2, totalLen/2)+findKth(nums1, nums2, totalLen/2+1)) / 2
	}
}

func findKth(nums1, nums2 []int, k int) int {
	if len(nums1) > len(nums2) {
		return findKth(nums2, nums1, k)
	}
	if len(nums1) == 0 {
		return nums2[k-1]
	}
	if k == 1 {
		return min(nums1[0], nums2[0])
	}

	pa := min(k/2, len(nums1))
	pb := k - pa
	if nums1[pa-1] < nums2[pb-1] {
		return findKth(nums1[pa:], nums2, k-pa)
	}
	return findKth(nums1, nums2[pb:], k-pb)
}

// 两两交换链表中的节点 https://leetcode.cn/problems/swap-nodes-in-pairs/description/
func swapPairs(head *ListNode) *ListNode {
	dummy := &ListNode{Next: head}
	prev := dummy

	for prev.Next != nil && prev.Next.Next != nil {
		first := prev.Next
		second := prev.Next.Next

		// Swap nodes
		prev.Next = second
		first.Next = second.Next
		second.Next = first

		// Move to next pair
		prev = first
	}

	return dummy.Next
}

// K 个一组翻转链表 https://leetcode.cn/problems/reverse-nodes-in-k-group/description/
func reverseKGroup(head *ListNode, k int) *ListNode {
	dummy := &ListNode{Next: head}
	pre := dummy
	for head != nil {
		tail := pre
		for i := 0; i < k; i++ {
			tail = tail.Next
			if tail == nil {
				return dummy.Next
			}
		}
		next := tail.Next
		head, tail = reverse(head, tail)
		pre.Next = head
		tail.Next = next
		pre = tail
		head = tail.Next
	}
	return dummy.Next
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
	pre := tail.Next
	cur := head
	for pre != tail {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return tail, head
}

// 合并两个有序数组 https://leetcode.cn/problems/merge-sorted-array/description/
func merge123(nums1 []int, m int, nums2 []int, n int) {
	i, j, k := m-1, n-1, m+n-1
	for i >= 0 && j >= 0 {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
		k--
	}

	for j >= 0 {
		nums1[k] = nums2[j]
		j--
		k--
	}
}

// 最长公共前缀 https://leetcode.cn/problems/longest-common-prefix/description/
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	prefix := strs[0]
	for i := 1; i < len(strs); i++ {
		for len(prefix) > 0 {
			if strings.HasPrefix(strs[i], prefix) {
				break
			} else {
				prefix = prefix[:len(prefix)-1]
			}
		}
	}

	return prefix
}

// 路径总和 https://leetcode.cn/problems/path-sum/description/
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil && root.Val == targetSum {
		return true
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

// 删除无效的括号 https://leetcode.cn/problems/remove-invalid-parentheses/description/
func removeInvalidParentheses(s string) []string {
    var results []string
    var visited = make(map[string]bool)
    var queue []string
    var found bool

    queue = append(queue, s)
    visited[s] = true

    for len(queue) > 0 {
        current := queue[0]
        queue = queue[1:]

        if isValid2(current) {
            results = append(results, current)
            found = true
        }

        if found {
            continue
        }

        for i := 0; i < len(current); i++ {
            if current[i] != '(' && current[i] != ')' {
                continue
            }

            next := current[:i] + current[i+1:]
            if !visited[next] {
                queue = append(queue, next)
                visited[next] = true
            }
        }
    }

    return results
}

func isValid2(s string) bool {
    count := 0
    for i := 0; i < len(s); i++ {
        if s[i] == '(' {
            count++
        } else if s[i] == ')' {
            count--
            if count < 0 {
                return false
            }
        }
    }
    return count == 0
}