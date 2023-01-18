package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_majorityElement(t *testing.T) {
	got := majorityElement([]int{3, 2, 3})
	assert.Equal(t, 3, got)
	got = majorityElement([]int{2, 2, 1, 1, 1, 2, 2})
	assert.Equal(t, 2, got)
}

func Test_permute(t *testing.T) {
	got := permute([]int{1, 2, 3})
	t.Log(got)
	got = permute([]int{0, 1})
	t.Log(got)
	got = permute([]int{1})
	t.Log(got)
	got = permute([]int{5, 4, 6, 2})
	t.Log(got)
}

func Test_search1(t *testing.T) {
	got := search1([]int{4, 5, 6, 7, 0, 1, 2}, 0)
	assert.Equal(t, 4, got)
	got = search1([]int{4, 5, 6, 7, 0, 1, 2}, 3)
	assert.Equal(t, -1, got)
	got = search1([]int{1}, 0)
	assert.Equal(t, -1, got)
	got = search1([]int{5, 1, 3}, 5)
	assert.Equal(t, 0, got)
	got = search1([]int{3, 4, 5, 6, 1, 2}, 2)
	assert.Equal(t, 5, got)
	got = search1([]int{3, 5, 1}, 3)
	assert.Equal(t, 0, got)
	got = search1([]int{4, 5, 6, 7, 8, 1, 2, 3}, 8)
	assert.Equal(t, 4, got)
}

func Test_canJump(t *testing.T) {
	got := canJump([]int{2, 3, 1, 1, 4})
	assert.Equal(t, true, got)
	got = canJump([]int{3, 2, 1, 0, 4})
	assert.Equal(t, false, got)
	got = canJump([]int{1, 2})
	assert.Equal(t, true, got)
	got = canJump([]int{0, 2, 3})
	assert.Equal(t, false, got)
	got = canJump([]int{2, 0, 0})
	assert.Equal(t, true, got)

}
