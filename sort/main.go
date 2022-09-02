package main

import "fmt"

func main() {

	// var arr = []int{44, 11, 55, 22, 66, 33}
	// result := bubbleSort(arr)
	// fmt.Println(result)

	// var arr = []int{44, 11, 55, 22, 66, 33}
	// result := selectionSort(arr)
	// fmt.Println(result)

	var arr = []int{44, 11, 55, 22, 66, 33}
	result := insertionSort(arr)
	fmt.Println(result)
}