package main

import (
	"fmt"
)

// Check if the i-th bit is set or not
func getBit(number, position int) (int, int) {
	fmt.Printf("binary of %v is %b \n", number, number)
	res := number & (1 << (position - 1))
	return res, position
}

// set the i-th bit
func setBit(number, position int) (int, int) {
	fmt.Printf("binary of %v is %b \n", number, number)
	res := number | (1 << (position - 1))
	return res, position
}

// set the i-th bit
func resetBit(number, position int) (int, int) {
	fmt.Printf("binary of %v is %b \n", number, number)
	mask := (1 << (position - 1))
	res := number &^ mask
	return res, position
}

// Check if a number is odd or not
func oddOrEven(n int) {
	bitResult := n & 1
	if bitResult == 1 {
		fmt.Println("number is odd : ", bitResult)
	} else {
		fmt.Println("number is  even : ", bitResult)
	}
}

func swapNumbers(a, b int) {

	fmt.Printf(" a : %v     b : %v\n", a, b)
	a = a ^ b
	b = a ^ b
	a = a ^ b

	fmt.Printf(" a : %v     b : %v\n", a, b)
}

//remaining
// func divideNumber(divident, divisor int) {

// 	divident = math.Abs(math.Inf(divi))
// 	var quotient int
// 	for divident >= divisor {
// 		divident = divident - divisor
// 		quotient++
// 	}
// 	fmt.Println("quotient = ", quotient)
// }

func singleNumber(nums []int) int {

	var ans int
	for _, v := range nums {
		ans = ans ^ v

	}
	return ans
}

// Count the number of set bits
func countSetBit(number int) int {
	fmt.Printf("binary of %v is %b \n", number, number)
	var count int
	for number > 0 {
		if (number & 1) == 1 {
			count++
		}
		number = number >> 1
	}
	return count
}

// Check if a number is power of 2 or not
func powerOfTwo(number int) bool {
	fmt.Printf("binary of %v is %b \n", number, number)
	var count int
	for number > 0 {
		if (number & 1) == 1 {
			count++
		}
		number = number >> 1
	}

	if count == 1 {
		return true
	} else {
		return false
	}
}
