package main

//final result of bubble sort will be after completeing one outer loop the biggest element will come to right hand side
func bubbleSort(arr []int) []int {

	//outer loop will iterate for len(arr)-1 times
	for i := 1; i < len(arr); i++ {
		//Inner loop will start from 0 till len(arr)-2
		//will check the jth and j+1th element
		//jth > j+1th swap both element 
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
	return arr
}

func selectionSort(arr []int) []int {

	//outer will iterarate for len(arr)-1 times
	for i := 0; i < len(arr)-1; i++ {
		//Inner loop will start iterataing from i+1 position 
		//check jth < ith element then swap it 
		for j := i+1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				// fmt.Printf("jth %v  : ith %v \n", arr[j], arr[i])
				arr[j], arr[i] = arr[i], arr[j]
			}
		}
	}
	return arr
}

func insertionSort(arr []int) []int{
	
	//one element array is always sorted so we will start from 1 index
	for i:=1; i<len(arr); i++ {
		//store the value of ith element
		//assign one variable to previous index   
		currElement := arr[i]
		j := i-1
		
		// check weather that element at jth index is geater than current element
		//if it geater then move element to right
		for j > -1 && arr[j] > currElement {
			arr[j+1] = arr[j]
			j=j-1
		}
		//there will be one vacant place at j+1 position to be filled up
		arr[j+1] = currElement
	}
	return arr
}