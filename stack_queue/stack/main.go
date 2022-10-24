package main

import "fmt"

func main() {

	/*****stack implementation******/
	var s = &Stack{}
	fmt.Println("isEmpty  : ", isEmpty(s))
	push(s, 5)
	push(s, 4)
	push(s, 3)
	push(s, 2)
	push(s, 1)
	printStack(s)
	// pop(s)
	// fmt.Println("After pop()")
	// printStack(s)
	// fmt.Println("Top elment : ", top(s))

	//implement 2 stack in one array
	// var arr [5]int
	// var s1 = &Stack{}
	// var s2 = &Stack{}
	// var top1 int = -1
	// var top2 int = len(arr)
	// push1(&arr, s1, 2, &top1, &top2)
	// push1(&arr, s1, 3, &top1, &top2)
	// push2(&arr, s2, 5, &top1, &top2)
	// push2(&arr, s2, 4, &top1, &top2)
	// fmt.Println("arr : ", arr)
	// pop1(&arr, s1, 3, &top1, &top2)
	// pop2(&arr, s2, 4, &top1, &top2)
	// fmt.Println("arr : ", arr)

	/*******Problem********/
	// res := isBalancedParenthesis("()[{}()]")
	// res := isBalancedParenthesis("[()")
	// fmt.Println("is Balanced parenthesis : ", res)

	// ans := ""
	// reverseString("love", &ans, 0)
	// fmt.Println("reverseString : ", ans)

	// deleteMiddle(s, s.getLength())
	// fmt.Println("after deleteing middle element")
	// printStack(s)

	// pushBottom(s, 9)
	// fmt.Println("after pushing new elment to bootom of stack")
	// printStack(s)

	reverseStack(s)
	print(s)

	//need to implement it again not working
	// sortStack(s)
	// print(s)

	//redundant bracket
	// isRedundantBraces := isRedundantBraces("((b*c))")
	// fmt.Println("isRedundantBraces : ", isRedundantBraces)

	//find minimum cost
	// findmin := findMinCost("{{{}")
	// fmt.Println(" findmin   :  ", findmin)

	//next smaller element
	// arr := []int{2, 1, 4, 3}
	// nextSmaller(arr)

	//largest rectangular area on histogram
	// 	arr := []int{2,1,5,6,2,3}
	// 	res:= largestRectangleArea(arr)
	//   fmt.Println("res  : ",  res)

	// mstackusingQueue()

	// minStack()

	// res := infixToPostfix("(p+q)*(m-n)")
	// fmt.Println("InfixToPostfix   :  ", res)

	// res := prefixToInfix("*-A/BC-/AKL")
	// fmt.Println("prefixToInfix   :  ", res)

	// res := prefixToPostfix("*+AB-CD")
	// fmt.Println("prefixToPostfix   :  ", res)

	// res := postfixToPrefix("ABC/-AK/L-*")
	// fmt.Println("postfixToPrefix   :  ", res)

	// res := postfixToInfix("ab*c+")
	// fmt.Println("postfixToInfix   :  ", res)

	// res := infixToPrefix("((a/b)+c)-(d+(e*f))")
	// fmt.Println("infixToPrefix   :  ", res)

	// nums := []int{5, 7, 1, 7, 6, 0}
	// res := nextGreaterElements2(nums)
	// fmt.Println("nextGreaterElements   :  ", res)

	// nums1 := []int{4, 1, 2}
	// nums2 := []int{1, 3, 4, 2}
	// res := nextGreaterElement1(nums1, nums2)
	// fmt.Println("nextGreaterElements   :  ", res)

	// nums := []int{3, 1, 2, 4, 0, 1, 3, 2}
	// nums := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	// res := trap(nums)
	// fmt.Println("trap water   :  ", res)

	// nums := []int{5, 7, 1, 7, 6, 0}
	// res := nextSmallerElements2(nums)
	// fmt.Println("nextSmallerElements2   :  ", res)

	// nums := []int{3, 4, 2, 7, 5, 8, 10, 6}
	// res := nextGreaterElementCount(nums)
	// fmt.Println("nextGreaterElementCount   :  ", res)
}
