package main

import "math"

//Link: https://leetcode.com/problems/perfect-squares/
func numSquares(n int) int {

	dp := make([]int, n+1)

	dp[0] = 0
	dp[1] = 1

	for i := 2; i <= n; i++ {

		min := math.MaxInt
		for j := 1; j*j <= i; j++ {
			rem := i - j*j
			if dp[rem] < min {
				min = dp[rem]
			}
		}
		dp[i] = min + 1
	}
	return dp[n]
}
