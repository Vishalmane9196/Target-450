package main

import "fmt"

func main() {

	// var arr = []int{2, 3, 7, 10, 13, 14, 17}
	// var target = 14

	// res := binarySearch(arr, target)
	// fmt.Println("res : ", res)
	// res := binarySearchRecurssive(arr, 0,len(arr),target)
	// fmt.Println("res : ", res)

	// var arr = []int{2,2,3,3,3,3,4}
	// var target = 3
	// res := countOccurances(arr, target)
	// fmt.Println("res : ", res)

	// var arr = []int{4, 6, 10, 12, 18, 20}
	// var target = 4
	// res := lowerBound(arr, target)
	// fmt.Printf("lower bound of %v is %v at index %v\n", target, arr[res], res)

	// var arr = []int{4, 6, 10, 12, 18, 20}
	// var target = 6
	// res := upperBound(arr, target)
	// fmt.Printf("lower bound of %v is %v at index %v\n", target, arr[res], res)

	// var arr = []int{4, 6, 10, 12, 18, 20}
	// var target = 7
	// res := searchInsertPosition(arr, target)
	// fmt.Printf("target %v is %v at index %v\n", target, arr[res], res)

	// var arr = []int{5, 7, 7, 8, 8, 10}
	// var target = 8
	// res := firstAndLastOccurance(arr, target)
	// fmt.Printf("the first and last occurance of %v is %v  \n", target, res)

	// var arr = [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	// var target = 8
	// res := search2DMatrix(arr, target)
	// fmt.Printf(" search2DMatrix  : %v", res)

	// var arr = [][]int{{10, 20, 15}, {21, 30, 14}, {7, 16, 32}}
	// res := findPeakGrid(arr)
	// fmt.Printf(" findPeakGrid  : %v\n", res)

	var arr = []int{30, 11, 23, 4, 20}
	hours := 5
	res := minEatingSpeed(arr, hours)
	fmt.Printf(" minEatingSpeed  : %v\n", res)

}
