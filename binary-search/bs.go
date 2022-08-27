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