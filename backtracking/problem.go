package main

import "fmt"

func main() {

	board := [][]byte{
		{'A', 'B', 'C', 'E'},
		{'S', 'F', 'C', 'S'},
		{'A', 'D', 'E', 'E'},
	}
	word := "ABCCED"
	res := exist(board, word)
	fmt.Println("exist : ", res)
}
