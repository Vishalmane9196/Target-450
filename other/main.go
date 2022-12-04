package main

import "fmt"

func main() {

	// number := 9669
	// res := maximum69Number(number)
	// fmt.Println("maximum69Number : ", res)

	// ax1 := -3
	// ay1 := 0
	// ax2 := 3
	// ay2 := 4
	// bx1 := 0
	// by1 := -1
	// bx2 := 9
	// by2 := 2
	// res := computeArea(ax1, ay1, ax2, ay2, bx1, by1, bx2, by2)
	// fmt.Println("computeArea : ", res)

	// number := 14
	// res := isUgly(number)
	// fmt.Println("isUgly : ", res)

	// board := [][]byte{
	// 	{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
	// 	{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
	// 	{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
	// 	{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
	// 	{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
	// 	{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
	// 	{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
	// 	{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
	// 	{'.', '.', '.', '.', '8', '.', '.', '7', '9'}}
	// res := isValidSudoku(board)
	// fmt.Println("isValidSudoku : ", res)

	word := []int{2, 5, 3, 9, 5, 3}
	res := minimumAverageDifference(word)
	fmt.Println("minimumAverageDifference : ", res)
}
