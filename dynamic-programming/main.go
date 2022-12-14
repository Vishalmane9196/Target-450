package main

import "fmt"

func main() {
	/********Problems********/
	// number := 12
	// res := numSquares(number)
	// fmt.Printf("perfect square of %v  : %v\n", number, res)

	// startTime := []int{1, 2, 3, 3}
	// endTime := []int{3, 4, 5, 6}
	// profit := []int{50, 10, 40, 70}
	// res := jobScheduling(startTime, endTime, profit)
	// fmt.Printf("jobScheduling  : %v\n", res)

	nums := []int{2, 7, 9, 3, 1}
	res := rob(nums)
	fmt.Printf("max amount that can be looted is : %v\n", res)

}
