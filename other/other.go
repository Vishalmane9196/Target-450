package main

//Link: https://leetcode.com/problems/maximum-69-number/
func maximum69Number(num int) int {
	var nums []int
	for num != 0 {
		nums = append(nums, num%10)
		num /= 10
	}

	used := false
	ans := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] == 6 && !used {
			nums[i] = 9
			used = true
		}
		ans = ans*10 + nums[i]
	}

	return ans
}
