package main

import (
	"fmt"
)

/***********Linked list implmentation***************/
// Node class
type Node struct {
	property int
	nextNode *Node
}

// LinkedList class
type LinkedList struct {
	headNode *Node
}

// insertAtStart inserts node at start of linked list
func (list *LinkedList) insertAtStart(property int) {
	var newNode = &Node{property, nil}

	if list.headNode != nil {
		newNode.nextNode = list.headNode
	}
	list.headNode = newNode
}

// insertAfter inserts newNode after nodeData node
func (list *LinkedList) insertAfter(nodeData int, newData int) {
	var newNode = &Node{newData, nil}

	currNode := list.getNode(nodeData)
	if currNode != nil {
		newNode.nextNode = currNode.nextNode
		currNode.nextNode = newNode
	}
}

// inserAtEnd method insert node at the end of linked list
func (list *LinkedList) inserAtEnd(newData int) {
	var node = &Node{newData, nil}

	lastNode := list.getLastNode()
	if lastNode != nil {
		lastNode.nextNode = node
	}
}

// deleteAtStart delete first node from list
func (list *LinkedList) deleteAtStart() {
	if !list.isListEmpty() {
		list.headNode = list.headNode.nextNode
	}
}

// removeAfterNode remove node from list
func (list *LinkedList) removeAfterNode(deleteData int) {

	deleteNode := list.getNode(deleteData)
	prevOfDeleteNode := list.getPreviousNode(deleteData)
	if deleteNode != nil {
		prevOfDeleteNode.nextNode = deleteNode.nextNode
		deleteNode.nextNode = nil
	}
}

// deleteAtEnd delete last node from linked list
func (list *LinkedList) deleteAtEnd() {

	lastNode := list.getLastNode()
	prevLastNode := list.getPreviousNode(lastNode.property)
	if lastNode != nil {
		prevLastNode.nextNode = nil
		lastNode.nextNode = nil
	}
}

// getLastNode returns the last Node pointer
func (list *LinkedList) getLastNode() *Node {

	var lastNode *Node
	for tempNode := list.headNode; tempNode != nil; tempNode = tempNode.nextNode {
		if tempNode.nextNode == nil {
			lastNode = tempNode
			break
		}
	}
	return lastNode
}

// getFirstNode returns the first Node of linked list
func (list *LinkedList) getFirstNode() *Node {
	return list.headNode
}

// getNode returns node having data attribute
func (list *LinkedList) getNode(data int) *Node {

	var currNode *Node
	for tempNode := list.headNode; tempNode != nil; tempNode = tempNode.nextNode {
		if tempNode.property == data {
			currNode = tempNode
			break
		}
	}
	return currNode
}

// getPreviousNode returns prvious node of given node
func (list *LinkedList) getPreviousNode(data int) *Node {
	var PrevNode *Node
	for tempNode := list.headNode; tempNode.nextNode != nil; tempNode = tempNode.nextNode {
		if tempNode.nextNode.property == data {
			PrevNode = tempNode
			break
		}
	}
	return PrevNode
}

// isListEmpty return the linked list is empty or not
func (list *LinkedList) isListEmpty() bool {
	return list.headNode == nil
}

// getLength returns length of linked list
func (list *LinkedList) getLength() int {

	var count int
	for travNode := list.headNode; travNode != nil; travNode = travNode.nextNode {
		count++
	}
	return count
}

// printLinkedlist print list by iterating over all nodes
func (list *LinkedList) printLinkedlist() {

	fmt.Printf(" [START]")
	for node := list.headNode; node != nil; node = node.nextNode {
		fmt.Printf(" <-> [%v]", node.property)
	}
	fmt.Printf(" <-> [END]\n")
}

// printLLByNode print linked list when node is given
func (node *Node) printLLByNode() {

	fmt.Printf(" [START]")
	for node := node; node != nil; node = node.nextNode {
		fmt.Printf(" <-> [%v]", node.property)
	}
	fmt.Printf(" <-> [END]\n")
}

// Link: https://leetcode.com/problems/reverse-linked-list/
func (list *LinkedList) reverseLinkedList() *Node {

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
	return prev
}
