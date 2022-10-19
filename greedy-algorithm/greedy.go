package main

import "sort"

//Link: https://takeuforward.org/data-structure/n-meetings-in-one-room/

// type pair struct {
// 	start int
// 	end   int
// }

// type pairs []pair

// func maxMeeting(start []int, end []int, n int) {
// 	var v []pair

// 	for i := 0; i < n; i++ {
// 		p := pair{start: start[i], end: end[i]}
// 		v = append(v, p)
// 	}
// 	fmt.Println(v)
// 	sort.IntSlice(v[0], v[len(v)-1])
// }

// func comparator(a, b pair) bool {
// 	return a.end < b.end
// }

// // sort
// func (p pairs) Len() int {
// 	return len(p)
// }

// func (p pairs) Swap(i, j int) {
// 	p[i], p[j] = p[j], p[i]
// }

// func (p pairs) Less(i, j pair) bool {
// 	return p[i] < p[j]
// }

// Link: https://practice.geeksforgeeks.org/problems/shop-in-candy-store1145/1
func candyStore(arr []int, n, k int) []int {
	sort.Ints(arr)

	var mini = 0
	var buy = 0
	var free = n - 1

	for buy <= free {
		mini += arr[buy]
		buy++
		free = free - k
	}

	var max = 0
	buy = n - 1
	free = 0

	for free <= buy {
		max += arr[buy]
		buy--
		free = free + k
	}
	var ans = []int{mini, max}
	return ans
}

//Link: https://practice.geeksforgeeks.org/problems/check-if-it-is-possible-to-survive-on-island4922/1

//Link: https://leetcode.com/problems/jump-game/
func canJump(nums []int) bool {

	n := len(nums)
	index := n - 1

	for i := n - 1; i >= 0; i-- {
		if nums[i]+i >= index {
			index = i
		}
	}
	return index == 0
}

//Link: https://leetcode.com/problems/jump-game-ii/
func jump(nums []int) int {

	var jump = 0
	var farthest = 0
	var current = 0
	n := len(nums)

	for i := 0; i < n-1; i++ {
		farthest = maxi(farthest, nums[i]+i)

		if i == current {
			current = farthest
			jump++
		}
	}
	return jump
}

func maxi(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func findContentChildren(greedArray []int, cookieSizeArray []int) int {

	sort.Ints(greedArray)
	sort.Ints(cookieSizeArray)

	greedIndex := 0
	sizeIndex := 0

	for greedIndex < len(greedArray) && sizeIndex < len(cookieSizeArray) {

		if cookieSizeArray[sizeIndex] >= greedArray[greedIndex] {
			greedIndex++
		}
		sizeIndex++
	}
	return greedIndex
}

//Link: https://takeuforward.org/data-structure/find-minimum-number-of-coins/
func minimumCoins(coinsArray []int, target int) int {

	n := len(coinsArray)
	count := 0
	for i := n - 1; i >= 0; i-- {
		for target >= coinsArray[i] {
			target -= coinsArray[i]
			count += 1
		}
		if target == 0 {
			break
		}
	}
	return count
}
