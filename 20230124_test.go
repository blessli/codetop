package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_productExceptSelf(t *testing.T) {
	got := productExceptSelf([]int{1, 2, 3, 4})
	assert.Equal(t, []int{24, 12, 8, 6}, got)
}

func Test_nextPermutation(t *testing.T) {
	nums := []int{1, 2, 3}
	nextPermutation(nums)
	assert.Equal(t, []int{1, 3, 2}, nums)
	nums = []int{3, 2, 1}
	nextPermutation(nums)
	assert.Equal(t, []int{1, 2, 3}, nums)
	nums = []int{1, 1, 5}
	nextPermutation(nums)
	assert.Equal(t, []int{1, 5, 1}, nums)
	nums = []int{1, 3, 2}
	nextPermutation(nums)
	assert.Equal(t, []int{2, 1, 3}, nums)
}

func Test_maxSlidingWindow(t *testing.T) {
	got := maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3)
	assert.Equal(t, []int{3, 3, 5, 5, 6, 7}, got)
	got = maxSlidingWindow([]int{1}, 1)
	assert.Equal(t, []int{1}, got)
}

func Test_trap(t *testing.T) {
	got := trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1})
	assert.Equal(t, 6, got)
	got = trap([]int{4, 2, 0, 3, 2, 5})
	assert.Equal(t, 9, got)
}
