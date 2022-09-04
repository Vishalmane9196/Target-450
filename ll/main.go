package main

import "fmt"

func main() {

	var linkedList LinkedList
	linkedList = LinkedList{}

	linkedList.insertAtStart(1)
	linkedList.insertAtStart(3)
	linkedList.inserAtEnd(4)
	linkedList.inserAtEnd(5)
	linkedList.insertAfter(1, 7)

	linkedList.printLinkedlist()
	linkedList.deleteAtStart()
	linkedList.printLinkedlist()
	linkedList.removeAfterNode(4)
	linkedList.deleteAtEnd()
	linkedList.printLinkedlist()
	fmt.Println("is list empty : ", linkedList.isListEmpty())

}
