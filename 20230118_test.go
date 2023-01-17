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
	got := permute([]int{1,2,3})
	t.Log(got)
	got = permute([]int{0, 1})
	t.Log(got)
	got = permute([]int{1})
	t.Log(got)
	got = permute([]int{5,4,6,2})
	t.Log(got)
}
