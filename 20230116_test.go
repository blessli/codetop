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
