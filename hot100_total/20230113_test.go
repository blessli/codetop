package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_search(t *testing.T) {
	got := search([]int{-1, 0, 3, 5, 9, 12}, 9)
	assert.Equal(t, 4, got)
}

func Test_singleNumber(t *testing.T) {
	got := singleNumber([]int{4,1,2,1,2})
	assert.Equal(t, 4, got)
}
