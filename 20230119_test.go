package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_subarraySum(t *testing.T) {
	got := subarraySum([]int{1, 1, 1}, 2)
	assert.Equal(t, 2, got)
	got = subarraySum([]int{1, 2, 3}, 3)
	assert.Equal(t, 2, got)
}

func Test_maxArea(t *testing.T) {
	got := maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7})
	assert.Equal(t, 49, got)
	got = maxArea([]int{1, 1})
	assert.Equal(t, 1, got)
	got = maxArea([]int{1, 2, 4, 3})
	assert.Equal(t, 4, got)
	got = maxArea([]int{1, 8, 6, 2, 5, 4, 8, 25, 7})
	assert.Equal(t, 49, got)
	got = maxArea([]int{2, 3, 4, 5, 18, 17, 6})
	assert.Equal(t, 17, got)
}

func Test_findDuplicate(t *testing.T) {
	got := findDuplicate([]int{1, 3, 4, 2, 2})
	assert.Equal(t, 2, got)
	got = findDuplicate([]int{3, 1, 3, 4, 2})
	assert.Equal(t, 3, got)
	got = findDuplicate([]int{8, 7, 1, 10, 17, 15, 18, 11, 16, 9, 19, 12, 5, 14, 3, 4, 2, 13, 18, 18})
	assert.Equal(t, 18, got)
	got = findDuplicate([]int{4, 4, 17, 15, 2, 1, 19, 11, 12, 13, 3, 18, 4, 4, 5, 9, 7, 14, 4, 16})
	assert.Equal(t, 4, got)
}

func Test_dailyTemperatures(t *testing.T) {
	got := dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73})
	assert.Equal(t, []int{1, 1, 4, 2, 1, 1, 0, 0}, got)
	got = dailyTemperatures([]int{30, 40, 50, 60})
	assert.Equal(t, []int{1, 1, 1, 0}, got)
	got = dailyTemperatures([]int{30, 60, 90})
	assert.Equal(t, []int{1, 1, 0}, got)
}
func Test_moveZeroes(t *testing.T) {
	nums := []int{0, 1, 0, 3, 12}
	moveZeroes(nums)
	assert.Equal(t, []int{1, 3, 12, 0, 0}, nums)
	nums = []int{0}
	moveZeroes(nums)
	assert.Equal(t, []int{0}, nums)
}
