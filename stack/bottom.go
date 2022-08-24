// package main

// import "fmt"

// type Stack struct {
//     stk []int
// }

// func (s *Stack) push(val int){
// 	s.stk = append(s.stk, val)
// }

// func (s *Stack) isEmpty() bool {
// 	return len(s.stk) == 0
// }

// func (s *Stack) getLength() int {
// 	return len(s.stk)
// }

// func (s *Stack) print() {
// 	fmt.Println("printing an array")
// 	fmt.Printf("[BEG] <-> ")
// 	for i:=0; i<len(s.stk); i++{
// 		fmt.Printf("[%v] <-> ",s.stk[i])
// 	}
// 	fmt.Printf("[END]\n")
// }

// func (s *Stack) pop() {
// 	stackLength := len(s.stk)-1
// 	// poppedElement := s.stk[stackLength]s
// 	// fmt.Println("removed the element : ", poppedElement)

// 	s.stk = s.stk[:stackLength]
// }

// func (s *Stack) top() int {
//   top := s.stk[len(s.stk)-1]
// 	// fmt.Println("top element : ", top )
// 	return top
// }

// //pushbottom code logic
// func (s *Stack) pushBottom(x int){

// 	//base case
// 	if s.isEmpty(){
//      s.push(x)
// 		 return;
// 	}

// 	 var num int = s.top()
// 	 s.pop()

// 	 //recurssive call
// 	 s.pushBottom(x)

// 	 s.push(num)
// }

// //delete middle element of stack
// func (s *Stack) deleteMiddle(count int){

// 	fmt.Println("length : ", s.getLength())

// //base case
// 	if count == s.getLength()/2 {
// 		s.pop()
// 		return;
// 	}

// 	var num int = s.top()
//   s.pop()
// 	//recurrsive call
// 	s.deleteMiddle(count+1)

// 	s.push(num)
// }

// func isValidParenthesis(str string) bool{

// 	var stack []rune

// 	for _, v:=range str {
// 		switch v {
// 		case '(', '[', '{' :
// 			 stack = append(stack, v)
// 		case ')' :
// 			if len(stack) >0 && stack[len(stack)-1] == '(' {
//          stack = stack[:len(stack)-1]
// 			}else{
// 				return false
// 			}
// 		case ']' :
// 			if len(stack) >0 && stack[len(stack)-1] == '[' {
//          stack = stack[:len(stack)-1]
// 			}else{
// 				return false
// 			}
// 		case '}' :
// 			if len(stack) >0 && stack[len(stack)-1] == '{' {
//          stack = stack[:len(stack)-1]
// 			}else{
// 				return false
// 			}
// 		}
// 	}
// 	return len(stack) == 0
// }

// func (s *Stack)reverseStack(){

// 	//base
// 	if s.isEmpty() {
// 		return
// 	}
	
// 	var num int = s.top()
// 	s.pop() 

// 	//recurssive call
// 	s.reverseStack()
// 	s.pushBottom(num)
// }

// func inserAtSortedOrder(s *Stack, val int){

// 	// fmt.Println(" val :  ", val)
// 	if s.isEmpty() || (!s.isEmpty() && s.top() < val) {
// 		s.push(val)
// 		return
// 	}

// 	var num int = s.top()
//   s.pop()

// 	//recurssive call
// 	inserAtSortedOrder(s, num)

// 	s.push(num)
// }

// func (s *Stack)sortStack(count int){

// 	fmt.Printf("sortstack :  %v \n", count)

// 	if s.isEmpty() {
// 		return
// 	}

// 	var num int = s.top()
// 	fmt.Println(" num :", num)
// 	s.pop()

//   count++
// 	//recurssive call
// 	s.sortStack(count)

// 	//insertAtSortedOrder
// 	inserAtSortedOrder(s, num)

// }

// func main(){
// 	var s Stack

// 	isempty := s.isEmpty()
// 	fmt.Println("isEmpty  : ", isempty)
// 	s.push(5)
// 	s.push(4)
// 	s.push(3)
// 	s.push(2)
// 	s.push(1)
// 	// s.print()
	
// 	s.pop()
	
// 	// s.print()
// 	s.push(9)
// 	s.print()
// 	// s.top()

// 	//stack Initialization done

// 	//inser element at bottom of stack

// 	// s.pushBottom(9)
// 	// s.print()

// 	// s.deleteMiddle(0)
// 	// s.print()

// 	// res := isValidParenthesis("()")
// 	// fmt.Println("is valid parenthesis : ", res)

// 	// s.reverseStack()
// 	// s.print()

// 	fmt.Println("**********************************************************")
// 	//need to implement it again not working
// 	// s.sortStack(0)
// 	// s.print()
	 






// }