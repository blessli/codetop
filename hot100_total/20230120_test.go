package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_combinationSum(t *testing.T) {
	got := combinationSum([]int{2, 3, 6, 7}, 7)
	assert.Equal(t, [][]int{{2, 2, 3}, {7}}, got)
	got = combinationSum([]int{2, 3, 5}, 8)
	assert.Equal(t, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}, got)
	got = combinationSum([]int{2}, 1)
	assert.Equal(t, [][]int{}, got)
}

func Test_maxProduct(t *testing.T) {
	got := maxProduct([]int{2, 3, -2, 4})
	assert.Equal(t, 6, got)
	got = maxProduct([]int{-2, 0, -1})
	assert.Equal(t, 0, got)
	got = maxProduct([]int{-3, -1, -1})
	assert.Equal(t, 3, got)
	got = maxProduct([]int{0, 2})
	assert.Equal(t, 2, got)
	got = maxProduct([]int{-2, 3, -4})
	assert.Equal(t, 24, got)
}

func Test_rob(t *testing.T) {
	got := rob([]int{1, 2, 3, 1})
	assert.Equal(t, 4, got)
	got = rob([]int{2, 7, 9, 3, 1})
	assert.Equal(t, 12, got)
}
