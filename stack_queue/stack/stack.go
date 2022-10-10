package main

import (
	"fmt"
	"math"
)

// Link: https://takeuforward.org/data-structure/implement-queue-using-array/
type Stack struct {
	arr []int
}

func push(s *Stack, val int) {
	s.arr = append(s.arr, val)
}

func isEmpty(s *Stack) bool {
	return len(s.arr) == 0
}

func (s *Stack) getLength() int {
	return len(s.arr)
}

func print(s *Stack) {
	fmt.Printf("[Bottom] <-> ")
	for i := 0; i < len(s.arr); i++ {
		fmt.Printf("[%v] <-> ", s.arr[i])
	}
	fmt.Printf("[Top]\n")
}

func pop(s *Stack) {
	stackLength := len(s.arr) - 1
	s.arr = s.arr[:stackLength]
}

func top(s *Stack) int {
	top := s.arr[len(s.arr)-1]
	return top
}

// Link: https://takeuforward.org/data-structure/check-for-balanced-parentheses/
func isBalancedParenthesis(str string) bool {

	var stack []rune
	for _, v := range str {
		switch v {
		case '(', '[', '{':
			stack = append(stack, v)
		case ')':
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case ']':
			if len(stack) > 0 && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		case '}':
			if len(stack) > 0 && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				return false
			}
		}
	}
	return len(stack) == 0
}

// implement 2 stack using array
func push1(arr *[5]int, s *Stack, val int, top1, top2 *int) {
	if *top2-*top1 > 1 {
		*top1++
		arr[*top1] = val
	} else {
		fmt.Println("Don't have space to store push element push1")
	}
}

func push2(arr *[5]int, s *Stack, val int, top1, top2 *int) {
	if *top2-*top1 > 1 {
		*top2--
		arr[*top2] = val
	} else {
		fmt.Println("Don't have space to store push element in push2")
	}
}

func pop1(arr *[5]int, s *Stack, val int, top1, top2 *int) {
	if *top1 != -1 {
		arr[*top1] = 0
		*top1--
	} else {
		fmt.Println("No element in pop1")
	}
}

func pop2(arr *[5]int, s *Stack, val int, top1, top2 *int) {
	if *top2 != len(arr) {
		arr[*top2] = 0
		*top2++
	} else {
		fmt.Println("No element in pop2")
	}
}

//reverse stack
// func reverseStack(str string){

// 	 var ans []rune
// 	for _,v:=range str{
// 		fmt.Println(" v   : ", v)
// 		ans = append(ans, v)
// 	}
// 	fmt.Println("reverse string : ", string(ans))
// }

// delete middle element
func deleteMiddle(s *Stack, stackLen int) {
	var count int = 0
	solve(s, count, stackLen)
}

func solve(s *Stack, count, stackLen int) {

	if count == stackLen/2 {
		pop(s)
		return
	}

	var num int = top(s)
	pop(s)
	//recurssive call
	count++
	solve(s, count, stackLen)

	push(s, num)
}

// pushbottom code logic
func pushBottom(s *Stack, val int) {

	//base case
	if isEmpty(s) {
		push(s, val)
		return
	}

	var num int = top(s)
	pop(s)

	//recurssive call
	pushBottom(s, val)

	push(s, num)
}

// reverse stack using recurssion
func reverseStack(s *Stack) {

	//base
	if isEmpty(s) {
		return
	}

	var num int = top(s)
	pop(s)

	//recurssive call
	reverseStack(s)
	pushBottom(s, num)
}

// sort Satck
func sortStack(s *Stack) {

	// fmt.Printf("sortstack :  %v \n", count)

	if isEmpty(s) {
		return
	}

	var num int = top(s)
	fmt.Println(" num :", num)
	pop(s)

	// count++
	//recurssive call
	sortStack(s)

	//insertAtSortedOrder
	inserAtSortedOrder(s, num)

}

func inserAtSortedOrder(s *Stack, val int) {

	if isEmpty(s) || (!isEmpty(s) && top(s) < val) {
		push(s, val)
		return
	}

	var num int = top(s)
	pop(s)
	//recurssive call
	inserAtSortedOrder(s, num)

	fmt.Println(" num insert : ", num)

	push(s, num)
}

// reduntdant bracket
func isRedundantBraces(str string) bool {

	var stack []rune
	var isRedundant bool

	for _, v := range str {
		switch v {
		case '(', '+', '-', '*', '/':
			stack = append(stack, v)
		default:
			if v == ')' {
				isRedundant = true
				for stack[len(stack)-1] != '(' {
					top := stack[len(stack)-1]

					if top == '+' || top == '-' || top == '*' || top == '/' {
						isRedundant = false
					}
					stack = stack[:len(stack)-1]
				}
				if isRedundant {
					return true
				}
				stack = stack[:len(stack)-1]
			}
		}
	}
	return false
	//  for _,v:=range str{
	// 	if v == '(' || v == '+' || v== '-' || v== '*' || v== '/' {
	// 			stack = append(stack, v)
	// 	}else{
	// 		if v == ')' {
	// 			isRedundant = true
	// 			for stack[len(stack)-1] != '(' {
	// 				top := stack[len(stack)-1]
	// 				fmt.Println("top  : ", top)

	// 				if top == '+' || top == '-' || top == '*' || top == '/' {
	// 					fmt.Println("top == '+' || top == '-' || top == '*' || top == '/'")
	// 					isRedundant = false
	// 					// break
	// 				}
	// 				stack = stack[ :len(stack)-1]
	// 			}

	// 			if isRedundant {
	// 				return true
	// 			}
	// 			stack = stack[ :len(stack)-1]
	// 		}
	// 	}
	//  }
	//  return false
}

// find min cost
func findMinCost(str string) int {

	var stack []rune
	//odd condition
	if len(str)%2 == 1 {
		return -1
	}

	for _, v := range str {
		if v == '{' {
			stack = append(stack, v)
		} else {
			//close braces
			if len(stack) > 0 && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, v)
			}
		}
	}

	//stack contains invalid express
	var a, b int
	// a for opening brances
	// b for closing braces
	for len(stack) != 0 {
		if stack[len(stack)-1] == '{' {
			a++
		} else {
			b++
		}
		stack = stack[:len(stack)-1]
	}

	var ans int = (a+1)/2 + (b+1)/2
	return ans
}

// next smaller element
func nextSmaller(arr []int) {
	var stack []int
	stack = append(stack, -1)
	var ans [4]int

	for i := len(arr) - 1; i >= 0; i-- {
		curr := arr[i]
		for stack[len(stack)-1] >= curr {
			stack = stack[:len(stack)-1]
		}
		ans[i] = stack[len(stack)-1]
		stack = append(stack, curr)
	}
	fmt.Println("stack : ", ans)

}

// largest rectangular area in histogram
func largestRectangleArea(heights []int) int {

	var n int = len(heights)
	// fmt.Println("n : ", n)

	// nextSmallerElement(heights, n)
	next := nextSmallerElement(heights, n)
	// return 0
	prev := prevSmallerElement(heights, n)

	var area int = 0

	for i := 0; i < n; i++ {
		var length int = heights[i]
		fmt.Println("length : ", length)

		if next[i] == -1 {
			next[i] = n
		}
		var breadth int = next[i] - prev[i] - 1
		fmt.Println("breadth : ", breadth)

		var newArea int = length * breadth
		fmt.Printf(" %v ===> %v\n", i, newArea)
		area = max(area, newArea)
		fmt.Printf(" %v ===> %v\n", i, area)
		break
	}
	return area
}

func nextSmallerElement(arr []int, n int) [6]int {
	var s []int
	s = append(s, -1)

	var ans [6]int
	// fmt.Println("arr : ", arr)

	for i := n - 1; i >= 0; i-- {
		// fmt.Println("i : ", i)
		curr := arr[i]
		// fmt.Println("CURR  : ", curr)

		for s[len(s)-1] != -1 && ans[len(s)-1] >= curr {
			// fmt.Println("Inside s[len(s)-s1] != -1 && ans[len(s)-1] >= curr ")
			s = s[:s[len(s)-1]]
		}
		// fmt.Println("ans : ", ans[i-1])
		ans[i] = s[len(s)-1]
		s = append(s, i)
	}
	return ans
}

func prevSmallerElement(arr []int, n int) [6]int {
	var s []int
	s = append(s, -1)

	var ans [6]int

	for i := 0; i < n; i++ {
		curr := arr[i]

		for s[len(s)-1] != -1 && ans[len(s)-1] >= curr {
			s = s[:s[len(s)-1]]
		}
		ans[i] = s[len(s)-1]
		s = append(s, i)
	}
	return ans
}

func max(a, b int) int {
	if a >= b {
		return a
	} else {
		return b
	}
}

// Link: https://leetcode.com/problems/implement-stack-using-queues/
type MyStack struct {
	q1    []int
	q2    []int
	front int //store the topmost elment of stack
}

// create a new MyStack object and return it
func Constructor() MyStack {
	return MyStack{[]int{}, []int{}, math.MinInt}
}

// push fuctionality of stack
func (this *MyStack) Push(x int) {
	//check is q2 is empty, if empty then push all element of q1 to q2 one by one
	if len(this.q2) == 0 {
		for len(this.q1) > 0 {
			var num int = this.q1[0]
			this.q1 = this.q1[1:]
			this.q2 = append(this.q2, num)
		}
	}
	//set the front and and append new element in queue q1
	if len(this.q1) == 0 {
		this.front = x
	}
	this.q1 = append(this.q1, x)
	//after adding new element to queue q1 add all prev elment in queue 1 from queue q1
	for len(this.q2) > 0 {
		this.q1 = append(this.q1, this.q2[0])
		this.q2 = this.q2[1:]
	}
}

func (this *MyStack) Pop() int {
	if len(this.q1) == 0 {
		return this.front
	}
	var num int = this.q1[0]
	this.q1 = this.q1[1:]
	return num
}

func (this *MyStack) Top() int {
	if len(this.q1) > 0 {
		return this.q1[0]
	}
	return this.front
}

func (this *MyStack) Empty() bool {
	return len(this.q1) == 0 && len(this.q2) == 0
}

func mstackusingQueue() {

	st := Constructor()
	st.Push(1)
	st.Push(2)
	st.Push(3)
	st.Push(4)
	st.Push(5)
	fmt.Println(" top : ", st.Top())
	fmt.Println("st : ", st)
	fmt.Println(" top : ", st.Pop())
	fmt.Println(" isEmpty : ", st.Empty())
	fmt.Println("st : ", st)
}

// Link: https://leetcode.com/problems/min-stack/
type MinStack struct {
	stack    []int
	minStack []int
}

func MinStackConstructor() MinStack {
	return MinStack{[]int{}, []int{}}
}

func (ms *MinStack) Push(val int) {
	// Push it into stack
	ms.stack = append(ms.stack, val)
	// If the new val is smaller current top in minStack, update
	var size = len(ms.minStack)
	if size == 0 {
		ms.minStack = append(ms.minStack, val)
	} else {
		if val <= ms.minStack[len(ms.minStack)-1] {
			ms.minStack = append(ms.minStack, val)
		}
	}
}

func (ms *MinStack) Pop() {
	if ms.stack[len(ms.stack)-1] == ms.minStack[len(ms.minStack)-1] {
		ms.stack = ms.stack[:len(ms.stack)-1]
		ms.minStack = ms.minStack[:len(ms.minStack)-1]
	} else {
		ms.stack = ms.stack[:len(ms.stack)-1]
	}
}

func (ms *MinStack) Top() int {
	return ms.stack[len(ms.stack)-1]
}

func (ms *MinStack) GetMin() int {
	return ms.minStack[len(ms.minStack)-1]
}

func minStack() {

	st := MinStackConstructor()
	st.Push(1)
	st.Push(2)
	st.Push(3)
	st.Push(4)
	st.Push(-1)
	fmt.Println(" top : ", st.Top())
	popEle := st.Top()
	fmt.Println(" top : ", st.Top())
	fmt.Printf(" popped element : %v  \n", popEle)
	fmt.Println(" getMin : ", st.GetMin())
	fmt.Println("stack : ", st)
}

// Link: https://takeuforward.org/data-structure/infix-to-postfix/
func infixToPostfix(exp string) string {
	fmt.Println("infix  : ", exp)
	var st []byte
	var ans string

	for i := 0; i < len(exp); i++ {
		char := exp[i]
		//if character is operand the add it to ans string
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'z' || char >= '0' && char <= '9' {
			ans += string(char)
		} else if string(char) == "(" {
			st = append(st, char)
		} else if string(char) == ")" {
			for string(st[len(st)-1]) != "(" {
				ans += string(st[len(st)-1])
				st = st[:len(st)-1]
			}
			st = st[:len(st)-1]
		} else {
			for len(st) > 0 && prec(exp[i]) <= prec(st[len(st)-1]) {
				ans += string(st[len(st)-1])
				st = st[:len(st)-1]
			}
			st = append(st, char)
		}
	}

	for len(st) > 0 {
		ans += string(st[len(st)-1])
		st = st[:len(st)-1]
	}
	return ans
}

func prec(c byte) int {
	if string(c) == "^" {
		return 3
	} else if string(c) == "/" || string(c) == "*" {
		return 2
	} else if string(c) == "+" || string(c) == "-" {
		return 1
	} else {
		return -1
	}
}

// Link: https://www.geeksforgeeks.org/convert-infix-prefix-notation/
// Link: https://www.youtube.com/watch?v=Gbs3kzYRdro&ab_channel=HelloWorld
func infixToPrefix(exp string) string {
	fmt.Println("infix  : ", exp)
	var st []byte
	var ans string

	for i := len(exp) - 1; i >= 0; i-- {
		char := exp[i]
		//if character is operand the add it to ans string
		if char >= 'a' && char <= 'z' || char >= 'A' && char <= 'z' || char >= '0' && char <= '9' {
			ans = string(char) + ans
		} else if string(char) == ")" {
			st = append(st, char)
		} else if string(char) == "(" {
			for string(st[len(st)-1]) != ")" {
				ans = string(st[len(st)-1]) + ans
				st = st[:len(st)-1]
			}
			st = st[:len(st)-1]
		} else {
			for len(st) > 0 && prec(exp[i]) <= prec(st[len(st)-1]) {
				ans = string(st[len(st)-1]) + ans
				st = st[:len(st)-1]
			}
			st = append(st, char)
		}
	}

	for len(st) > 0 {
		ans = string(st[len(st)-1]) + ans
		st = st[:len(st)-1]
	}

	// return reverse(ans)
	return ans
}

// reverse a string
func reverse(str string) string {
	var result string
	for _, v := range str {
		result = string(v) + result
	}
	return result
}

// Link: https://www.geeksforgeeks.org/prefix-infix-conversion/
func prefixToInfix(exp string) string {
	fmt.Println("prefix  : ", exp)

	var st = []string{}
	n := len(exp)

	for i := n - 1; i >= 0; i-- {
		c := exp[i]

		if isOperator(c) {
			op1 := st[len(st)-1]
			st = st[:len(st)-1]
			op2 := st[len(st)-1]
			st = st[:len(st)-1]

			temp := "(" + op1 + string(c) + op2 + ")"
			st = append(st, temp)
		} else {
			st = append(st, string(c))
		}
	}
	lastEle := st[len(st)-1]
	return lastEle
}

func isOperator(c byte) bool {
	switch c {
	case '+', '-', '*', '/', '^', '%':
		return true
	}
	return false
}

// Link: https://www.geeksforgeeks.org/prefix-postfix-conversion/
func prefixToPostfix(exp string) string {
	fmt.Println("prefix  : ", exp)

	var st = []string{}
	n := len(exp)

	for i := n - 1; i >= 0; i-- {
		c := exp[i]

		if isOperator(c) {
			op1 := st[len(st)-1]
			st = st[:len(st)-1]
			op2 := st[len(st)-1]
			st = st[:len(st)-1]

			temp := op1 + op2 + string(c)
			st = append(st, temp)
		} else {
			st = append(st, string(c))
		}
	}
	lastEle := st[len(st)-1]
	return lastEle
}

// Link: https://www.geeksforgeeks.org/postfix-prefix-conversion/
func postfixToPrefix(exp string) string {
	fmt.Println("postfix  : ", exp)

	var st = []string{}
	n := len(exp)

	for i := 0; i < n; i++ {
		c := exp[i]

		if isOperator(c) {
			op1 := st[len(st)-1]
			st = st[:len(st)-1]
			op2 := st[len(st)-1]
			st = st[:len(st)-1]

			temp := string(c) + op2 + op1
			st = append(st, temp)
		} else {
			st = append(st, string(c))
		}
	}
	lastEle := st[len(st)-1]
	return lastEle
}

// Link: https://www.geeksforgeeks.org/postfix-to-infix/
func postfixToInfix(exp string) string {
	fmt.Println("postfix  : ", exp)

	var st = []string{}
	n := len(exp)

	for i := 0; i < n; i++ {
		c := exp[i]

		if isOperator(c) {
			op1 := st[len(st)-1]
			st = st[:len(st)-1]
			op2 := st[len(st)-1]
			st = st[:len(st)-1]

			temp := "(" + op2 + string(c) + op1 + ")"
			st = append(st, temp)
		} else {
			st = append(st, string(c))
		}
	}
	lastEle := st[len(st)-1]
	return lastEle
}

// Link: https://leetcode.com/problems/next-greater-element-ii/
func nextGreaterElements2(nums []int) []int {

	n := len(nums)
	nge := make([]int, n)
	st := make([]int, 0)

	for i := range nums {
		nge[i] = -1
	}
	fmt.Println("nge  : ", nge)

	for i := 2*n - 1; i >= 0; i-- {

		for len(st) > 0 && st[len(st)-1] <= nums[i%n] {
			st = st[:len(st)-1]
		}

		if i < n {
			if len(st) > 0 {
				nge[i] = st[len(st)-1]
			}
		}
		st = append(st, nums[i%n])
	}
	return nge
}

// Link: https://leetcode.com/problems/next-greater-element-i/
func nextGreaterElement1(nums1 []int, nums2 []int) []int {
	st := make([]int, 0)
	arrMap := make(map[int]int, 0)
	res := make([]int, 0)

	for i := len(nums2) - 1; i >= 0; i-- {

		for len(st) > 0 && st[len(st)-1] <= nums2[i] {
			st = st[:len(st)-1]
		}
		if len(st) == 0 {
			arrMap[nums2[i]] = -1
			st = append(st, nums2[i])
		} else {
			arrMap[nums2[i]] = st[len(st)-1]
			st = append(st, nums2[i])
		}
	}

	for _, v := range nums1 {
		if val, exist := arrMap[v]; exist {
			res = append(res, val)
		}
	}
	return res
}

// Link: https://leetcode.com/problems/trapping-rain-water/
func trap(height []int) int {

	n := len(height)
	left := make([]int, n)
	right := make([]int, n)
	ans := 0
	//store max of rain at each index from left
	left[0] = height[0]
	for i := 1; i < n; i++ {
		left[i] = maxi(left[i-1], height[i])
	}
	//store max of rain at each index from right
	right[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		right[i] = maxi(right[i+1], height[i])
	}
	for i := 0; i < n; i++ {
		a1 := mini(left[i], right[i]) - height[i]
		ans += a1
	}
	return ans
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

// Link: https://www.geeksforgeeks.org/next-smaller-element/
func nextSmallerElements2(nums []int) []int {
	fmt.Println("nums   : ", nums)

	n := len(nums)
	nge := make([]int, n)
	st := make([]int, 0)

	for i := range nums {
		nge[i] = -1
	}
	fmt.Println("nge  : ", nge)

	for i := n - 1; i >= 0; i-- {

		for len(st) > 0 && st[len(st)-1] >= nums[i%n] {
			st = st[:len(st)-1]
		}

		if i < n {
			if len(st) > 0 {
				nge[i] = st[len(st)-1]
			}
		}
		st = append(st, nums[i%n])
	}
	return nge
}

// ramaining
// Link: https://leetcode.com/problems/next-greater-element-ii/
func nextGreaterElementCount(nums []int) []int {
	fmt.Println("nums   : ", nums)
	n := len(nums)
	nge := make([]int, n)
	st := make([]int, 0)

	for i := range nums {
		nge[i] = 0
	}
	fmt.Println("nge  : ", nge)

	for i := n - 1; i >= 0; i-- {

		for len(st) > 0 && st[len(st)-1] > nums[i] {
			st = st[:len(st)-1]
		}

		if i < n {
			if len(st) > 0 {
				nge[i] = len(st)
			}
		}
		st = append(st, nums[i%n])
	}
	return nge
}
