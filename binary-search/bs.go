package main

func binarySearch(arr []int, target int) int {
	var start = 0
	var end = len(arr) - 1

	for start < end {
		var mid = start + ((end - start) / 2)

		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return -1
}

func binarySearchRecurssive(arr []int, start, end, target int) int {

	//base condition
	if start > end {
		return -1
	}

	//processing
	var mid = start + ((end - start) / 2)

	if arr[mid] == target {
		return mid
	} else if arr[mid] < target {
		//recurssive relation
		return binarySearchRecurssive(arr, mid+1, end, target)
	} else {
		//recurssive relation
		return binarySearchRecurssive(arr, start, mid-1, target)
	}
}

// Link: https://leetcode.com/problems/binary-search/
func getbsIdx(arr []int, target int) int {
	//find the occurnace of target and return its index
	var s = 0
	var e = len(arr) - 1

	for s <= e {

		var m = s + ((e - s) / 2)

		if arr[m] == target {
			return m
		} else if arr[m] < target {
			s = m + 1
		} else {
			e = m - 1
		}
	}
	return -1
}

func countOccurances(arr []int, target int) int {

	idx := getbsIdx(arr, target)
	if idx == -1 {
		return -1
	}

	var k int = 1
	left := idx - 1

	for left > 0 && arr[left] == target {
		k++
		left--
	}
	right := idx + 1
	for right < len(arr) && arr[right] == target {
		k++
		right++
	}
	return k
}
