package main

import "fmt"

func (list *LinkedList) middleOfLinkedList() {
	if list.headNode == nil || list.headNode.nextNode == nil {
		fmt.Println("middle : ", list.headNode)
		return
	}

	if list.headNode.nextNode.nextNode == nil {
		fmt.Println("middle : ", list.headNode.nextNode)
		return
	}

	var slow *Node = list.headNode
	var fast *Node = list.headNode.nextNode

	for fast != nil {
		fast = fast.nextNode
		if fast != nil {
			fast = fast.nextNode
		}
		slow = slow.nextNode
	}
	fmt.Println("middle : ", slow)
}

func (list *LinkedList) detectCycle() {

	if list.headNode == nil || list.headNode.nextNode == nil {
		fmt.Println("there is no cycle in linked list")
		return
	}

	// floyd algo for cycle detection
	var slow *Node = list.headNode
	var fast *Node = list.headNode

	for slow != nil && fast != nil {
		fast = fast.nextNode

		if fast != nil {
			fast = fast.nextNode
		}
		slow = slow.nextNode

		if slow == fast {
			fmt.Println("there is cycle in linked list")
			return
		}
	}
	fmt.Println("there is no cycle in linked list")
	return
}

//Implement stack using linked list
//Link: https://www.educative.io/answers/how-to-implement-a-stack-using-a-linked-list-in-go
//Link: https://takeuforward.org/data-structure/implement-stack-using-linked-list/
type stackLinkedList struct {
	head *Node
	size int
}

func Constructor() stackLinkedList {
	return stackLinkedList{nil, 0}
}

func (s *stackLinkedList) Size() int {
	return s.size
}

func (s *stackLinkedList) IsEmpty() bool {
	return s.size == 0
}

func (s *stackLinkedList) Top() int {
	if s.IsEmpty() {
		fmt.Println("stack is empty")
		return 0
	}
	return s.head.property
}

func (s *stackLinkedList) Push(x int) {
	s.head = &Node{x, s.head}
	s.size++
}

func (s *stackLinkedList) Pop() (int, bool) {
	if s.IsEmpty() {
		fmt.Println("stack is empty")
		return 0, false
	}
	poppedElement := s.head.property
	s.head = s.head.nextNode
	s.size--
	return poppedElement, true
}

func stackUsingLinkedList() {

	st := Constructor()
	st.Push(1)
	st.Push(2)
	st.Push(3)
	st.Push(4)
	st.Push(5)
	fmt.Println(" top : ", st.Top())
	popEle, _ := st.Pop()
	fmt.Println(" top : ", st.Top())
	fmt.Printf(" popped element : %v  \n", popEle)
	fmt.Println(" isEmpty : ", st.IsEmpty())
	fmt.Println("stack : ", st)
}

//
//Implement queue using linked list
//Link: https://www.educative.io/answers/how-to-implement-a-stack-using-a-linked-list-in-go
//Link: https://takeuforward.org/data-structure/implement-stack-using-linked-list/
type queueLinkedList struct {
	head *Node
	size int
}

func QConstructor() queueLinkedList {
	return queueLinkedList{nil, 0}
}

func (s *queueLinkedList) Size() int {
	return s.size
}

func (s *queueLinkedList) IsEmpty() bool {
	return s.size == 0
}

func (s *queueLinkedList) Top() int {
	if s.IsEmpty() {
		fmt.Println("stack is empty")
		return 0
	}
	return s.head.property
}

func (s *queueLinkedList) Push(x int) {
	if s.size == 0 {
		s.head = &Node{x, nil}
		s.size++
	} else {
		temp := s.head
		for temp.nextNode != nil {
			temp = temp.nextNode
		}
		var newNode = &Node{x, nil}
		temp.nextNode = newNode
		s.size++
	}
}

func (s *queueLinkedList) Pop() (int, bool) {
	if s.IsEmpty() {
		fmt.Println("stack is empty")
		return 0, false
	}
	poppedElement := s.head.property
	s.head = s.head.nextNode
	s.size--
	return poppedElement, true
}

func queueUsingLinkedList() {

	q := QConstructor()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Push(5)
	fmt.Println(" top : ", q.Top())
	popEle, _ := q.Pop()
	fmt.Printf(" popped element : %v  \n", popEle)
	fmt.Println(" top : ", q.Top())
	fmt.Println(" isEmpty : ", q.IsEmpty())
	popEle, _ = q.Pop()
	fmt.Printf(" popped element : %v  \n", popEle)
	popEle, _ = q.Pop()
	fmt.Printf(" popped element : %v  \n", popEle)
	popEle, _ = q.Pop()
	fmt.Printf(" popped element : %v  \n", popEle)
	fmt.Println(" top : ", q.Top())
	fmt.Println("stack : ", q)

}
