package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findKthLargest(t *testing.T) {
	got := findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2)
	assert.Equal(t, 5, got)
	got = findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4)
	assert.Equal(t, 4, got)
}

func Test_topKFrequent(t *testing.T) {
	got := topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)
	assert.Equal(t, []int{1, 2}, got)
	got = topKFrequent([]int{1}, 1)
	assert.Equal(t, []int{1}, got)
}

func Test_canFinish(t *testing.T) {
	got := canFinish(2, [][]int{{1, 0}})
	assert.Equal(t, true, got)
	got = canFinish(2, [][]int{{1, 0}, {0, 1}})
	assert.Equal(t, false, got)
	got = canFinish(5, [][]int{{1, 4}, {2, 4}, {3, 1}, {3, 2}})
	assert.Equal(t, true, got)
	got = canFinish(3, [][]int{{1, 0}, {2, 0}, {0, 2}})
	assert.Equal(t, false, got)
}

// func Test_findUnsortedSubarray(t *testing.T) {
// 	got := findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15})
// 	assert.Equal(t, 5, got)
// 	got = findUnsortedSubarray([]int{1, 2, 3, 4})
// 	assert.Equal(t, 0, got)
// 	got = findUnsortedSubarray([]int{1})
// 	assert.Equal(t, 0, got)
// }

func Test_countBits(t *testing.T) {
	got := countBits(5)
	assert.Equal(t, []int{0, 1, 1, 2, 1, 2}, got)
	got = countBits(2)
	assert.Equal(t, []int{0, 1, 1}, got)
	got = countBits(16)
	assert.Equal(t, []int{0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4, 1}, got)
}
