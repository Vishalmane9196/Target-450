package main

import "fmt"

func main() {
	/********Problems********/
	// stones := [][]int{{0, 0}, {0, 2}, {1, 1}, {2, 0}, {2, 2}}
	// res := removeStones(stones)
	// fmt.Println("res : ", res)

	// maze := [][]byte{{"+", "+", "+"}, {".", ".", "."}, {"+", "+", "+"}}
	maze := [][]byte{{'+', '+', '+'}, {'.', '.', '.'}, {'+', '+', '+'}}
	entrance := []int{1, 0}
	res := nearestExit(maze, entrance)
	fmt.Println("nearestExit : ", res)

}
