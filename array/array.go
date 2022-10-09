package main

import (
	"fmt"
	"math"
	"sort"
)

// LinkL : https://takeuforward.org/data-structure/find-the-largest-element-in-an-array/
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

// Link : https://takeuforward.org/data-structure/check-if-an-array-is-sorted/
func isSorted(arr []int, n int) bool {

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
	for i := 1; i < n; i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

// Link: https://takeuforward.org/data-structure/remove-duplicates-in-place-from-sorted-array/
func removeDuplicates(arr []int) ([]int, int) {

	//approach 1
	// var set = make(map[int]bool)
	// newArr := make([]int, 0)

	// for i := 0; i < len(arr); i++ {
	// 	set[arr[i]] = true
	// }

	// for v := range set {
	// 	newArr = append(newArr, v)
	// }
	// return newArr, len(newArr)

	//approach 2
	var i int = 0
	for j := 1; j < len(arr); j++ {
		if arr[i] != arr[j] {
			i = i + 1
			arr[i] = arr[j]
		}
	}
	return arr, i + 1
}

// Link: https://takeuforward.org/data-structure/left-rotate-the-array-by-one/
func leftRotateByOne(arr []int) []int {
	fmt.Println("original array", arr)

	//approach one
	//by creating new slice
	// temp := make([]int, len(arr))
	// for i := 1; i < len(arr); i++ {
	// 	//move given array's elment to new array
	// 	temp[i-1] = arr[i]
	// }

	// temp[len(arr)-1] = arr[0]
	// return temp

	//approach two
	// Modify original array
	var temp int = arr[0]

	for i := 1; i < len(arr); i++ {
		arr[i-1] = arr[i]
	}
	arr[len(arr)-1] = temp
	return arr
}

// Link: https://takeuforward.org/data-structure/rotate-array-by-k-elements/
func rotateElementToRightByKPosition(arr []int, k int) []int {

	fmt.Printf("original array : %v  Position : %v\n", arr, k)
	//approach one
	// if len(arr) == 0 || len(arr) == 1 {
	// 	return arr
	// }
	// var temp = make([]int, len(arr))

	// for i := 0; i < len(arr); i++ {
	// 	temp[(i+k)%len(arr)] = arr[i]
	// }
	// return temp

	//approach 2
	//reverse 0 till n-k-1 and n-k to n-1 and 0 till n-1
	n := len(arr)
	k = k % n
	reverse(&arr, 0, n-k-1)
	reverse(&arr, n-k, n-1)
	reverse(&arr, 0, n-1)

	return arr
}

// Link: https://takeuforward.org/data-structure/rotate-array-by-k-elements/
func rotateElementToLeftByKPosition(arr []int, k int) []int {

	fmt.Printf("original array : %v  Position : %v\n", arr, k)

	//approach 1
	//reverse 0 till k-1 and k to n-1 and 0 till n-1
	n := len(arr)
	k = k % n
	reverse(&arr, 0, k-1)
	reverse(&arr, k, n-1)
	reverse(&arr, 0, n-1)

	return arr
}

// helper function to rotateElementToRightByKPosition and rotateElementToLeftByKPosition
func reverse(arr *[]int, low, high int) {
	for low <= high {
		(*arr)[low], (*arr)[high] = (*arr)[high], (*arr)[low]
		low++
		high--
	}
	fmt.Printf("arr : %v\n", arr)
}

// Link: https://takeuforward.org/data-structure/move-all-zeros-to-the-end-of-the-array/
func shiftZeroToEnd(arr []int) []int {
	fmt.Println("original array : ", arr)

	//approach one
	//create new temp slice of int move all non-zero to it and fill other position to 0 till we reach length of arr
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
		if arr[i] == 0 {
			k = i
			break
		}
	}

	var i = k
	var j = k + 1
	//finding non-zero elment and swap it with i which is first occurance of zero
	for i < len(arr) && j < len(arr) {
		if arr[j] != 0 {
			arr[i], arr[j] = arr[j], arr[i]
			i++
		}
		j++
	}
	return arr
}

// Link: https://takeuforward.org/data-structure/linear-search-in-c/
func search(arr []int, searchNum int) int {

	for i, v := range arr {
		if v == searchNum {
			return i
		}
	}
	return -1
}

// Link: https://takeuforward.org/data-structure/union-of-two-sorted-arrays/
func union(arr1, arr2 []int) []int {

	//approach 1
	//interate over both array find store the number and it's frequency in map and then append the all numbers to new array
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

	for i, _ := range arr1 {
		set[arr1[i]] = true
	}

	for i, _ := range arr2 {
		set[arr2[i]] = true
	}

	for i, _ := range set {
		union = append(union, i)
	}
	return union

	//approach 3
	//using two pointer
}

// Link: https://takeuforward.org/data-structure/intersection-of-two-sorted-arrays/
func intersection(arr1, arr2 []int) []int {

	//approach 1
	//inter over bith array find store the number and it's frequency in map and then append the all numbers to new array
	var freq = map[int]int{}
	var intersection = []int{}

	for _, v := range arr1 {
		freq[v]++
	}

	for _, v := range arr2 {
		freq[v]++
	}

	// fmt.Println("union : ", union)
	for i, v := range freq {
		// fmt.Printf(" %v ===> %v \n",i,v)
		if v > 1 {
			intersection = append(intersection, i)
		}
	}
	return intersection

	//approach 2
	//using two pointer
}

// Link: https://takeuforward.org/data-structure/longest-subarray-with-given-sum-k/
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
		for j+1 < n && sum+arr[j+1] <= k {
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
	} else {
		return y
	}
}

// Link: https://takeuforward.org/data-structure/search-in-a-sorted-2d-matrix/
func search2DArray(mat [][]int, target int) bool {

	var m = len(mat)
	var n = len(mat[0])
	var lo = 0
	var hi = m*n - 1

	if m == 0 {
		return false
	}

	for lo <= hi {
		var mid int = lo + ((hi - lo) / 2)

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

// Link: https://www.geeksforgeeks.org/find-the-missing-number/
func missingNumber(arr []int) int {
	var sum int
	for i := 0; i < len(arr); i++ {
		sum = sum + arr[i]
	}
	var totalSum = len(arr) * ((len(arr) + 1) / 2)
	fmt.Println("sum : ", sum)
	fmt.Println(" total sum : ", totalSum)

	missNo := totalSum - sum
	return missNo
}

// Link: https://takeuforward.org/data-structure/sort-an-array-of-0s-1s-and-2s/
func sort012s(arr []int) []int {

	var low int = 0
	var mid int = 0
	var high int = len(arr) - 1

	for mid <= high {

		switch arr[mid] {
		case 0:
			//swap low and mid
			arr[low], arr[mid] = arr[mid], arr[low]
			low++
			mid++
		case 1:
			mid++
		case 2:
			//swap mid and high
			arr[high], arr[mid] = arr[mid], arr[high]
			high--
		}
	}
	return arr
}

// Link: https://takeuforward.org/data-structure/two-sum-check-if-a-pair-with-given-sum-exists-in-array/
func twoSum(arr []int, target int) []int {
	var res []int
	var mp = make(map[int]int)

	for i := 0; i < len(arr); i++ {

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

// Link: https://takeuforward.org/data-structure/kadanes-algorithm-maximum-subarray-sum-in-an-array/
func maxSum(nums []int) int {

	var maxTillNow = nums[0]
	var currSum = 0

	for i := 0; i < len(nums); i++ {
		currSum += nums[i]
		maxTillNow = max(currSum, maxTillNow)

		if currSum < 0 {
			currSum = 0
		}
	}
	return maxTillNow
}

// Link: https://takeuforward.org/data-structure/minimum-in-rotated-sorted-array/
func findMin(nums []int) int {

	low, mid := 0, 0
	high := len(nums) - 1

	for low < high {
		mid = low + (high-low)/2

		if nums[mid] > nums[high] {
			low = mid + 1
		} else {
			high = mid
		}
	}
	return nums[low]
}

// Link: https://takeuforward.org/data-structure/spiral-traversal-of-matrix/
func spiralTraversalPfMatrix(arr [][]int) []int {
	var left = 0
	var right = len(arr[0]) - 1
	var top = 0
	var bottom = len(arr) - 1
	var res = []int{}

	for left <= right && top <= bottom {

		//print top row line
		for i := left; i <= right; i++ {
			res = append(res, arr[top][i])
			// fmt.Printf("%v ", arr[top][i])
		}
		top++

		//print rightmost coloum
		for i := top; i <= bottom; i++ {
			res = append(res, arr[i][right])
			// fmt.Printf("%v ", arr[i][right])
		}
		right--

		if top <= bottom {
			//print bottom row
			for i := right; i >= left; i-- {
				res = append(res, arr[bottom][i])
				// fmt.Printf("%v ", arr[bottom][i])
			}
			bottom--
		}

		if left <= right {
			//print leftmost column
			for i := bottom; i >= top; i-- {
				res = append(res, arr[i][left])
				// fmt.Printf("%v ", arr[i][left])
			}
			left++
		}
	}
	return res
}

// Link: https://takeuforward.org/data-structure/rotate-image-by-90-degree/
func rotateMatrixBy90(arr [][]int) [][]int {

	printMatrix(arr)
	row := len(arr)
	column := len(arr[0])

	if row <= 0 {
		return arr
	}
	//transpose matrix
	for i := 0; i < row; i++ {
		for j := i + 1; j < column; j++ {
			temp := arr[i][j]
			arr[i][j] = arr[j][i]
			arr[j][i] = temp
		}
	}

	halfColumn := column / 2
	//reverse the columns of each row
	for i := 0; i <= row-1; i++ {
		for j := 0; j < halfColumn; j++ {
			temp := arr[i][j]
			arr[i][j] = arr[i][column-j-1]
			arr[i][column-j-1] = temp

		}
	}
	return arr
}

func printMatrix(arr [][]int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr[0]); j++ {
			fmt.Printf("%v ", arr[i][j])
		}
		fmt.Println()
	}
}

// Link: https://takeuforward.org/data-structure/stock-buy-and-sell/
func maxProfit(prices []int) int {

	var minSoFar = prices[0]
	var maxProfit = 0

	for i := 0; i < len(prices); i++ {
		minSoFar = mini(minSoFar, prices[i])
		profit := prices[i] - minSoFar
		maxProfit = maxi(profit, maxProfit)
	}
	return maxProfit
}

func mini(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func maxi(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Link: https://leetcode.com/problems/max-consecutive-ones/
func findMaxConsecutiveOnes(nums []int) int {

	count, max := 0, 0

	for _, v := range nums {
		if v == 0 {
			count = 0
		} else if v == 1 {
			count = count + 1
			if count > max {
				max = count
			}
		}
	}
	return max
}

// Link: https://practice.geeksforgeeks.org/problems/row-with-max-1s0023/1?utm_source=youtube&utm_medium=collab_striver_ytdescription&utm_campaign=row-with-max-1s
// Link: https://medium.com/enjoy-algorithm/find-the-row-with-the-maximum-number-of-1s-3193b568c78
func rowWithMax1s(arr [][]int) int {
	max1, max1Row, cur := 0, 0, 0

	for i := 0; i < len(arr); i++ {
		cur = 0
		for j := 0; j < len(arr[0]); j++ {
			if arr[i][j] == 1 {
				cur++
			}
		}
		if cur > max1 {
			max1Row = i
		}
	}
	return max1Row
}

// Link: https://leetcode.com/problems/single-number/
func singleNumber(nums []int) int {

	var ans int
	for _, v := range nums {
		ans = ans ^ v

	}
	return ans
}

// Link: https://takeuforward.org/data-structure/search-element-in-a-rotated-sorted-array/
// Link: https://leetcode.com/problems/search-in-rotated-sorted-array/
func searchInSortedRotatedArray1(arr []int, target int) int {
	lo := 0
	hi := len(arr) - 1

	for lo <= hi {
		mid := lo + (hi-lo)/2

		if arr[mid] == target {
			return mid
		}

		//check if lef half is sorted or not
		if arr[lo] <= arr[mid] {
			if arr[lo] <= target && arr[mid] >= target {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else {
			if arr[mid] <= target && arr[hi] >= target {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}

// Link: https://leetcode.com/problems/search-in-rotated-sorted-array-ii/
func searchInSortedRotatedArray2(nums []int, target int) bool {

	lo := 0
	hi := len(nums) - 1

	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return true
		}
		if nums[mid] < nums[hi] {
			if target >= nums[mid] && target <= nums[hi] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		} else if nums[mid] > nums[hi] {
			if target >= nums[lo] && target <= nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else {
			hi--
		}
	}
	return false
}

// Link: https://takeuforward.org/data-structure/longest-consecutive-sequence-in-an-array/
func longestConsecutive(arr []int) int {

	//brute force
	// if len(arr) == 0 {
	// 	return 0
	// }

	// sort.Ints(arr)
	// var ans = 1
	// var prev = arr[0]
	// var cur = 1

	// for i := 1; i < len(arr); i++ {
	// 	if arr[i] == prev+1 {
	// 		cur++
	// 	} else if arr[i] != prev {
	// 		cur = 1
	// 	}
	// 	prev = arr[i]
	// 	ans = maxi(ans, cur)
	// }
	// return ans

	//optmised solution
	set := make(map[int]bool)
	for _, v := range arr {
		set[v] = true
	}
	var longestStreak = 0

	for i := 0; i < len(arr); i++ {
		if _, exist := set[arr[i]-1]; !exist {
			currentNum := arr[i]
			currentStreak := 1

			_, exist1 := set[currentNum+1]
			for exist1 {
				currentNum += 1
				currentStreak += 1
				_, exist1 = set[currentNum+1]
			}
			longestStreak = maxi(longestStreak, currentStreak)
		}
	}
	return longestStreak
}

// Link: https://takeuforward.org/data-structure/set-matrix-zero/
func setZeros(arr [][]int) [][]int {

	//brute force
	// row := len(arr)
	// col := len(arr[0])

	// for i := 0; i < row; i++ {
	// 	for j := 0; j < col; j++ {

	// 		if arr[i][j] == 0 {
	// 			//make left row of current matrix element to -1
	// 			idx := i - 1
	// 			for idx >= 0 {
	// 				if arr[idx][j] != 0 {
	// 					arr[idx][j] = -1
	// 				}
	// 				idx--
	// 			}
	// 			//make right row of current matrix element to -1
	// 			idx = i + 1
	// 			for idx < row {
	// 				if arr[idx][j] != 0 {
	// 					arr[idx][j] = -1
	// 				}
	// 				idx++
	// 			}

	// 			//make upper column of current matrix element to -1
	// 			idx = j - 1
	// 			for idx >= 0 {
	// 				if arr[i][idx] != 0 {
	// 					arr[i][idx] = -1
	// 				}
	// 				idx--
	// 			}

	// 			//make lower column of current matrix element to -1
	// 			idx = j + 1
	// 			for idx < col {
	// 				if arr[i][idx] != 0 {
	// 					arr[i][idx] = -1
	// 				}
	// 				idx++
	// 			}
	// 		}

	// 	}
	// }

	// for i := 0; i < row; i++ {
	// 	for j := 0; j < col; j++ {
	// 		if arr[i][j] <= 0 {
	// 			arr[i][j] = 0
	// 		}
	// 	}
	// }
	// return arr

	//optimised solution
	// row := len(arr)
	// col := len(arr[0])
	// var rowArr = make([]int, row)
	// var colArr = make([]int, col)

	// for i, _ := range rowArr {
	// 	rowArr[i] = -1
	// }

	// for i, _ := range colArr {
	// 	colArr[i] = -1
	// }

	// for i := 0; i < row; i++ {
	// 	for j := 0; j < col; j++ {
	// 		if arr[i][j] == 0 {
	// 			rowArr[i] = 0
	// 			colArr[j] = 0
	// 		}
	// 	}
	// }

	// for i := 0; i < row; i++ {
	// 	for j := 0; j < col; j++ {
	// 		if rowArr[i] == 0 || colArr[j] == 0 {
	// 			arr[i][j] = 0
	// 		}
	// 	}
	// }
	// return arr

	//more optmised solution
	row := len(arr)
	col := len(arr[0])
	col0 := 1

	for i := 0; i < row; i++ {
		//check whetherthe 0 is present in first column
		if arr[i][0] == 0 {
			col0 = 0
		}
		for j := 0; j < col; j++ {
			if arr[i][j] == 0 {
				//updated the first column and first row index to zero
				arr[i][0] = 0
				arr[0][j] = 0
			}
		}
	}

	//trvaersersing to arr backwardly
	for i := row - 1; i >= 0; i-- {
		for j := col - 1; j >= 0; j-- {
			if arr[i][0] == 0 || arr[0][j] == 0 {
				arr[i][j] = 0
			}
		}
		if col0 == 0 {
			arr[i][0] = 0
		}
	}
	return arr
}

// Link: https://takeuforward.org/data-structure/leaders-in-an-array/
func printLeaders(arr []int) []int {

	//brute force
	// var leaderArr = make([]int, 0)

	// for i := 0; i < len(arr); i++ {
	// 	leader := true
	// 	for j := i + 1; j < len(arr); j++ {
	// 		if arr[j] > arr[i] {
	// 			leader = false
	// 			break
	// 		}
	// 	}
	// 	if leader {
	// 		leaderArr = append(leaderArr, arr[i])
	// 	}
	// }
	// return leaderArr

	//optimised solution
	var n = len(arr)
	var leaderArr = make([]int, 0)
	var max = arr[n-1]
	leaderArr = append(leaderArr, max)

	for i := n - 2; i >= 0; i-- {

		if arr[i] > max {
			leaderArr = append(leaderArr, arr[i])
			max = arr[i]
		}
	}
	return leaderArr
}

func pascalTraingle(numRows int) [][]int {

	base := make([][]int, numRows)

	for col := 0; col < numRows; col++ {
		base[col] = make([]int, col+1)
		base[col][0], base[col][col] = 1, 1
	}
	base[0][0] = 1

	for row := 2; row < numRows; row++ {
		for col := 1; col < len(base[row])-1; col++ {
			base[row][col] = base[row-1][col-1] + base[row-1][col]
		}
	}
	return base
}

// Link: https://leetcode.com/problems/3sum-closest/
func threeSumClosest(nums []int, target int) int {
	ans := 0
	diff := math.MaxInt
	// fmt.Println("diff : ", diff)
	sort.Ints(nums)
	// fmt.Println("nums : ", nums)
	for i := 0; i < len(nums); i++ {
		first := nums[i]
		start := i + 1
		end := len(nums) - 1

		for start < end {
			// fmt.Println("ans : ", ans)
			if (first + nums[start] + nums[end]) == target {
				return target
			} else if Abs(first+nums[start]+nums[end]-target) < diff {
				diff = Abs(first + nums[start] + nums[end] - target)
				ans = first + nums[start] + nums[end]
			}
			if first+nums[start]+nums[end] > target {
				end--
			} else {
				start++
			}
		}
	}
	return ans
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
