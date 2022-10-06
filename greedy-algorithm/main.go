package main

import "fmt"

func main() {
	// n := 6
	// start := []int{1, 3, 0, 5, 8, 5}
	// end := []int{2, 4, 6, 7, 9, 9}
	// maxMeeting(start, end, n)

	// N := 4
	// K := 2
	// candies := []int{3, 2, 1, 4}
	// res := candyStore(candies, N, K)
	// fmt.Println("candyStore : ", res)

	nums := []int{3, 2, 1, 0, 4}
	res := canJump(nums)
	fmt.Println("canjump : ", res)

}
