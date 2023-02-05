package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numTrees(t *testing.T) {
	got := numTrees(3)
	assert.Equal(t, 5, got)
}

func Test_sortColors(t *testing.T) {
	nums := []int{2, 0, 1}
	sortColors(nums)
	assert.Equal(t, []int{0, 1, 2}, nums)
	nums = []int{2, 0, 2, 1, 1, 0}
	sortColors(nums)
	assert.Equal(t, []int{0, 0, 1, 1, 2, 2}, nums)
}

func Test_groupAnagrams(t *testing.T) {
	ans := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	t.Log(ans)
}

func Test_canPartition(t *testing.T) {
	got := canPartition([]int{1, 5, 11, 5})
	assert.Equal(t, true, got)
	got = canPartition([]int{1, 2, 3, 5})
	assert.Equal(t, false, got)
	got = canPartition([]int{3, 3, 3, 4, 5})
	assert.Equal(t, true, got)
	got = canPartition([]int{1, 2, 5})
	assert.Equal(t, false, got)
	got = canPartition([]int{1, 2, 3, 4, 5, 6, 7})
	assert.Equal(t, true, got)
}
