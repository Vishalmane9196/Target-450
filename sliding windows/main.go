package main

import "fmt"

func main() {

	//consecutive 3
	// var nums = []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	// res := findMaxConsecutiveOnes(nums)
	// fmt.Println("max consective 1  problem solution is :", res)

	//consecutive 3
	// var nums = []int{1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0}
	// var k = 2
	// res := longestOnes(nums, k)
	// fmt.Println("max consective 3  problem solution is :", res)

	//fruit into  basket
	var nums = []int{1, 2, 3, 2, 2}
	res := totalFruit(nums)
	fmt.Println("Fruit Into Baskets is :", res)

}
