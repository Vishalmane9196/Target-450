package main

import "fmt"

type Stack struct {
	stk []int
}

func (s *Stack) push(val int) {
	s.stk = append(s.stk, val)
}

func (s *Stack) isEmpty() bool {
	return len(s.stk) == 0
}

func (s *Stack) getLength() int {
	return len(s.stk)
}

func (s *Stack) printStack() {
	fmt.Println("printing the Stack")
	fmt.Printf("[BEG] <-> ")
	for i := 0; i < len(s.stk); i++ {
		fmt.Printf("[%v] <-> ", s.stk[i])
	}
	fmt.Printf("[END]\n")
}

func (s *Stack) pop() int {
	stackLength := len(s.stk) - 1
	poppedElement := s.stk[stackLength]
	s.stk = s.stk[:stackLength]
	return poppedElement
}

func (s *Stack) top() int {
	top := s.stk[len(s.stk)-1]
	return top
}

func main() {

	/******stack implementaion using oop******/
	var s Stack
	isempty := s.isEmpty()
	fmt.Println("isEmpty  : ", isempty)
	s.push(5)
	s.push(4)
	s.push(3)
	s.push(2)
	s.push(1)
	s.printStack()

	s.pop()
	s.push(9)
	s.printStack()
	s.top()
}
