package main

import "fmt"

func main() {

	//consecutive 3
	// var nums = []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	// res := findMaxConsecutiveOnes(nums)
	// fmt.Println("max consective 1  problem solution is :", res)

	//consecutive 3
	// var nums = []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	// var k = 2
	// res := longestOnes(nums, k)
	// fmt.Println("max consective 3  problem solution is :", res)

	//fruit into  basket
	// var nums = []int{1, 2, 3, 2, 2}
	// res := totalFruit(nums)
	// fmt.Println("Fruit Into Baskets is :", res)

	// var nums = []int{12, 1, 78, 90, 57, 89, 56}
	// res := maxSlidingWindow(nums, 3)
	// fmt.Println("maxSlidingWindow :", res)

	// var nums = []int{1, 2, 3, 4, 5}
	// var nums = []int{2, 1, 5, 0, 4, 6}
	// var nums = []int{2, 1, 5, 0, 4, 6}
	// res := increasingTriplet(nums)
	// fmt.Println("increasingTriplet :", res)

	// res := lengthOfLongestSubstring("pwwkew")
	// fmt.Println("lengthOfLongestSubstring :", res)

	// res := characterReplacement("ABAB", 2)
	// res := characterReplacement("AAAA", 0)
	// fmt.Println("characterReplacement :", res)

	nums := []int{1, 1, 2, 1, 1}
	target := 3
	res := numberOfSubarrays(nums, target)
	fmt.Println("numberOfSubarrays :", res)

}
