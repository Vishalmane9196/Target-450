package main

import "math"

func nthroot(n, m int) int {

	var l = 1
	var h = m

	for l <= h {
		mid := l+(l-h)/2
		temp := math.Pow(mid, n)
	}
}