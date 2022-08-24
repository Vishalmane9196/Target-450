// //push
// //pop
// //top
// //print
// //length
// //isempty

// package main

// import "fmt"

// type Stack struct {
//     stk []int
// 		// top int	
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
// 	poppedElement := s.stk[stackLength]
// 	fmt.Println("removed the element : ", poppedElement)

// 	s.stk = s.stk[:stackLength]
// }

// func (s *Stack) top(){
//   top := s.stk[len(s.stk)-1]
// 	fmt.Println("top element : ", top )
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
// 	s.print()
	
// 	s.pop()
	
// 	s.print()
// 	s.push(9)
// 	s.print()
// 	s.top()
// }
