package main

import "fmt"

/*
left shift == multiply

right shift == divide
*/
func main() {

	// res, position := getBit(19, 3)
	// fmt.Printf("bit at %v position  is %v\n", position, res)

	// res, position := setBit(19, 3)
	// fmt.Printf("Binary is %b and bit is changed at %v ", res, position)

	// res, position := resetBit(19, 3)
	// fmt.Printf("Binary is %b and bit is changed at %v ", res, position)

	// oddOrEven(4)
	// swapNumbers(5, 9)

	// var nums = []int{4, 1, 2, 1, 2}
	// res := singleNumber(nums)
	// fmt.Println(" unique number  : ", res)

	res := powerOfTwo(64)
	fmt.Printf("Is number is power of 2 : %v\n", res)
}
