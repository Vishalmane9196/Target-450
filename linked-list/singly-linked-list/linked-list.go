package main

import (
	"fmt"
)

// Node class
type Node struct {
	property int
	nextNode *Node
}

// LinkedList class
type LinkedList struct {
	headNode *Node
}

// insertAtStart method will be used for isnerting node at starting position
func (list *LinkedList) insertAtStart(property int) {
	var newNode = &Node{property, nil}
	
	if list.headNode != nil {
		newNode.nextNode = list.headNode
	}
	list.headNode = newNode
}

// insertAfter method  adds a node with nodeProperty after node with property
func (list *LinkedList) insertAfter(nodeData int, newData int) {
	var newNode = &Node{newData, nil}
	var currNode *Node

	currNode = list.getNode(nodeData)
	if currNode != nil {
		newNode.nextNode = currNode.nextNode
		currNode.nextNode = newNode
	}
}

// inserAtEnd method adds the node with property to the end
func (list *LinkedList) inserAtEnd(newData int) {
	var node = &Node{newData, nil}
	var lastNode *Node

	lastNode = list.getLastNode()
	if lastNode != nil {
		lastNode.nextNode = node
	}
}

// deleteAtStart method remove the first node from linked list
func (list *LinkedList) deleteAtStart() {
	if !list.isListEmpty() {
		list.headNode = list.headNode.nextNode
	}
}

// removeNode method remove the data from linked list
func (list *LinkedList) removeAfterNode(deleteData int) {
	var deleteNode = &Node{}
	var prevOfDeleteNode = &Node{}

	deleteNode = list.getNode(deleteData)
	prevOfDeleteNode = list.getPreviousNode(deleteData)
	if deleteNode != nil {
		prevOfDeleteNode.nextNode = deleteNode.nextNode
		deleteNode.nextNode = nil
	}
}

// deleteAtEnd method remove the last node from linked list
func (list *LinkedList) deleteAtEnd() {
	var lastNode = &Node{}
	var prevLastNode = &Node{}

	lastNode = list.getLastNode()
	prevLastNode = list.getPreviousNode(lastNode.property)
	if lastNode != nil {
		prevLastNode.nextNode = nil
		lastNode.nextNode = nil
	}
}

// getLastNode method returns the last Node
func (list *LinkedList) getLastNode() *Node {
	var tempNode *Node
	var lastNode *Node

	for tempNode = list.headNode; tempNode != nil; tempNode = tempNode.nextNode {
		if tempNode.nextNode == nil {
			lastNode = tempNode
		}
	}
	return lastNode
}

// getFirstNode method returns the first Node of linked list
func (list *LinkedList) getFirstNode() *Node {
	return list.headNode
}

// getNode method returns Node given parameter property
func (list *LinkedList) getNode(data int) *Node {
	var tempNode *Node
	var currNode *Node
	for tempNode = list.headNode; tempNode != nil; tempNode = tempNode.nextNode {
		if tempNode.property == data {
			currNode = tempNode
			break
		}
	}
	return currNode
}

// getPreviousNode method returns 2nd last elment of given property
func (list *LinkedList) getPreviousNode(data int) *Node {
	var tempNode *Node
	var PrevNode *Node
	for tempNode = list.headNode; tempNode.nextNode != nil; tempNode = tempNode.nextNode {
		if tempNode.nextNode.property == data {
			PrevNode = tempNode
			break
		}
	}
	return PrevNode
}

// isListEmpty method return the linked list is empty or not
func (list *LinkedList) isListEmpty() bool {
	return list.headNode == nil
}

func (list *LinkedList) getLength() int {
   var travNode = &Node{}
   var count int

   for travNode = list.headNode; travNode != nil; travNode = travNode.nextNode {
      count++
   }
   return count
}

// printLinkedlist method iterates over LinkedList
func (list *LinkedList) printLinkedlist() {
	var node *Node

	fmt.Printf(" [START]")
	for node = list.headNode; node != nil; node = node.nextNode {
		fmt.Printf(" <-> [%v]", node.property)
	}
	fmt.Printf(" <-> [END]\n")
}


//not working
func (list *LinkedList) reverseLinkedList() {

	var prev *Node
	var curr = list.headNode
	for curr != nil {
        //stored next pointer reference
        tmp := curr.nextNode
        //changed the next of current to prev
        curr.nextNode = prev 
        //shifted prev and curr poniter by 1 ahead
        prev = curr
        curr = tmp
	}
}