package main

import (
	"fmt"
	"strings"
)

func exist(board [][]byte, word string) bool {
	m := len(board)
	n := len(board[0])
	ans := false
	vis := make([][]bool, m)
	for i := 0; i < m; i++ {
		vis[i] = make([]bool, n)
	}
	count := len(word)
	dist := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	var dfs func(int, int, int)
	dfs = func(x, y, pos int) {
		if ans {
			return
		}
		if pos == count {
			ans = true
			return
		}
		for i := 0; i < 4; i++ {
			xx := x + dist[i][0]
			yy := y + dist[i][1]
			if xx >= 0 && xx < m && yy >= 0 && yy < n && pos < count && board[xx][yy] == word[pos] && !vis[xx][yy] {
				vis[xx][yy] = true
				dfs(xx, yy, pos+1)
				vis[xx][yy] = false
			}
		}
	}
	for i := 0; i < m && !ans; i++ {
		for j := 0; j < n && !ans; j++ {
			if board[i][j] == word[0] {
				vis[i][j] = true
				dfs(i, j, 1)
				vis[i][j] = false
			}
		}
	}
	return ans
}

func wordBreak(s string, wordDict []string) bool {
	count := len(wordDict)
	ans := false
	vis := map[string]bool{}
	var dfs func(string)
	dfs = func(curr string) {
		if ans {
			return
		}
		if curr == s {
			ans = true
			return
		}
		for i := 0; i < count; i++ {
			now := curr + wordDict[i]
			if len(now) <= len(s) && strings.HasPrefix(s, now) && !vis[now] {
				vis[now] = true
				dfs(now)
			}
		}
	}
	dfs("")
	return ans
}
