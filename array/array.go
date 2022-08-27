package main

import "fmt"

func getLargest(arr []int) {

	//approach 1
	var max = arr[0]

	for _, v := range arr {
		if v > max {
			max = v
		}
	}
	fmt.Println("largest element  : ", max)

	//approach 2 
	//sort the array and return the last element of array

}

func isSorted(arr []int, n int) bool{

	//approach one
	// for i:=0; i<n; i++ {
	// 	for j:= i+1; j<n;j++ {
  //     if arr[j] < arr[i] {
	// 			return false
	// 		}
	// 	}
	// }
	// return true

	//approach 2
	for i:=1; i<n; i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

func removeDuplicates(arr []int) {

	//approach 1
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

	//approach 2
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

//Left rotate array by one

func solve(arr []int) []int{

	//approach one
	// var temp = make([]int, len(arr))
  // for i:=1; i < len(arr); i++ {
	// 	temp[i-1] = arr[i]
	// }

	// temp[len(arr)-1] = arr[0]

	// return temp

	//approach two

	var temp int = arr[0]

	for i:=1; i < len(arr); i++ {
		arr[i-1] = arr[i]
	}
	arr[len(arr)-1] = temp

	return arr
}

func rotateElement(arr []int, k int) []int{
	if len(arr) == 0 || len(arr) == 1 {
		return arr
	}

	var temp = make([]int, len(arr))

	for i:=0; i < len(arr); i++ {
		temp[(i+k)%len(arr)] = arr[i]
	}
	return temp 	
}

//move all zeros to end of array
func zeroToEnd(arr []int) []int{

	//approach one
	//create new temp array move all non-zero to it and fill other position to 0 till we reach length of arr
	// var temp = make([]int, len(arr))
  // var k int

	// for i:=0; i < len(arr); i++ {
	// 	if arr[i] != 0 {
	// 		temp[k] = arr[i]
	// 		k++
	// 	}
	// }
	// // fmt.Println("temp : ", temp)

	// for i:=k ; i < len(arr); i++ {
	// 	temp[k]=0
	// 	k++
	// }
	// return temp

	//approach two
	//two pointer approach


	//finding the first ocuurance of zero
	var k int
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0{
			k=i
			break
		}
	} 

	var m = k
	var n = k+1

	for m < len(arr) && n < len(arr) {
		if arr[n] != 0 {
			arr[m],arr[n] = arr[n], arr[m]
			m++
		}
		n++
	}
	return arr
	
}

//linear search
func search(arr []int, searchNum int) int{

	for i,v:=range arr {
		if v == searchNum {
			return i
		}
	}
	return -1
}

func union(arr1, arr2 []int) []int{

//approach 1 
//inter over bith array find store the number and it's frequency in map and then append the all numbers to new array
	// var freq = map[int]int{}
	// var union = []int{}

	// for _,v:=range arr1{
	// 	freq[v]++
	// }

	// for _,v:=range arr2{
	// 	freq[v]++
	// }

	// fmt.Println("union : ", union)
	// for i,_:=range freq{
	// 	// fmt.Printf(" %v ===> %v \n",i,v)
	// 	union = append(union, i)
	// }
	// return union

	//approach 2
	//create set which can store both array

	var set = map[int]bool{}
  var union = []int{}

	for i,_:=range arr1{
		set[arr1[i]] = true
	}

	for i,_:=range arr2{
		set[arr2[i]] = true
	}
  
	for i,_:=range set{
		union = append(union, i)
	}
	return union

	//approach 3 
	//using two pointer 
}


func intersection(arr1, arr2 []int) []int{

//approach 1 
//inter over bith array find store the number and it's frequency in map and then append the all numbers to new array
	var freq = map[int]int{}
	var intersection = []int{}

	for _,v:=range arr1{
		freq[v]++
	}

	for _,v:=range arr2{
		freq[v]++
	}

	// fmt.Println("union : ", union)
	for i,v:=range freq{
		// fmt.Printf(" %v ===> %v \n",i,v)
		if v > 1 {
			intersection = append(intersection, i)
		}
	}
	return intersection

	//approach 2
	//using two pointer 
}


func longsubArr(arr []int, k int) int {
 

	//approach one
	//iterate over array and check ithe sum equal to k or not

	// var maxLength int
	// for i:=0; i< len(arr); i++ {
	// 	var sum  = 0
	// 	for j:=i; j < len(arr); j++ {
	// 		sum = sum + arr[j]
	// 		if sum == k {
	// 			maxLength = max(maxLength,j-i+1)
	// 		}
	// 		if sum > k {
	// 			break
	// 		}
	// 	}
	// }
	// return maxLength

	//approach two
	//sliding window technique

	var i, j = 0, -1
	var sum, maxlength int
	var n = len(arr)

	for i < n {
		for j+1 < n &&  sum+arr[j+1]<=k {
			j++
			sum = sum + arr[j]
		}
		if sum == k {
			maxlength = max(maxlength, j-i+1)
		}
		sum = sum - arr[i]
		i++
	}
	return maxlength
}

func max(x, y int) int {
	if x >= y {
		return x
	}else{
		return y
	}
}

func sort2DArray(mat [][]int, target int) bool {

	var m = len(mat)
	var n = len(mat[0])
	var lo = 0
	var hi = m*n-1

	if m == 0 {
		return false
	}

	for lo <= hi {
		var mid int = lo + ((hi-lo)/2)

		if mat[mid/n][mid%n] == target {
			return true
		}
		if mat[mid/n][mid%n] < target {
			lo = lo + 1
		}
		if mat[mid/n][mid%n] > target {
			hi = hi - 1
		}
	}
 return false
}

func missingNumber(arr []int) int {
var sum int
	for i:=0; i< len(arr); i++ {
		sum = sum + arr[i]
	}
	var totalSum = len(arr) * ((len(arr)+1)/2)
	fmt.Println("sum : ", sum)
	fmt.Println(" total sum : ", totalSum)

	missNo := totalSum - sum
	return missNo
}

func sort012s(arr []int) []int{

	var low int = 0
	var mid int = 0
	var high int = len(arr)-1

	for mid <= high {

		switch arr[mid] {
			case 0 :
				arr[low], arr[mid] = arr[mid], arr[low]
				low++
				mid++
			case 1 :
				mid++
			case 2 :
				arr[high], arr[mid] = arr[mid], arr[high]
				high--
		}
	}
	return arr 
}

func twoSum(arr []int, target int)[]int{
	var res []int
  var mp  = make(map[int]int)

	for i:=0; i< len(arr); i++ {
		
		num1 := target - arr[i]
	  num2, ok := mp[num1]
		if ok {

			res = append(res, i)
			res = append(res, num2)
			mp[arr[i]] = i
		}
		mp[arr[i]] = i
	}
	return res
}