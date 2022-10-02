package main

import "math"

// func nthroot(n, m int) int {

// 	var l = 1
// 	var h = m

// 	for l <= h {
// 		mid := l + (l-h)/2
// 		temp := math.Pow(mid, n)
// 	}
// }

//Link: https://www.geeksforgeeks.org/implementing-upper_bound-and-lower_bound-in-c/
func lowerBound(arr []int, target int) int {
	lo := 0
	hi := len(arr) - 1

	for hi-lo > 1 {
		mid := lo + (hi-lo)/2
		if arr[mid] < target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	if arr[lo] >= target {
		return lo
	}
	if arr[hi] >= target {
		return hi
	}
	return -1
}

func upperBound(arr []int, target int) int {
	lo := 0
	hi := len(arr) - 1

	for hi-lo > 1 {
		mid := lo + (hi-lo)/2
		if arr[mid] <= target {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	if arr[lo] > target {
		return lo
	}
	if arr[hi] > target {
		return hi
	}
	return -1
}

//Link: https://www.geeksforgeeks.org/search-insert-position-of-k-in-a-sorted-array/
func searchInsertPosition(arr []int, target int) int {

	lo := 0
	hi := len(arr) - 1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return hi + 1
}

//Link: https://practice.geeksforgeeks.org/problems/check-if-an-array-is-sorted0701/1?utm_source=youtube&utm_medium=collab_striver_ytdescription&utm_campaign=check-if-an-array-is-sorted
func checkSortedArray(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i+1] < arr[i] {
			return false
		}
	}
	return true
}

//Link: https://leetcode.com/problems/find-first-and-last-position-of-element-in-sorted-array/
func firstAndLastOccurance(nums []int, target int) []int {

	firstIndex := findFirstIndex(nums, target)
	lastIndex := findLastIndex(nums, target)

	return []int{firstIndex, lastIndex}
}

func findLastIndex(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	lastIndex := -1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			lastIndex = mid
			left = mid + 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return lastIndex
}

func findFirstIndex(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	firstIndex := -1

	for left <= right {
		mid := (left + right) / 2

		if nums[mid] == target {
			firstIndex = mid
			right = mid - 1
		} else if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return firstIndex
}

//Link: https://leetcode.com/problems/single-element-in-a-sorted-array/
func singleNonDuplicate(nums []int) int {

	lo := 0
	hi := len(nums) - 2

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == nums[mid^1] {
			lo = mid + 1
		} else {
			hi = mid - 1
		}
	}
	return nums[lo]
}

//Link: https://leetcode.com/problems/search-a-2d-matrix/
func search2DMatrix(matrix [][]int, target int) bool {
	var m = len(matrix)
	var n = len(matrix[0])
	var left = 0
	var right = m*n - 1

	if m == 0 {
		return false
	}

	for left <= right {

		var middle = left + ((right - left) / 2)
		i, j := middle/n, middle%n

		if matrix[i][j] == target {
			return true
		}
		if matrix[i][j] < target {
			left = middle + 1
		}
		if matrix[i][j] > target {
			right = middle - 1
		}
	}
	return false
}

//Link: https://leetcode.com/problems/find-a-peak-element-ii/description/
func findPeakGrid(mat [][]int) []int {
	startRow := 0
	endRow := len(mat) - 1
	res := make([]int, 0)

	for endRow >= startRow {
		middleRow := startRow + (endRow-startRow)/2
		// fmt.Printf(" startRow  :  %v || middleRow :  %v || endRow : %v\n", startRow, middleRow, endRow)
		// will get max element position for that row
		rowmax := maxRowElementPosition(mat[middleRow], len(mat[middleRow])-1)

		// middle row is the first row
		if middleRow == 0 {
			if mat[middleRow][rowmax] > mat[middleRow+1][rowmax] {
				res = append(res, middleRow, rowmax)
				return res
			}
		}

		//middle row is the last row
		if middleRow == len(mat)-1 {
			if mat[middleRow][rowmax] > mat[middleRow-1][rowmax] {
				res = append(res, middleRow, rowmax)
				return res
			}
		}
		//  checking max element of the row with it's upper and lower row
		if mat[middleRow][rowmax] > mat[middleRow+1][rowmax] && mat[middleRow][rowmax] > mat[middleRow-1][rowmax] {
			res = append(res, middleRow, rowmax)
			return res
		}
		// if max is lesser than next rows same column element, will move startRow to the nextRow
		if mat[middleRow][rowmax] < mat[middleRow+1][rowmax] {
			startRow = middleRow + 1
		} else { // otherwise move the endRow to current row -1
			endRow = middleRow - 1
		}

	}
	// we didn't find peak element in matrix
	res = append(res, -1, -1)
	return res
}

func maxRowElementPosition(arr []int, end int) int {
	max := 0

	for i := 0; i <= end; i++ {
		if arr[i] > arr[max] {
			max = i
		}
	}
	return max
}

//Link: https://leetcode.com/problems/koko-eating-bananas/
func minEatingSpeed(piles []int, h int) int {

	low := 1
	high := maxi(piles)

	ans := 0
	for low <= high {
		mid := low + (high-low)/2
		time := timefunc(mid, piles)

		if time <= h {
			ans = mid
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return ans
}

func timefunc(mid int, arr []int) int {
	sum := 0
	for i := 0; i < len(arr); i++ {
		sum += int(math.Ceil(float64(arr[i]) / float64(mid)))
	}
	return sum
}

func maxi(arr []int) int {
	max := arr[0]
	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	return max
}
