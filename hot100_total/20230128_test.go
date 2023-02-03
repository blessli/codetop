package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findUnsortedSubarray(t *testing.T) {
	got := findUnsortedSubarray([]int{2, 6, 4, 8, 10, 9, 15})
	assert.Equal(t, 5, got)
	got = findUnsortedSubarray([]int{1, 2, 3, 4})
	assert.Equal(t, 0, got)
	got = findUnsortedSubarray([]int{1})
	assert.Equal(t, 0, got)
}

func Test_findAnagrams(t *testing.T) {
	got := findAnagrams("cbaebabacd", "abc")
	assert.Equal(t, []int{0, 6}, got)
	got = findAnagrams("abab", "ab")
	assert.Equal(t, []int{0, 1, 2}, got)
}
