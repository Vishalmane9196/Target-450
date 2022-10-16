package main

import (
	"fmt"
	"math"
)

// https://leetcode.com/problems/max-consecutive-ones/
func findMaxConsecutiveOnes(nums []int) int {

	count, max := 0, 0

	for _, v := range nums {
		if v == 0 {
			count = 0
		} else if v == 1 {
			count = count + 1
			if count > max {
				max = count
			}
		}
	}
	return max
}

// https://leetcode.com/problems/max-consecutive-ones-iii/
func longestOnes(nums []int, k int) int {
	zeroCount := 0
	i := 0
	result := 0

	for j := 0; j < len(nums); j++ {
		if nums[j] == 0 {
			zeroCount++
		}
		for zeroCount > k {
			if nums[i] == 0 {
				zeroCount--
			}
			i++
		}
		result = maxi(result, j-i+1)
	}
	return result
}

func maxi(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// https://leetcode.com/problems/fruit-into-baskets/
func totalFruit(fruits []int) int {

	var lastFruit = -1
	var secLastfruit = -1
	var countOfLastFruit = 0
	var currentMax = 0
	var maximum = 0

	for _, v := range fruits {

		//condition 1
		if v == lastFruit || v == secLastfruit {
			currentMax = currentMax + 1
		} else {
			currentMax = countOfLastFruit + 1
		}

		//condition 2
		if v == lastFruit {
			countOfLastFruit++
		} else {
			countOfLastFruit = 1
		}

		//condition 3
		if v != lastFruit {
			secLastfruit = lastFruit
			lastFruit = v
		}
		maximum = maxi(currentMax, maximum)
		// fmt.Printf("  %v || %v  || %v  || %v  || %v  \n", lastFruit, secLastfruit, currentMax, maximum, countOfLastFruit)
	}
	return maximum
}

// Link: https://leetcode.com/problems/sliding-window-maximum/
func maxSlidingWindow(nums []int, k int) []int {

	n := len(nums)
	dequeue := make([]int, 0) //to store elment in decreasing order
	ans := make([]int, 0)     //stire the ans

	i := 0
	for i = 0; i < k; i++ {
		//if len > 0 and last elment of deque is less than current elment of num
		//that means current elment is an poterntial candidate so popping the top element of dequeue
		for len(dequeue) > 0 && dequeue[len(dequeue)-1] < nums[i] {
			dequeue = dequeue[:len(dequeue)-1]
		}
		//append element to qequeue
		dequeue = append(dequeue, nums[i])
	}

	ans = append(ans, dequeue[0])
	for ; i < n; i++ {
		//check top elment of dequeue is in the range of current slidding window if not then remove that elment
		if dequeue[0] == nums[i-k] {
			dequeue = dequeue[1:]
		}
		for len(dequeue) > 0 && dequeue[len(dequeue)-1] < nums[i] {
			dequeue = dequeue[:len(dequeue)-1]
		}
		dequeue = append(dequeue, nums[i])
		ans = append(ans, dequeue[0])
	}
	return ans
}

// Link: https://leetcode.com/problems/sliding-window-maximum/
func increasingTriplet(nums []int) bool {

	// n := len(nums)
	// k := 3
	// dequeue := make([]int, 0) //to store elment in decreasing order
	// ans := true

	// i := 0
	// for i = 0; i < k; i++ {
	// 	//if len > 0 and last elment of deque is less than current elment of num
	// 	//that means current elment is an poterntial candidate so popping the top element of dequeue
	// 	if len(dequeue) > 0 && dequeue[len(dequeue)-1] > nums[i] {
	// 		ans = false
	// 	}
	// 	//append element to qequeue
	// 	dequeue = append(dequeue, nums[i])
	// }

	// for ; i < n; i++ {
	// 	//check top elment of dequeue is in the range of current slidding window if not then remove that elment
	// 	if dequeue[0] == nums[i-k] {
	// 		dequeue = dequeue[1:]
	// 	}
	// 	if len(dequeue) > 0 && dequeue[len(dequeue)-1] > nums[i] {
	// 		ans = false
	// 	} else {
	// 		ans = true
	// 	}
	// }
	// return ans

	//optimal approach
	n := len(nums)
	if n < 3 {
		return false
	}

	left := math.MaxInt
	mid := math.MaxInt

	for i := 0; i < n; i++ {

		if nums[i] > mid {
			return true
		} else if nums[i] > left && nums[i] < mid {
			mid = nums[i]
		} else if nums[i] < left {
			left = nums[i]
		}
	}
	return false
}

// Link: https://leetcode.com/problems/longest-substring-without-repeating-characters/
func lengthOfLongestSubstring(s string) int {

	set := map[byte]bool{}
	max := 0

	start := 0
	end := 0

	for start < len(s) {
		_, ok := set[s[start]]

		if !ok {
			if start-end+1 > max {
				max = start - end + 1
			}
			set[s[start]] = true
			start++
		} else {
			delete(set, s[end])
			end++
		}
	}
	return max
}

// remaining
// Link: https://leetcode.com/problems/longest-repeating-character-replacement/
func characterReplacement(s string, k int) int {

	set := map[byte]bool{}
	max := 0

	start := 0
	end := 0
	repeatCount := 0

	for start < len(s) {
		_, ok := set[s[start]]

		fmt.Println("max : ", max)
		if !ok {
			if start-end+1 > max {
				max = start - end + 1
			}
			set[s[start]] = true
			start++
		} else {
			if repeatCount < k {
				start++
				repeatCount++
				max = maxi(start-end+1, max)
			} else {
				delete(set, s[end])
				end++
				repeatCount--
			}
		}
	}
	return max
}

// Link: https://leetcode.com/problems/count-number-of-nice-subarrays/
func numberOfSubarrays(nums []int, k int) int {

	// var num, j, ans int = 0, 0, 0

	// for i := 0; i < len(nums); i++ {
	// 	for (num & nums[i]) != 0 {
	// 		//keep removing elment from back unless problem is resolved
	// 		//xor
	// 		num ^= nums[j]
	// 		j++
	// 	}
	// 	//no problem
	// 	//set bit
	// 	num |= nums[i]
	// 	ans = maxi(ans, i-j+1)
	// }
	// return ans

	n := len(nums)
	res := 0
	count := 0
	j := 0

	for i := 0; i < n; i++ {
		if nums[i]&1 == 0 {
			// dq = append(dq, i)
			k = k - 1
			count = 0
		}

		for k == 0 {
			k += nums[j] & 1
			j += 1
			count += 1
		}
		res += count
	}
	return res
}
