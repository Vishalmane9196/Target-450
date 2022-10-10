package main

//https://leetcode.com/problems/max-consecutive-ones/
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

//https://leetcode.com/problems/max-consecutive-ones-iii/
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

//https://leetcode.com/problems/fruit-into-baskets/
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

//Link: https://leetcode.com/problems/sliding-window-maximum/
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
