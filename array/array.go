package main

import "fmt"

func getLargest(arr []int) {

	//aproach 1
	var max = arr[0]

	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	fmt.Println("largest element  : ", max)

	//aproach 2 
	//sort the array and return the last element of array

}

func isSorted(arr []int, n int) bool{

	//aproach one
	// for i:=0; i<n; i++ {
	// 	for j:= i+1; j<n;j++ {
  //     if arr[j] < arr[i] {
	// 			return false
	// 		}
	// 	}
	// }
	// return true

	//aproach 2
	for i:=1; i<n; i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

func removeDuplicates(arr []int) {

	//aproach 1
	// var set = map[int]bool{}

	// for i:=0; i< len(arr); i++ {
	// 	set[arr[i]]=true
	// }
	// fmt.Println("set : ", set)

	// var k = len(set)
	// var newArr []int 

	// for i:=0; i < k ; i++ {
	// 	newArr = append(newArr, i)
	// }

	// fmt.Println(newArr)
	// fmt.Println("k : ", k)

	//aproach 2

	var i int = 0 
	for j:=1; j < len(arr); j++ {
		if arr[i] != arr[j] {
			i = i + 1
			arr[i] = arr[j]
		}
	}
	fmt.Println("length of non-duplicate array : ", i+1)
	fmt.Println("arr : ", arr)
}