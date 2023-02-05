package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_findDisappearedNumbers(t *testing.T) {
	got := findDisappearedNumbers([]int{4, 3, 2, 7, 8, 2, 3, 1})
	assert.Equal(t, []int{5, 6}, got)
	got = findDisappearedNumbers([]int{1, 1})
	assert.Equal(t, []int{2}, got)
}

func Test_hammingDistance(t *testing.T) {
	got := hammingDistance(1, 4)
	assert.Equal(t, 2, got)
	got = hammingDistance(3, 1)
	assert.Equal(t, 1, got)
}

func Test_decodeString(t *testing.T) {

	got := decodeString("3[a]2[bc]")
	assert.Equal(t, "aaabcbc", got)
	got = decodeString("3[a2[c]]")
	assert.Equal(t, "accaccacc", got)
	got = decodeString("2[abc]3[cd]ef")
	assert.Equal(t, "abcabccdcdcdef", got)
	got = decodeString("abc3[cd]xyz")
	assert.Equal(t, "abccdcdcdxyz", got)
	got = decodeString("3[z]2[2[y]pq4[2[jk]e1[f]]]ef")
	assert.Equal(t, "zzzyypqjkjkefjkjkefjkjkefjkjkefyypqjkjkefjkjkefjkjkefjkjkefef", got)
}
