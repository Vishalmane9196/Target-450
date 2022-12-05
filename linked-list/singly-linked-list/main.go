package main

import "fmt"

func main() {

	/***********Linked list implmentation***************/
	linkedList := &LinkedList{}

	linkedList.insertAtStart(1)
	linkedList.insertAtStart(3)
	linkedList.inserAtEnd(4)
	linkedList.inserAtEnd(5)
	linkedList.insertAfter(1, 7)
	linkedList.insertAfter(7, 6)
	linkedList.insertAfter(7, 9)

	// linkedList.printLinkedlist()
	// linkedList.deleteAtStart()
	// linkedList.printLinkedlist()
	// linkedList.removeAfterNode(4)
	// linkedList.deleteAtEnd()
	// linkedList.printLinkedlist()
	// fmt.Println(" is list empty : ", linkedList.isListEmpty())
	// linkedList.getLength()
	// fmt.Println(" Linked list length  : ", linkedList.getLength())
	// tempList := linkedList.reverseLinkedList()
	// tempList.printLLByNode()
	linkedList.printLinkedlist()

	/**********Problem***********/
	middle := linkedList.middleOfLinkedList()
	fmt.Println("middle elment of linked list: ", middle.property)

	// linkedList.headNode.nextNode.nextNode.nextNode.nextNode = linkedList.headNode.nextNode
	// res := linkedList.hasCycle()
	// fmt.Println("hasCycle : ", res)

	// linkedList.headNode.nextNode.nextNode.nextNode.nextNode = linkedList.headNode.nextNode
	// res := linkedList.detectCycleNode()
	// fmt.Println("detectCycleNode : ", res)

	// stackUsingLinkedList()
	// queueUsingLinkedList()
}
