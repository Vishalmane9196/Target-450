package main

import (
	"fmt"
)

// Node class
type Node struct {
	property     int
	nextNode     *Node
	prevNode *Node
}

// LinkedList class
type LinkedList struct {
	headNode *Node
	// tailNode *Node
}

//insertAtStart list.method will be used for isnerting node at starting position
func (list *LinkedList) insertAtStart(property int) {
	var newNode = &Node{property, nil, nil}

	if list.headNode != nil {
		newNode.nextNode = list.headNode
		list.headNode.prevNode = newNode
	}
	list.headNode = newNode
}

//inserAtEnd method adds the node with property to the end
func (list *LinkedList) inserAtEnd(property int) {
	var newNode = &Node{property, nil, nil}

	var lastNode = &Node{}
	lastNode = list.getLastNode()

	if lastNode != nil {
		lastNode.nextNode = newNode
		newNode.prevNode = lastNode
	}
}

//insert_after method  adds a node with nodeProperty after node with property
func (list *LinkedList) insertAfter(existingData int, newData int) {
	var newNode = &Node{newData, nil, nil}
	var dataNode *Node

	dataNode = list.getNode(existingData)
	if dataNode != nil {
		newNode.nextNode = dataNode.nextNode
		newNode.prevNode = dataNode
		dataNode.nextNode.prevNode = newNode
		dataNode.nextNode = newNode
	}
}

func (list *LinkedList) delAtStart(){

	if list.headNode == nil {
       fmt.Println("The List is empty")
	}else{
        if list.headNode.nextNode == nil {
			list.headNode = nil
		}else{
			list.headNode = list.headNode.nextNode
			list.headNode.prevNode = nil
		}	
	}
}

func (list *LinkedList) delAtEnd(){
  var lastNode = &Node{}
  lastNode = list.getLastNode()

  if lastNode != nil {
	lastNode.prevNode.nextNode = nil
	lastNode.nextNode = nil
  }
}

//remove_data method remove the data from linked list
func (list *LinkedList) delNode(data int){
	 var deleteNode = &Node{}

	 deleteNode = list.getNode(data)
	 if deleteNode != nil {
		deleteNode.prevNode.nextNode = deleteNode.nextNode
		deleteNode.nextNode.prevNode = deleteNode.prevNode
	 }
}

//getFirstNode method returns the first Node of linked list
func (list *LinkedList) getFirstNode() *Node {
  return list.headNode
}

//getLastNode method returns the last Node
func (list *LinkedList) getLastNode() *Node {
	var travNode *Node
	var lastNode *Node

	for travNode = list.headNode; travNode != nil; travNode = travNode.nextNode {
		if travNode.nextNode == nil {
			lastNode = travNode
		}
	}
	return lastNode
}

//getNode method returns Node given parameter property
func (list *LinkedList) getNode(data int) *Node {
	var tempNode *Node
	var travNode *Node

	for travNode = list.headNode; travNode != nil; travNode = travNode.nextNode {
		if travNode.property == data {
			tempNode = travNode
			break
		}
	}
	return tempNode
}

//getprevNode method returns 2nd last elment of given property
func (list *LinkedList) getPreviousNode(data int) *Node {
	var tempNode *Node
	var travNode *Node
	for travNode = list.headNode; travNode.nextNode!= nil; travNode = travNode.nextNode {
		if tempNode.nextNode.property == data {
			tempNode = travNode
			break
		}
	}
	return tempNode
}

//IterateList method of LinkedList
func (list *LinkedList) printLinkedList() {
	var travNode *Node

	fmt.Printf(" [START]")
	travNode = list.headNode
	for travNode != nil {
		fmt.Printf(" <-> [%v]", travNode.property)
		travNode = travNode.nextNode
	}
	fmt.Printf(" <-> [END]\n")
}

//isListEmpty method return the linked list is empty or not
func(list *LinkedList) isListEmpty() bool {
	return list.headNode == nil
}

//NodeBetweenValues method of LinkedList
func (list *LinkedList) NodeBetweenValues(firstNodeData, secondNodeData int) *Node {
	var travNode *Node
	var tempNode *Node

	for travNode = list.headNode; travNode != nil; travNode = travNode.nextNode {
		if travNode.prevNode != nil && travNode.nextNode != nil {
			if travNode.prevNode.property == firstNodeData && travNode.nextNode.property == secondNodeData {
				tempNode = travNode
				break
			}
		}
	}
	return tempNode
}