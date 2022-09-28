package main

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
