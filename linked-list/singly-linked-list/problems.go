package main

import "fmt"

//Link: https://leetcode.com/problems/middle-of-the-linked-list/
func (list *LinkedList) middleOfLinkedList() *Node {
	if list.headNode == nil || list.headNode.nextNode == nil {
		return list.headNode
	}

	if list.headNode.nextNode.nextNode == nil {
		return list.headNode.nextNode
	}

	slow := list.headNode
	fast := list.headNode.nextNode
	for fast != nil {
		fast = fast.nextNode
		if fast != nil {
			fast = fast.nextNode
		}
		slow = slow.nextNode
	}
	return slow
}

//Link: https://leetcode.com/problems/linked-list-cycle/
func (list *LinkedList) hasCycle() bool {

	/*
		approach 1 - Hashing Approach:
		Traverse the list one by one and keep putting the node addresses in a Hash Table. At any point, if NULL is reached then return false,
		and if the next of the current nodes points to any of the previously stored nodes in  Hash then return true.
	*/

	/*
		Solution 2: 	Have a visited flag with each node.
		Approach: This solution requires modifications to the basic linked list data structure.
		Traverse the linked list and keep marking visited nodes.
		If you see a visited node again then there is a loop.
	*/

	/*
		Solution 3: Floyd’s Cycle-Finding Algorithm
		Traverse linked list using two pointers.
		Move one pointer(slow) by one and another pointer(fast) by two.
		If these pointers meet at the same node then there is a loop. If pointers do not meet then linked list doesn’t have a loop.
	*/

	//check first two nodes of linked list empty or not
	if list.headNode == nil || list.headNode.nextNode == nil {
		return false
	}

	//set slow and fast pointer to first elment of linked list
	var slow *Node = list.headNode
	var fast *Node = list.headNode

	for slow != nil && fast != nil {
		fast = fast.nextNode

		if fast != nil {
			fast = fast.nextNode
		}
		slow = slow.nextNode

		/* floyd algo for cycle detection
		At one point slow and fast will become equal if linked list has cycle present in it
		*/
		if slow == fast {
			return true
		}
	}
	return false
}

func (list *LinkedList) detectCycleNode() int {

	if list.headNode == nil || list.headNode.nextNode == nil {
		return -1
	}

	var slow *Node = list.headNode
	var fast *Node = list.headNode

	for slow != nil && fast != nil {
		fast = fast.nextNode

		if fast != nil {
			fast = fast.nextNode
		}
		slow = slow.nextNode

		if slow == fast {
			return slow.property
		}
	}
	return -1
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

// Top return the head value
func (s *stackLinkedList) Top() int {
	if s.IsEmpty() {
		fmt.Println("stack is empty")
		return 0
	}
	return s.head.property
}

//Push create new node and keep head in new node's next pointer
//like a pushfront operation
func (s *stackLinkedList) Push(x int) {
	s.head = &Node{x, s.head}
	s.size++
}

//Pop move head to head's next and decrease size by 1
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

func (s *stackLinkedList) PrintStack() {
	if !s.IsEmpty() {
		fmt.Printf("[START]")
		for trav := s.head; trav != nil; trav = trav.nextNode {
			fmt.Printf(" <-> [%v]", trav.property)
		}
		fmt.Printf(" <-> [END]\n")
	}
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
	fmt.Printf(" popped element : %v  \n", popEle)
	fmt.Println(" top : ", st.Top())
	fmt.Println(" isEmpty : ", st.IsEmpty())
	st.PrintStack()
}

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

//Push create new node and add the newly created node in last node's nextnode
//if list is empty then first pushed node will become head node
//like a pushback operation
func (s *queueLinkedList) Push(x int) {

	//if list is empty then set new node as head node and increase size by 1
	if s.size == 0 {
		s.head = &Node{x, nil}
		s.size++
	} else {
		//temp variable store the last node of list
		temp := s.head
		for temp.nextNode != nil {
			temp = temp.nextNode
		}
		//set the last node's next poniter to newly created node and increase size by 1
		var newNode = &Node{x, nil}
		temp.nextNode = newNode
		s.size++
	}
}

//Pop move head to head's next and decrease size by 1
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

func (s *queueLinkedList) PrintQueue() {
	if !s.IsEmpty() {
		fmt.Printf("[START]")
		for trav := s.head; trav != nil; trav = trav.nextNode {
			fmt.Printf(" <-> [%v]", trav.property)
		}
		fmt.Printf(" <-> [END]\n")
	}
}

func queueUsingLinkedList() {

	q := QConstructor()
	q.Push(1)
	q.Push(2)
	q.Push(3)
	q.Push(4)
	q.Push(5)
	q.PrintQueue()
	fmt.Println(" top : ", q.Top())
	popEle, _ := q.Pop()
	fmt.Printf(" popped element : %v  \n", popEle)
	fmt.Println(" top : ", q.Top())
	fmt.Println(" isEmpty : ", q.IsEmpty())
	q.PrintQueue()
}
