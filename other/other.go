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

//Link: https://leetcode.com/problems/rectangle-area/
func computeArea(ax1 int, ay1 int, ax2 int, ay2 int, bx1 int, by1 int, bx2 int, by2 int) int {

	first := area(ax1, ay1, ax2, ay2)
	second := area(bx1, by1, bx2, by2)

	leftOverlap := max(ax1, bx1)
	downOverlap := max(ay1, by1)

	rightOverlap := min(ax2, bx2)
	upOverlap := min(ay2, by2)

	overlap := area(leftOverlap, downOverlap, rightOverlap, upOverlap)

	return first + second - overlap

}

// Returns 0 if right < left or up < down
func area(left, down, right, up int) int {
	if left > right || down > up {
		return 0
	}
	return (right - left) * (up - down)
}

// math.Max accepts and returns floats and makes a mess
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
