package main

import (
	"fmt"
	"sync"
)

type Node struct {
	key       int
	value     int
	leftNode  *Node
	rightNode *Node
}

type BinarySearchTree struct {
	rootNode *Node
	lock     sync.RWMutex
}

// Insert inserts the Item t in the tree
func (bst *BinarySearchTree) InsertElement(key int, value int) {
	bst.lock.Lock()
	defer bst.lock.Unlock()

	var newNode = &Node{key, value, nil, nil}
	if bst.rootNode == nil {
		bst.rootNode = newNode
	} else {
		insertNode(bst.rootNode, newNode)
	}
}

// internal function to find the correct place for a node in a tree
func insertNode(rootNode *Node, newNode *Node) {
	if newNode.key < rootNode.key {
		if rootNode.leftNode == nil {
			rootNode.leftNode = newNode
		} else {
			insertNode(rootNode.leftNode, newNode)
		}
	} else {
		if rootNode.rightNode == nil {
			rootNode.rightNode = newNode
		} else {
			insertNode(rootNode.rightNode, newNode)
		}
	}
}

// InOrderTraverseTree method
func (tree *BinarySearchTree) InOrderTraverseTree() {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	fmt.Println("In-order traversal :  ")
	inOrderTraverseTree(tree.rootNode)
}

// inOrderTraverseTree method
func inOrderTraverseTree(Node *Node) {
	if Node != nil {
		inOrderTraverseTree(Node.leftNode)
		fmt.Printf("[%v]\t", Node.value)
		inOrderTraverseTree(Node.rightNode)
	}
}

// PreOrderTraverse method
func (tree *BinarySearchTree) PreOrderTraverseTree() {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	fmt.Println()
	fmt.Println("Pre-order traversal :  ")
	preOrderTraverseTree(tree.rootNode)
}

// preOrderTraverseTree method
func preOrderTraverseTree(Node *Node) {
	if Node != nil {
		fmt.Printf("[%v]\t", Node.value)
		preOrderTraverseTree(Node.leftNode)
		preOrderTraverseTree(Node.rightNode)
	}
}

// PostOrderTraverseTree method
func (tree *BinarySearchTree) PostOrderTraverseTree() {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	fmt.Println()
	fmt.Println("Post-order traversal :  ")
	postOrderTraverseTree(tree.rootNode)
}

// postOrderTraverseTree method
func postOrderTraverseTree(Node *Node) {
	if Node != nil {
		postOrderTraverseTree(Node.leftNode)
		postOrderTraverseTree(Node.rightNode)
		fmt.Printf("[%v]\t", Node.value)
	}
}

// MinNode method
func (tree *BinarySearchTree) MinNode() *int {
	tree.lock.RLock()
	defer tree.lock.RUnlock()

	var Node *Node
	Node = tree.rootNode
	if Node == nil {
		return (*int)(nil)
	}
	for {
		if Node.leftNode == nil {
			return &Node.value
		}
		Node = Node.leftNode
	}
}

// MaxNode method
func (tree *BinarySearchTree) MaxNode() *int {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	var Node *Node
	Node = tree.rootNode
	if Node == nil {
		return (*int)(nil)
	}
	for {
		if Node.rightNode == nil {
			return &Node.value
		}
		Node = Node.rightNode
	}
}

// SearchNode method
func (tree *BinarySearchTree) SearchNode(key int) bool {
	tree.lock.RLock()
	defer tree.lock.RUnlock()
	return searchNode(tree.rootNode, key)
}

// searchNode method
func searchNode(parentNode *Node, searchKey int) bool {
	if parentNode == nil {
		return false
	}
	if searchKey < parentNode.key {
		return searchNode(parentNode.leftNode, searchKey)
	}
	if searchKey > parentNode.key {
		return searchNode(parentNode.rightNode, searchKey)
	}
	return true
}

// RemoveNode method
func (tree *BinarySearchTree) RemoveNode(deleteKey int) {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	removeNode(tree.rootNode, deleteKey)
}

// removeNode method
func removeNode(parentNode *Node, deleteKey int) *Node {
	if parentNode == nil {
		return nil
	}

	if deleteKey < parentNode.key {
		parentNode.leftNode = removeNode(parentNode.leftNode, deleteKey)
		return parentNode
	}
	if deleteKey > parentNode.key {
		parentNode.rightNode = removeNode(parentNode.rightNode, deleteKey)
		return parentNode
	}

	if parentNode.leftNode == nil && parentNode.rightNode == nil {
		parentNode = nil
		return nil
	}

	if parentNode.leftNode == nil {
		parentNode = parentNode.rightNode
		return parentNode
	}
	if parentNode.rightNode == nil {
		Node := parentNode.leftNode
		return Node
	}
	// var leftmostrightNode *Node
	leftmostrightNode := parentNode.rightNode
	for {

		if leftmostrightNode != nil && leftmostrightNode.leftNode != nil {
			leftmostrightNode = leftmostrightNode.leftNode
		} else {
			break
		}
	}
	parentNode.key, parentNode.value = leftmostrightNode.key, leftmostrightNode.value
	parentNode.rightNode = removeNode(parentNode.rightNode, parentNode.key)
	return parentNode
}

// String method
func (tree *BinarySearchTree) String() {
	tree.lock.Lock()
	defer tree.lock.Unlock()
	fmt.Println("************************************************")
	stringify(tree.rootNode, 0)
	fmt.Println("************************************************")
}

// stringify method
func stringify(Node *Node, level int) {
	if Node != nil {
		format := ""
		for i := 0; i < level; i++ {
			format += "     "
		}
		format += "**> "
		level++
		stringify(Node.leftNode, level)
		fmt.Printf(format+"%d\n", Node.key)
		stringify(Node.rightNode, level)
	}
}

// print method
func print(tree *BinarySearchTree) {
	if tree != nil {

		fmt.Println(" Value", tree.rootNode.value)
		fmt.Printf("Root Tree Node")
		printNode(tree.rootNode)
	} else {
		fmt.Printf("Nil\n")
	}
}

// printNode method
func printNode(Node *Node) {
	if Node != nil {
		fmt.Println(" Value", Node.value)
		fmt.Printf("Node Left")
		printNode(Node.leftNode)
		fmt.Printf("Node Right")
		printNode(Node.rightNode)
	} else {
		fmt.Printf("Nil\n")
	}

}
