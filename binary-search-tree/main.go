package main

import "fmt"

func main() {

	tree := &BinarySearchTree{}

	elements := []int{5, 6, 3, 4, 2, 1}

	for _, v := range elements {
		tree.insert(v)
	}
	tree.PrintBinarySearchTree()
	res := tree.searchBST(3)
	fmt.Println("Search in BST : ", res)

	fmt.Printf("\nMaximum element is %d\n", FindMax(tree.root))
	fmt.Printf("\nMinimum element is %d\n", FindMin(tree.root))
	fmt.Printf("\nkth smallest element is %d\n", kthSmallest(tree.root, 3))
	fmt.Printf("\nkth largest element is %d\n", kthLargest(tree.root, 3))
	fmt.Printf("\nis valid bst:  %v\n", isValidBST(tree.root))
	fmt.Printf("\nceil of %v in bst is : %v\n", 4, ceil(tree.root, 4))
	fmt.Printf("\nfloor of %v in bst is : %v\n", 3, floor(tree.root, 4))
	fmt.Printf("\nlowestCommonAncestor of %v  and %v in bst is : %v\n", 3, 6, lowestCommonAncestor(tree.root, &Node{Val: 3}, &Node{Val: 6}))
	fmt.Printf("\nsuccessor of %v in bst is : %v  \n", 3, inOrderSuccessor(tree.root, &Node{Val: 3}))
	fmt.Printf("\npredecessor of %v in bst is : %v  \n", 3, inOrderpredecessor(tree.root, &Node{Val: 3}))
}
