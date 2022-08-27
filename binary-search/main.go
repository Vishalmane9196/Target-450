package main

import "fmt"

func main() {

	var arr = []int{2, 3, 7, 10, 13, 14, 17}
	var target = 14

	res := binarySearch(arr, target)
	fmt.Println("res : ", res)

}