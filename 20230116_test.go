package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_numSquares(t *testing.T) {
	got := numSquares(12)
	assert.Equal(t, 3, got)
	got = numSquares(13)
	assert.Equal(t, 2, got)
}

func Test_uniquePaths(t *testing.T) {
	got := uniquePaths(3, 7)
	assert.Equal(t, 28, got)
	got = uniquePaths(7, 3)
	assert.Equal(t, 28, got)
	got = uniquePaths(3, 2)
	assert.Equal(t, 3, got)
	got = uniquePaths(3, 3)
	assert.Equal(t, 6, got)
}

func Test_uniquePaths_(t *testing.T) {
	got := uniquePaths_(3, 7)
	assert.Equal(t, 28, got)
	got = uniquePaths_(7, 3)
	assert.Equal(t, 28, got)
	got = uniquePaths_(3, 2)
	assert.Equal(t, 3, got)
	got = uniquePaths_(3, 3)
	assert.Equal(t, 6, got)
}

func Test_longestPalindrome(t *testing.T) {
	got := longestPalindrome("babad")
	assert.Equal(t, "bab", got)
	got = longestPalindrome("cbbd")
	assert.Equal(t, "bb", got)
	got = longestPalindrome("a")
	assert.Equal(t, "a", got)
	got = longestPalindrome("aacabdkacaa")
	assert.Equal(t, "aca", got)
}

func Test_longestPalindrome_(t *testing.T) {
	got := longestPalindrome_("babad")
	assert.Equal(t, "bab", got)
	got = longestPalindrome_("cbbd")
	assert.Equal(t, "bb", got)
	got = longestPalindrome_("a")
	assert.Equal(t, "a", got)
	got = longestPalindrome_("aacabdkacaa")
	assert.Equal(t, "aca", got)
}
