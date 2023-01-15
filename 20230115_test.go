package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_searchRange(t *testing.T) {
	got := searchRange([]int{5, 7, 7, 8, 8, 10}, 8)
	assert.Equal(t, []int{3, 4}, got)
	got = searchRange([]int{5, 7, 7, 8, 8, 10}, 6)
	assert.Equal(t, []int{-1, -1}, got)
	got = searchRange([]int{}, 0)
	assert.Equal(t, []int{-1, -1}, got)
	got = searchRange([]int{1}, 1)
	assert.Equal(t, []int{0, 0}, got)
	got = searchRange([]int{2, 2}, 1)
	assert.Equal(t, []int{-1, -1}, got)
	got = searchRange([]int{2, 2}, 2)
	assert.Equal(t, []int{0, 1}, got)
}

func Test_searchMatrix(t *testing.T) {
	got := searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 5)
	assert.Equal(t, true, got)
	got = searchMatrix([][]int{{1, 4, 7, 11, 15}, {2, 5, 8, 12, 19}, {3, 6, 9, 16, 22}, {10, 13, 14, 17, 24}, {18, 21, 23, 26, 30}}, 20)
	assert.Equal(t, false, got)
}

func Test_longestConsecutive(t *testing.T) {
	got := longestConsecutive([]int{100, 4, 200, 1, 3, 2})
	assert.Equal(t, 4, got)
	got = longestConsecutive([]int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1})
	assert.Equal(t, 9, got)
}
