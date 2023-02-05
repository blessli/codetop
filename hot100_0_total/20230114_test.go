package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_maxProfit(t *testing.T) {
	got := maxProfit([]int{7, 1, 5, 3, 6, 4})
	assert.Equal(t, 5, got)
	got = maxProfit([]int{7, 6, 4, 3, 1})
	assert.Equal(t, 0, got)
}

func Test_lengthOfLIS(t *testing.T) {
	got := lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})
	assert.Equal(t, 4, got)
	got = lengthOfLIS([]int{0, 1, 0, 3, 2, 3})
	assert.Equal(t, 4, got)
	got = lengthOfLIS([]int{7, 7, 7, 7, 7, 7, 7})
	assert.Equal(t, 1, got)
	got = lengthOfLIS2([]int{10, 9, 2, 5, 3, 7, 101, 18})
	assert.Equal(t, 4, got)
	got = lengthOfLIS2([]int{0, 1, 0, 3, 2, 3})
	assert.Equal(t, 4, got)
	got = lengthOfLIS2([]int{7, 7, 7, 7, 7, 7, 7})
	assert.Equal(t, 1, got)
}

func Test_coinChange(t *testing.T) {
	got := coinChange([]int{1, 2, 5}, 11)
	assert.Equal(t, 3, got)
	got = coinChange([]int{2}, 3)
	assert.Equal(t, -1, got)
	got = coinChange([]int{1}, 0)
	assert.Equal(t, 0, got)
	got = coinChange([]int{411, 412, 413, 414, 415, 416, 417, 418, 419, 420, 421, 422}, 9864)
	assert.Equal(t, 24, got)
}

func Test_merge(t *testing.T) {
	got := merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	t.Log(got)
	got = merge([][]int{{1, 4}, {4, 5}})
	t.Log(got)
}
