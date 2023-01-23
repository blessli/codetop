package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_largestRectangleArea(t *testing.T) {
	got := largestRectangleArea([]int{2, 1, 5, 6, 2, 3})
	assert.Equal(t, 10, got)
	got = largestRectangleArea([]int{2, 4})
	assert.Equal(t, 4, got)
	got = largestRectangleArea([]int{1})
	assert.Equal(t, 1, got)
	got = largestRectangleArea([]int{9, 0})
	assert.Equal(t, 9, got)
	got = largestRectangleArea([]int{2, 0, 2})
	assert.Equal(t, 2, got)
	got = largestRectangleArea([]int{2, 1, 2})
	assert.Equal(t, 3, got)
	got = largestRectangleArea([]int{4, 2, 0, 3, 2, 5})
	assert.Equal(t, 6, got)
	got = largestRectangleArea([]int{3, 6, 5, 7, 4, 8, 1, 0})
	assert.Equal(t, 20, got)
}

func Test_findTargetSumWays(t *testing.T) {
	got := findTargetSumWays([]int{1, 1, 1, 1, 1}, 3)
	assert.Equal(t, 5, got)
	got = findTargetSumWays([]int{1}, 1)
	assert.Equal(t, 1, got)
	got = findTargetSumWays([]int{7, 7, 17, 1, 46, 38, 8, 32, 35, 18, 43, 48, 9, 17, 6, 6, 42, 10, 2, 32}, 38)
	assert.Equal(t, 6462, got)
}

func Test_rotate(t *testing.T) {
	matrix := [][]int{{5, 1, 9, 11}, {2, 4, 8, 10}, {13, 3, 6, 7}, {15, 14, 12, 16}}
	rotate(matrix)
	t.Log(matrix)
}

func Test_generateParenthesis(t *testing.T) {
	got := generateParenthesis(3)
	t.Log(got)
	got = generateParenthesis(1)
	t.Log(got)
}
