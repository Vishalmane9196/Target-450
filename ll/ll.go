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
	var newNode = &Node{}
	newNode.property = property
	newNode.nextNode = nil
	if list.headNode != nil {
		newNode.nextNode = list.headNode
	}
	list.headNode = newNode
}

// insertAfter method  adds a node with nodeProperty after node with property
func (list *LinkedList) insertAfter(nodeData int, newData int) {
	var newNode = &Node{}
	newNode.property = newData
	newNode.nextNode = nil

	var currNode *Node

	currNode = list.getNode(nodeData)
	if currNode != nil {
		newNode.nextNode = currNode.nextNode
		currNode.nextNode = newNode
	}
}

// inserAtEnd method adds the node with property to the end
func (list *LinkedList) inserAtEnd(newData int) {
	var node = &Node{}
	node.property = newData
	node.nextNode = nil

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
	var prevDeleteNode = &Node{}

	deleteNode = list.getNode(deleteData)
	prevDeleteNode = list.getPreviousNode(deleteData)

	if deleteNode != nil {
		prevDeleteNode.nextNode = deleteNode.nextNode
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
func (linkedList *LinkedList) getLastNode() *Node {
	var tempNode *Node
	var lastNode *Node
	for tempNode = linkedList.headNode; tempNode != nil; tempNode = tempNode.nextNode {
		if tempNode.nextNode == nil {
			lastNode = tempNode
		}
	}
	return lastNode
}

// getFirstNode method returns the first Node of linked list
func (linkedList *LinkedList) getFirstNode() *Node {
	return linkedList.headNode
}

// getNode method returns Node given parameter property
func (linkedList *LinkedList) getNode(data int) *Node {
	var tempNode *Node
	var nodeWith *Node
	for tempNode = linkedList.headNode; tempNode != nil; tempNode = tempNode.nextNode {
		if tempNode.property == data {
			nodeWith = tempNode
			break
		}
	}
	return nodeWith
}

// getPreviousNode method returns 2nd last elment of given property
func (linkedList *LinkedList) getPreviousNode(data int) *Node {
	var tempNode *Node
	var nodeWith *Node
	for tempNode = linkedList.headNode; tempNode.nextNode != nil; tempNode = tempNode.nextNode {
		if tempNode.nextNode.property == data {
			nodeWith = tempNode
			break
		}
	}
	return nodeWith
}

// isListEmpty method return the linked list is empty or not
func (l *LinkedList) isListEmpty() bool {
	return l.headNode == nil
}

// printLinkedlist method iterates over LinkedList
func (linkedList *LinkedList) printLinkedlist() {
	var node *Node
	fmt.Printf(" [START]")
	for node = linkedList.headNode; node != nil; node = node.nextNode {
		fmt.Printf(" <-> [%v]", node.property)
	}
	fmt.Printf(" <-> [END]\n")
}
