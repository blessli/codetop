package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_wordBreak(t *testing.T) {
	got := wordBreak("leetcode", []string{"leet", "code"})
	assert.Equal(t, true, got)
	got = wordBreak("applepenapple", []string{"apple", "pen"})
	assert.Equal(t, true, got)
	got = wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})
	assert.Equal(t, false, got)
}

func Test_exist(t *testing.T) {
	got := exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "SEE")
	assert.Equal(t, true, got)
	got = exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCB")
	assert.Equal(t, false, got)
	got = exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'C', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCCED")
	assert.Equal(t, true, got)
	got = exist([][]byte{{'A', 'B', 'C', 'E'}, {'S', 'F', 'E', 'S'}, {'A', 'D', 'E', 'E'}}, "ABCESEEEFS")
	assert.Equal(t, true, got)
	got = exist([][]byte{{'a', 'a'}}, "aaa")
	assert.Equal(t, false, got)
	got = exist([][]byte{{'A', 'A', 'A', 'A', 'A', 'A'}, {'A', 'A', 'A', 'A', 'A', 'A'}, {'A', 'A', 'A', 'A', 'A', 'A'}, {'A', 'A', 'A', 'A', 'A', 'A'}, {'A', 'A', 'A', 'A', 'A', 'B'}, {'A', 'A', 'A', 'A', 'B', 'A'}}, "AAAAAAAAAAAAABB")
	assert.Equal(t, false, got)
	got = exist([][]byte{{'C', 'A', 'A'}, {'A', 'A', 'A'}, {'B', 'C', 'D'}}, "AAB")
	assert.Equal(t, true, got)
}
