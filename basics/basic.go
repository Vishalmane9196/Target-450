package main

import (
	"fmt"
	"math"
	"sort"
	"unicode"
)

/***************Hashing Problems************/\
// Link:https://takeuforward.org/data-structure/count-frequency-of-each-element-in-the-array/
func Frequency(arr []int) map[int]int {
	fmt.Println("arr : ", arr)

	//approach 1
	// var freMap = map[int]int{}
	// for i := 0; i < len(arr); i++ {
	// 	freMap[arr[i]]++
	// }
	// return freMap

	//approach 2
	var freMap = map[int]int{}
	for i := 0; i < len(arr); i++ {
		if freMap[arr[i]] == 0 {
			freMap[arr[i]] = 1
		} else {
			freMap[arr[i]]++
		}
	}
	return freMap
}

// Link:https://takeuforward.org/data-structure/count-frequency-of-each-element-in-the-array/
func FrequencyOfCharacters(str string) map[string]int {
	fmt.Println("str : ", str)

	//approach 1
	var freMap = map[string]int{}
	for i := 0; i < len(str); i++ {
		if freMap[string(str[i])] == 0 {
			freMap[string(str[i])] = 1
		} else {
			freMap[string(str[i])]++
		}
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

/*********** Recurssion Problems*************/
// question: Understand recursion by print something N times
func printNTimes(n int) {
	if n == 0 {
		return
	}
	word := "hello"
	fmt.Printf("%v ", word)

	printNTimes(n - 1)
}

// question: Print 1 to N using recursion
func print1ToN(index, n int) {
	if index == n {
		return
	}
	fmt.Printf("%v ", index+1)

	print1ToN(index+1, n)
}

// question: Print N to 1 using recursion
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

// question: Factorial of N numbers
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

// question: fibonacci number 
func fib(n int) int {

	if n <= 1 {
		return n
	}
	return fib(n-1) + fib(n-2)
}

/***********String Problems*************/
// Link: https://leetcode.com/problems/valid-palindrome/
func isPalindromeString(s string) bool {

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
// pangram sentence is the sentense which contains all the letters from english alphabet at least once
func checkIfPangram(sentence string) bool {

	n := len(sentence)

	if n < 26 {
		return false
	}

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

/**********Math Problems**********/
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

// Link:https://takeuforward.org/maths/check-if-a-number-is-armstrong-number-or-not/
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

// Link: https://takeuforward.org/data-structure/print-all-divisors-of-a-given-number/
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

// Link:https://takeuforward.org/data-structure/find-gcd-of-two-numbers/
func gcd(num1, num2 int) int {

	//approach 1
	// min := mini(num1, num2)
	// var ans int
	// for i := 1; i <= min; i++ {

	// 	if num1%i == 0 && num2%i == 0 {
	// 		ans = i
	// 	}
	// }
	// return ans

	//approach 2
	return gcdHelper(num1, num2)

}

func gcdHelper(a, b int) int {
	if b == 0 {
		return a
	}
	return gcdHelper(b, a%b)
}

// Link: https://takeuforward.org/data-structure/check-if-a-number-is-prime-or-not/
func isPrime(num int) bool {

	//approach 1
	// for i := 2; i < num; i++ {
	// 	if num%i == 0 {
	// 		return false
	// 	}
	// }
	// return true

	// approach 2
	for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
		fmt.Println("i : ", i)
		if num%i == 0 {
			return false
		}
	}
	return true
}
