package main

import "fmt"

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

//final result of bubble sort will be after completeing one outer loop the biggest element will come to right hand side
func improvedBubbleSort(arr []int) []int {

	var swapped bool
	//outer loop will iterate for len(arr)-1 times
	for i := 1; i < len(arr); i++ {
		//Inner loop will start from 0 till len(arr)-2
		//will check the jth and j+1th element
		//jth > j+1th swap both element 
		for j := 0; j < len(arr)-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
				swapped = true
			}			
		}
		
		if swapped == false {
			fmt.Println("inside break")
			break
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


func mergeSortTechnique(a1, a2 []int, m, n int) []int {

	i, j, k := 0, 0, 0

	var b = make([]int, m+n)
	for {

		if a1[i] <= a2[j] {
			b[k] = a1[i]
			k++
			i++
			if i == m {
				for j < n {
					b[k] = a2[j]
					k++
					j++
				}
				break
			}
			} else {
			b[k] = a2[j]
			k++
			j++
			if j == n {
				for i < m {
					b[k] = a1[i]
					k++
					i++
				}
				break
			}
		}
	}
	return b
}

func mergeSort(items []int) []int {
    var num = len(items)
     
    if num == 1 {
        return items
    }
     
    middle := int(num / 2)
    var (
        left = make([]int, middle)
        right = make([]int, num-middle)
    )
    for i := 0; i < num; i++ {
        if i < middle {
            left[i] = items[i]
        } else {
            right[i-middle] = items[i]
        }
    }
    return merge(mergeSort(left), mergeSort(right))
}
 
func merge(left, right []int) (result []int) {
    result = make([]int, len(left) + len(right))
     
    i := 0
    for len(left) > 0 && len(right) > 0 {
        if left[0] < right[0] {
            result[i] = left[0]
            left = left[1:]
        } else {
            result[i] = right[0]
            right = right[1:]
        }
        i++
    }
     
    for j := 0; j < len(left); j++ {
        result[i] = left[j]
        i++
    }
    for j := 0; j < len(right); j++ {
        result[i] = right[j]
        i++
    }
    return
}


func bubbleSortRecurssion(arr []int, m int) {

	//base case
	if m == 1 {
		return
	}

	var isSwapped bool
	//processing
	for i:=0; i<m; i++ {
		if arr[i] > arr[i+1] {
			 arr[i], arr[i+1] = arr[i+1], arr[i]
			 isSwapped = true
		}
	}

	if !isSwapped {
		return
	}
	//recursive relation
	bubbleSortRecurssion(arr, m-1)
}

func insertionSortRecurssive(arr []int, n int) {

	//base case
	if n <=1 {
		return
	}
	
	//recurssive relation
	insertionSortRecurssive(arr,n-1)

	//processing
	key := arr[n-1]
	j:=n-2

	for j>-1 && arr[j] > key {
		arr[j+1]=arr[j]
		j--
	}
	arr[j+1]= key
}

func partition(arr []int, low, high int) ([]int, int) {
	pivot := arr[high]
	i := low
	for j := low; j < high; j++ {
		if arr[j] < pivot {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
	}
	arr[i], arr[high] = arr[high], arr[i]
	return arr, i
}

func quickSort(arr []int, low, high int) []int {
	if low < high {
		var p int
		arr, p = partition(arr, low, high)
		arr = quickSort(arr, low, p-1)
		arr = quickSort(arr, p+1, high)
	}
	return arr
}