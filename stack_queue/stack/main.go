package main

import "fmt"

func main() {

	//stack Initialization and calling basic functions
	// var s = &Stack{}
	// isempty := isEmpty(s)
	// fmt.Println("isEmpty  : ", isempty)
	// push(s, 5)
	// push(s, 4)
	// push(s, 3)
	// push(s, 2)
	// push(s, 1)
	// fmt.Println("After push()")
	// print(s)
	// pop(s)
	// fmt.Println("After pop()")
	// print(s)
	// push(s,1)
	// topEle := top(s)
	// fmt.Println("Top elment : ",topEle )

	//implement stack
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

	//reverse stack
	//need to think about it
	// reverseStack("love")

	//delete middle elment of stack
	// deleteMiddle(s, s.getLength())
	// fmt.Println("after deleteing middle element")
	// print(s)

	//valid parenthesis
	// res := isBalancedParenthesis("( )[ { } ( ) ]")
	res := isBalancedParenthesis("[ ( )")
	fmt.Println("is Balanced parenthesis : ", res)

	//inser element at bottom of stack
	// pushBottom(s,9)
	// print(s)

	//reverse stak using recurssion
	// reverseStack(s)
	// print(s)

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

}
