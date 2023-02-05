package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_longestValidParentheses(t *testing.T) {
	got := longestValidParentheses("(()")
	assert.Equal(t, 2, got)
	got = longestValidParentheses(")()())")
	assert.Equal(t, 4, got)
}

func Test_countSubstrings(t *testing.T) {
	got := countSubstrings("abc")
	assert.Equal(t, 3, got)
	got = countSubstrings("aaa")
	assert.Equal(t, 6, got)
}

func Test_maxProfit2(t *testing.T) {
	got := maxProfit2([]int{1, 2, 3, 0, 2})
	assert.Equal(t, 3, got)
	got = maxProfit2([]int{1})
	assert.Equal(t, 0, got)
}
