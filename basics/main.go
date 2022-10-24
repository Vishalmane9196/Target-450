package main

import "fmt"

func main() {

	/**************************************************/
	//Problems on hashing

	// arr := []int{10, 5, 10, 15, 10, 5}
	// res := Frequency(arr)
	// fmt.Println("frequency : ", res)

	str := "leetcode"
	res := FrequencyOfCharacters(str)
	fmt.Println("FrequencyOfCharacters : ", res)

	// arr := []int{1, 2, 3, 1, 5, 2, 3, 1}
	// res := lowestHighestFrequency(arr)
	// fmt.Printf(" max : %v  min: %v \n", res[0], res[1])

	// arr := []int{7, 8, 4, 5, 4, 1, 1, 7, 7, 2, 5}
	// res := getMaximumDiff(arr)
	// fmt.Printf(" max difference : %v\n", res)

	/**************************************************/
	//Problems on recurssion

	// printNTimes(5)
	// print1ToN(0, 5)
	// printNTo1(5, 5)

	// sum := 0
	// sumOfN(1, 5, &sum)
	// fmt.Printf(" sum : %v\n", sum)

	// ans := 0
	// sumOfN(1, 5, &ans)
	// fmt.Printf(" sum : %v\n", ans)

	// ans := 1
	// factorial(5, &ans)
	// fmt.Printf(" factorial : %v\n", ans)

	// arr := []int{5, 4, 3, 2, 1}
	// reverseArray(&arr, 0, len(arr)-1)
	// fmt.Printf(" reverseArray : %v\n", arr)

	// arr := []int{5, 4, 3, 2, 1}
	// reverseArray(&arr, 0, len(arr)-1)
	// fmt.Printf(" reverseArray : %v\n", arr)

	/********String Problems********/
	// str := "A man, a plan, a canal: Panama"
	// res := isPalindromeString(str)
	// fmt.Printf(" isPalindromeString : %v\n", res)

	// str := "thequickbrownfoxjumpsoverthelazydog"
	// str := "leetcode"
	// res := checkIfPangram(str)
	// fmt.Printf(" checkIfPangram : %v\n", res)

	/**********Math Problems**********/
	// res := countDigit(123)
	// fmt.Printf("countDigit : %v\n", res)

	// res := reverseNumber(123)
	// fmt.Printf("reverseNumber : %v\n", res)

	// res := isPalindromeNumber(121)
	// fmt.Printf("isPalindromeNumber : %v\n", res)

	// res := armstrongNumebr(153)
	// fmt.Printf("armstrongNumebr : %v\n", res)

	// res := printDivisors(36)
	// fmt.Printf("printDivisors : %v\n", res)

	// num1 := 4
	// num2 := 8
	// res := gcd(num1, num2)
	// fmt.Printf("gcd : %v\n", res)

	// num := 11
	// res := isPrime(num)
	// ans := num != 1 && res == true
	// fmt.Printf("isPrime : %v\n", ans)

}
