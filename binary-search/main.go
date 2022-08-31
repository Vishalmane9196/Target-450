package main

import "fmt"

func main() {

	// var arr = []int{2, 3, 7, 10, 13, 14, 17}
	// var target = 14
	
	// res := binarySearch(arr, target)
	// fmt.Println("res : ", res)
	// res := binarySearchRecurssive(arr, 0,len(arr),target)
	// fmt.Println("res : ", res)
	
	var arr = []int{2,2,3,3,3,3,4}
	var target = 3
	res := countOccurances(arr, target)
	fmt.Println("res : ", res)

}