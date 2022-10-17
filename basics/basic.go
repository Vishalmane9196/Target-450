package main

import (
	"fmt"
	"math"
	"sort"
	"unicode"
)

/**************************************************/
//Problems on hashing

// Link:https://takeuforward.org/data-structure/count-frequency-of-each-element-in-the-array/
func Frequency(arr []int) map[int]int {

	var freMap = map[int]int{}
	for i := 0; i < len(arr); i++ {
		freMap[arr[i]]++
	}
	return freMap
}

// Link:https://takeuforward.org/data-structure/count-frequency-of-each-element-in-the-array/
func lowestHighestFrequency(arr []int) []int {
	sort.Ints(arr)
	min, max, count := len(arr), 0, 1

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == arr[i+1] {
			count += 1
			continue
		} else {
			max = maxi(max, count)
			min = mini(min, count)
			count = 0
		}
	}
	temp := []int{max, min}
	return temp
}

func maxi(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func mini(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// Link: https://www.tutorialcup.com/interview/hashing/difference-between-highest-and-least-frequencies-in-an-array.html
func getMaximumDiff(arr []int) int {
	sort.Ints(arr)
	min, max, count := len(arr), 0, 0

	for i := 0; i < len(arr)-1; i++ {
		if arr[i] == arr[i+1] {
			count += 1
			continue
		} else {
			max = maxi(max, count)
			min = mini(min, count)
			count = 0
		}
	}
	return max - min
}

/**************************************************/
//Problems on recurssion

// Link: Understand recursion by print something N times
func printNTimes(n int) {
	if n == 0 {
		return
	}
	word := "hello"
	fmt.Printf("%v ", word)

	printNTimes(n - 1)
}

// Link: Print 1 to N using recursion
func print1ToN(index, n int) {
	if index == n {
		return
	}
	fmt.Printf("%v ", index+1)

	print1ToN(index+1, n)
}

// Link: printNTo1
func printNTo1(index, n int) {
	if index == 0 {
		return
	}
	fmt.Printf("%v ", index)

	printNTo1(index-1, n)
}

// Link: https://takeuforward.org/data-structure/sum-of-first-n-natural-numbers/
func sumOfN(index, n int, sum *int) {
	if index > n {
		return
	}

	*sum += index

	sumOfN(index+1, n, sum)
}

// Factorial of N numbers
func factorial(n int, ans *int) int {
	if n <= 1 {
		return 1
	}

	*ans = n * factorial(n-1, ans)
	return *ans
}

// Link: https://takeuforward.org/data-structure/reverse-a-given-array/
func reverseArray(arr *[]int, i, j int) {
	if i > j {
		return
	}

	(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	reverseArray(arr, i+1, j-1)
}

func fib(n int) int {

	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

// Link: https://leetcode.com/problems/valid-palindrome/
func isPalindrome(s string) bool {

	leftIndex, rightIndex := 0, len(s)-1

	for leftIndex < rightIndex {
		leftChar := rune(s[leftIndex])
		rightChar := rune(s[rightIndex])

		if !unicode.IsLetter(leftChar) && !unicode.IsDigit(leftChar) {
			leftIndex++
		} else if !unicode.IsLetter(rightChar) && !unicode.IsDigit(rightChar) {
			rightIndex--
		} else if unicode.ToLower(leftChar) == unicode.ToLower(rightChar) {
			leftIndex++
			rightIndex--
		} else {
			return false
		}
	}
	return true
}

// Link: https://leetcode.com/problems/check-if-the-sentence-is-pangram/
func checkIfPangram(sentence string) bool {

	charCount := [26]int{}
	for i, _ := range charCount {
		charCount[i] = 0
	}

	for i := 0; i < len(sentence); i++ {
		charCount[sentence[i]-97]++
	}

	for i := 0; i < len(charCount); i++ {
		if charCount[i] == 0 {
			return false
		}
	}
	return true
}

/**************************************************/
//Problems on Math
func countDigit(number int) int {
	count := 0
	for number != 0 {
		number = number / 10
		count += 1
	}
	return count
}

// Link:https://leetcode.com/problems/reverse-integer/
func reverseNumber(number int) int {
	ans := 0
	for number != 0 {
		lastDigit := number % 10
		ans = ans*10 + lastDigit
		number = number / 10
		if ans > math.MaxInt32 || ans < math.MinInt32 {
			return 0
		}
	}
	return ans
}

// Link: https://leetcode.com/problems/palindrome-number/
func isPalindromeNumber(x int) bool {

	if x < 0 {
		return false
	}

	reverseNumber := reverseNumber(x)
	return x == reverseNumber
}

func armstrongNumebr(n int) bool {
	originalno := n
	count := 0
	temp := n
	for temp != 0 {
		count++
		temp = temp / 10
	}
	sumofpower := 0
	for n != 0 {
		digit := n % 10
		sumofpower += int(math.Pow(float64(digit), float64(count)))
		n /= 10
	}
	return (sumofpower == originalno)
}

func printDivisors(n int) []int {

	//approach 1
	// ans := []int{}
	// for i := 1; i <= n; i++ {
	// 	if n%i == 0 {
	// 		ans = append(ans, i)
	// 	}
	// }
	// return ans

	//approach 2
	ans := []int{}
	for i := 1; i <= int(math.Sqrt(float64(n))); i++ {
		if n%i == 0 {
			ans = append(ans, i)
			if i != n/i {
				ans = append(ans, n/i)
			}
		}
	}
	return ans

}
