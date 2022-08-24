package main

import "fmt"

type Stack struct {
	stk []int
}

func push(s *Stack, val int) {
	s.stk = append(s.stk, val)
}

func isEmpty(s *Stack) bool {
	return len(s.stk) == 0
}

func (s *Stack) getLength() int {
	return len(s.stk)
}

func print(s *Stack) {
	// fmt.Println("printing an array")
	fmt.Printf("[BEG] <-> ")
	for i := 0; i < len(s.stk); i++ {
		fmt.Printf("[%v] <-> ", s.stk[i])
	}
	fmt.Printf("[END]\n")
}

func pop(s *Stack) {
	stackLength := len(s.stk) - 1
	// poppedElement := s.stk[stackLength]s
	// fmt.Println("removed the element : ", poppedElement)

	s.stk = s.stk[:stackLength]
}

func top(s *Stack) int {
	// stacklength := s.stk[len(s.stk)-1]
	// fmt.Println("stacklength : ", stacklength )
	top := s.stk[len(s.stk)-1]
	// fmt.Println("top element : ", top )
	return top
}

func isValidParenthesis(str string) bool {

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

//implement 2 stack using array
func push1(arr *[5]int, s *Stack, val int , top1, top2 *int) {
	if *top2 - *top1 > 1 {
     *top1++
		 arr[*top1] = val
	}else{
		fmt.Println("Don't have space to store push element push1")
	}
}

func push2(arr *[5]int, s *Stack, val int , top1, top2 *int) {
	if *top2 - *top1 > 1 {
     *top2--
		 arr[*top2] = val
	}else{
		fmt.Println("Don't have space to store push element in push2")
	}
}

func pop1(arr *[5]int, s *Stack, val int , top1, top2 *int) {
	if *top1 != -1 {
		arr[*top1] = 0
     *top1--
	}else{
		fmt.Println("No element in pop1")
	}
}

func pop2(arr *[5]int, s *Stack, val int , top1, top2 *int) {
	if *top2 != len(arr) {
     arr[*top2] = 0
     *top2++
	}else{
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

//delete middle element
func deleteMiddle(s *Stack, stackLen int) {
	var count int = 0
  solve(s, count, stackLen)
}

func solve(s *Stack, count, stackLen int){

	if count == stackLen/2 {
		pop(s)
		return
	}

	var num int = top(s)
	pop(s)
	//recurssive call
	count++
	solve(s , count, stackLen)

	push(s, num)
}

//pushbottom code logic
func pushBottom(s *Stack, val int){

	//base case
	if isEmpty(s){
     push(s, val)
		 return;
	}

	 var num int = top(s)
	 pop(s)

	 //recurssive call
	 pushBottom(s, val)

	 push(s, num)
}

//reverse stack using recurssion
func reverseStack(s *Stack){

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


//sort Satck
func sortStack(s *Stack){

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

func inserAtSortedOrder(s *Stack, val int){

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

//reduntdant bracket
func isRedundantBraces(str string) bool {

	var stack []rune
	var isRedundant bool

	for _, v := range str {
		switch v {
		case '(', '+', '-', '*', '/':
			stack = append(stack, v)
		default :
		  if v == ')' {
				isRedundant = true
				for stack[len(stack)-1] != '(' {
					top := stack[len(stack)-1]
	
						if top == '+' || top == '-' || top == '*' || top == '/' {
							isRedundant = false
						}
						stack = stack[ :len(stack)-1]
				}
				if isRedundant {
					return true
				}
				stack = stack[ :len(stack)-1]
			}
		}
	}
	return  false
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

//find min cost
func findMinCost(str string) int {

	var stack []rune
	//odd condition
	if len(str) %2 == 1 {
		return -1
	}

	for _,v:= range str {
    if v == '{' {
			stack = append(stack, v)
		}else{
			//close braces
			if len(stack) > 0 && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			}else{
				stack = append(stack, v)
			}
		}
	}

		//stack contains invalid express
		var a,b int
		// a for opening brances 
		// b for closing braces
		for len(stack) != 0 {
			if stack[len(stack)-1] == '{' {
				a++;
			}else{
				b++
			}
			stack = stack[:len(stack)-1]
		}

		var ans int = (a+1)/2 + (b+1)/2
		return ans
}

//next smaller element
func nextSmaller(arr []int) {
  var stack []int 
	stack = append(stack, -1)
	var ans [4]int

	for i:= len(arr)-1; i>=0; i-- {
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
    
    for i:=0; i < n; i++ {
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
    
    for i:=n-1; i >=0; i-- {
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
    
    for i:=0; i < n; i++ {
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
    }else{
        return b
    }
}
