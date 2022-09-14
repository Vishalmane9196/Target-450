package main

func main() {

	var linkedList LinkedList

	linkedList = LinkedList{}

	linkedList.insertAtStart(1)
	linkedList.insertAtStart(2)
	linkedList.insertAtStart(3)
	linkedList.insertAtStart(4)
	linkedList.inserAtEnd(5)
	linkedList.insertAfter(2,9)
	linkedList.printLinkedList()
	linkedList.delAtStart()
	linkedList.delAtEnd()
	linkedList.delNode(2)
	linkedList.printLinkedList()
	linkedList.reverseLinkedList()
	linkedList.printLinkedList()
	
}