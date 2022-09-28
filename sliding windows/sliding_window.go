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
