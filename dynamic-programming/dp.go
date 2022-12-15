package main

import (
	"math"
	"sort"
)

// Link: https://leetcode.com/problems/perfect-squares/
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

/*
Link: https://leetcode.com/problems/maximum-profit-in-job-scheduling/
Link: https://www.youtube.com/watch?v=rZLvA1rsLy4&ab_channel=TanishqChaudhary
*/
type t struct {
	s int
	e int
	p int
}

func jobScheduling(st []int, et []int, p []int) int {
	r := make([]t, len(st))
	for i := range st {
		r[i] = t{s: st[i], e: et[i], p: p[i]}
	}

	sort.Slice(r, func(i, j int) bool { return r[i].s <= r[j].s })

	var dp func(int) int
	dp = func(i int) int {
		if i >= len(r) {
			return 0
		}

		a := r[i].p
		for j := i + 1; j < len(r); j++ {
			if r[i].e <= r[j].s {
				a += dp(j)
				break
			}
		}
		na := dp(i + 1)
		if na > a {
			return na
		}
		return a
	}
	dp = c(dp)
	return dp(0)
}

func c(f func(int) int) func(int) int {
	var m = make(map[int]int)
	d := func(arg int) int {
		if val, ok := m[arg]; ok {
			return val
		}
		val := f(arg)
		m[arg] = val
		return val
	}
	return d
}

// Link: https://leetcode.com/problems/house-robber/
func rob(nums []int) int {
	rob1, rob2 := 0, 0
	for i := 0; i < len(nums); i++ {
		temp := maxi(nums[i]+rob1, rob2)
		rob1 = rob2
		rob2 = temp
	}
	return rob2
}

func maxi(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Link: https://leetcode.com/problems/longest-common-subsequence/
func longestCommonSubsequence(text1 string, text2 string) int {
	l1 := len(text1)
	l2 := len(text2)
	dp := make([][]int, l1+1)
	for i, _ := range dp {
		dp[i] = make([]int, l2+1)
	}
	for i := 0; i < l1; i++ {
		for j := 0; j < l2; j++ {
			if text1[i] == text2[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				dp[i+1][j+1] = maxi(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	return dp[l1][l2]
}
