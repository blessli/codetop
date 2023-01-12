package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numIslands(t *testing.T) {
	grid := [][]byte{{'1', '1', '1', '1', '0'}, {'1', '1', '0', '1', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '0', '0', '0'}}
	got := numIslands(grid)
	assert.Equal(t, 1, got)
	grid = [][]byte{{'1', '1', '0', '0', '0'}, {'1', '1', '0', '0', '0'}, {'0', '0', '1', '0', '0'}, {'0', '0', '0', '1', '1'}}
	got = numIslands(grid)
	assert.Equal(t, 3, got)
}

// func Test_reverseList(t *testing.T) {
// 	head := &ListNode{
// 		Val:  1,
// 		Next: &ListNode{},
// 	}
// 	curr := head
// 	for i := 2; i <= 5; i++ {
// 		newNode := &ListNode{
// 			Val:  i,
// 			Next: &ListNode{},
// 		}
// 		head.Next = newNode
// 		head = head.Next
// 	}
// 	got := reverseList(curr)
// 	body, _ := json.Marshal(got)
// 	t.Logf("reverseList: %+v", string(body))
// }

func Test_minPathSum(t *testing.T) {
	grid := [][]int{{1, 3, 1}, {1, 5, 1}, {4, 2, 1}}
	got := minPathSum(grid)
	assert.Equal(t, 7, got)
	grid = [][]int{{1, 2, 3}, {4, 5, 6}}
	got = minPathSum(grid)
	assert.Equal(t, 12, got)
}

func Test_climbStairs(t *testing.T) {
	got := climbStairs(2)
	assert.Equal(t, 2, got)
	got = climbStairs(3)
	assert.Equal(t, 3, got)
}
