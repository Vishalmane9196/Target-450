package main

//Link: https://leetcode.com/problems/word-search/
func exist(board [][]byte, word string) bool {

	rowLength := len(board)
	if rowLength == 0 {
		return false
	}
	colLength := len(board[0])
	if colLength == 0 {
		return false
	}

	for i := 0; i < rowLength; i++ {
		for j := 0; j < colLength; j++ {
			if existHelper(board, word, i, j, rowLength, colLength, 0) {
				return true
			}

		}
	}
	return false
}

func existHelper(board [][]byte, word string, i, j, row, col, k int) bool {
	if k >= len(word) {
		return true
	}

	if i < 0 || i >= row || j < 0 || j >= col || board[i][j] == '.' || word[k] != board[i][j] {
		return false
	}

	if len(word) == 1 && word[k] == board[i][j] {
		return true
	}

	board[i][j] = '.'
	temp := false

	x := []int{0, 0, -1, 1}
	y := []int{-1, 1, 0, 0}

	for index := 0; index < 4; index++ {
		temp = temp || existHelper(board, word, i+x[index], j+y[index], row, col, k+1)
	}

	board[i][j] = word[k]
	return temp
}
