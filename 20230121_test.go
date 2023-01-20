package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLRUCache_Get(t *testing.T) {
	obj := Constructor1(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	got := obj.Get(1)
	assert.Equal(t, 1, got)
	obj.Put(3, 3)
	got = obj.Get(2)
	assert.Equal(t, -1, got)
	obj.Put(4, 4)
	got = obj.Get(1)
	assert.Equal(t, -1, got)
	t.Log(obj.list)
	got = obj.Get(3)
	assert.Equal(t, 3, got)
	got = obj.Get(4)
	assert.Equal(t, 4, got)
}

func Test_subsets(t *testing.T) {
	got := subsets([]int{1, 2, 3})
	t.Log(got)
	got = subsets([]int{0})
	t.Log(got)
}

func Test_letterCombinations(t *testing.T) {
	got:=letterCombinations("23")
	t.Log(got)
}
