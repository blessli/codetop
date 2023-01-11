package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxSubArray(t *testing.T) {
	got := maxSubArray([]int{5, 4, -1, 7, 8})
	assert.Equal(t, 23, got)
	got = maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4})
	assert.Equal(t, 6, got)
}

func Test_lengthOfLongestSubstring(t *testing.T) {
	got := lengthOfLongestSubstring("pwwkew")
	assert.Equal(t, 3, got)
	got = lengthOfLongestSubstring("abcabcbb")
	assert.Equal(t, 3, got)
}

func Test_threeSum(t *testing.T) {
	got := threeSum([]int{-1, 0, 1, 2, -1, -4})
	t.Log(got)
	got = threeSum([]int{0, 1, 1})
	t.Log(got)
	got = threeSum([]int{0, 0, 0})
	t.Log(got)
}

func Test_isValid(t *testing.T) {
	got := isValid("()")
	assert.Equal(t, true, got)
	got = isValid("()[]{}")
	assert.Equal(t, true, got)
	got = isValid("(]")
	assert.Equal(t, false, got)
	got = isValid("((((()))))[]{}")
	assert.Equal(t, true, got)
}

func Test_twoSum(t *testing.T) {
	got := twoSum([]int{2, 7, 11, 15}, 9)
	assert.Equal(t, []int{0, 1}, got)
	got = twoSum([]int{3, 2, 4}, 6)
	assert.Equal(t, []int{1, 2}, got)
	got = twoSum([]int{3, 3}, 6)
	assert.Equal(t, []int{0, 1}, got)
}
