package main

import "fmt"

func main() {
	var tree *BinarySearchTree = &BinarySearchTree{}
	tree.InsertElement(8, 8)
	tree.InsertElement(3, 3)
	tree.InsertElement(10, 10)
	tree.InsertElement(1, 1)
	tree.InsertElement(6, 6)
	tree.InsertElement(5, 5)
	tree.InOrderTraverseTree()
	tree.PreOrderTraverseTree()
	tree.PostOrderTraverseTree()
	fmt.Println()
	tree.String()
	fmt.Println()
	maxDepth := maxDepth(tree.rootNode)
	fmt.Println("max depth : ", maxDepth)
	isBalanced := isBalanced(tree.rootNode)
	fmt.Println("is balanced : ", isBalanced)
	diameterOfTree := diameterOfTree(tree.rootNode, 0)
	fmt.Println("diameterOfTree : ", diameterOfTree)

}
