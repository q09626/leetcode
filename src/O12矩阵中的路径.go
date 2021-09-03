package main

func exist(board [][]byte, word string) bool {
	var dfs func(i, j, index int) bool
	dfs = func(i, j, index int) bool {
		// 越界、不符，提前剪枝
		if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] != word[index] {
			return false
		}
		// 遍历完字符串，提前剪枝
		if index == len(word)-1 {
			return true
		}
		board[i][j] = ' '
		res := dfs(i+1, j, index+1) || dfs(i-1, j, index+1) || dfs(i, j+1, index+1) || dfs(i, j-1, index+1)
		board[i][j] = word[index]
		return res
	}

	// 遍历所有位置作为起点
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if dfs(i, j, 0) {
				return true
			}
		}
	}
	return false
}
