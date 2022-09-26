package main

import "fmt"

func main() {
	var tree *BinarySearchTree = &BinarySearchTree{}
	tree.InsertElement(8, 8)
	tree.InsertElement(3, 3)
	tree.InsertElement(10, 10)
	tree.InsertElement(1, 1)
	tree.InsertElement(6, 6)
	tree.InOrderTraverseTree()
	tree.PreOrderTraverseTree()
	tree.PostOrderTraverseTree()
	fmt.Println()
	tree.String()
	fmt.Println()
}
