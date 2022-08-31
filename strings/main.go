package main

import "fmt"

func main() {

	// input := "(()())(())"
	// res := removeOutermostParenthesis(input)
	// fmt.Println("res : ", res)

	// input := "the sky is blue"
	// res := reverseWords(input)
	// fmt.Println("res : ", res)

	// input := "34564"
	// res := largestOddNumber(input)
	// fmt.Println("res : ", res)

	// input :=  []string{"flower","flow","flight"}
	// res := longestCommonPrefix(input)
	// fmt.Println("res : ", res)

	// var s,t = "paper","title"
	// var s,t = "foo","bar"
	// res := isIsomorphic(s,t)
	// fmt.Println("res : ", res)


	var s,t = "abcde","cdeab"
	res := rotateString(s,t)
	fmt.Println("res : ", res)

}