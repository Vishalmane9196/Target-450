package main

import "fmt"

func main() {

	//binary tree implementationand demo
	var tree *BinaryTree = &BinaryTree{}
	tree.InsertElement(8, 8)
	tree.InsertElement(3, 3)
	tree.InsertElement(10, 10)
	tree.InsertElement(1, 1)
	tree.InsertElement(6, 6)
	tree.InsertElement(5, 5)
	tree.InsertElement(12, 12)
	tree.InOrderTraverseTree()
	tree.PreOrderTraverseTree()
	tree.PostOrderTraverseTree()
	fmt.Println()
	tree.PrintBinaryTree()
	fmt.Println()
	// tree.RemoveNode(10)
	// tree.PrintBinaryTree()
	// fmt.Println(" postorder : ", postorderTraversalIterative(tree.rootNode))
	// fmt.Println(" inorder : ", inorderTraversalIterative(tree.rootNode))
	// fmt.Println(" postorder : ", postOrderTraversalUsing1Stack(tree.rootNode))
	// fmt.Println(" preorder : ", preorderTraversalIterative(tree.rootNode))

	// pre, in, post := allTraversal(tree.rootNode)
	// fmt.Printf(" preOrder : %v \n inOrder : %v \n postOrder : %v \n ", pre, in, post)

	//problems
	// levelOrder := levelOrder(tree.rootNode)
	// fmt.Println("levelOrder : ", levelOrder)

	// fmt.Println(" Min node of tree : ", *tree.MinNode())
	// fmt.Println(" Max node of tree : ", *tree.MaxNode())
	// fmt.Println(" Search node of tree : ", tree.SearchNode(1))

	// maxDepth := maxDepth(tree.rootNode)
	// fmt.Println("max depth : ", maxDepth)

	// isBalanced := isBalanced(tree.rootNode)
	// fmt.Println("is balanced : ", isBalanced)

	// diameterOfTree := diameterOfBinaryTree(tree.rootNode)
	// fmt.Println("diameterOfBinaryTree : ", diameterOfTree)

	// maxPathSum := maxPathSum(tree.rootNode)
	// fmt.Println("maxPathSum : ", maxPathSum)

	// res := preorderTraversalIterative(tree.rootNode)
	// fmt.Println("preorderTraversalIterative : ", res)

	// res := preorderTraversalIterative(tree.rootNode)
	// fmt.Println("preorderTraversalIterative : ", res)

	// res := rightSideView(tree.rootNode)
	// fmt.Println("rightSideView : ", res)

	// res = rightSideViewRec(tree.rootNode)
	// fmt.Println("rightSideViewRec : ", res)

	// res := bottomView(tree.rootNode)
	// fmt.Println("bottomView : ", res)

	// res := topView(tree.rootNode)
	// fmt.Println("topView : ", res)

	// res := verticalTraversal(tree.rootNode)
	// fmt.Println("verticalTraversal : ", res)

	// res := printBoundary(tree.rootNode)
	// fmt.Println("printBoundary : ", res)

	// res := zigzagLevelOrder(tree.rootNode)
	// fmt.Println("zigzagLevelOrder : ", res)

	// var arr = make([]int, 0)
	// getPath(tree.rootNode, &arr, 5)
	// fmt.Println("getPath : ", arr)

	// res := widthOfBinaryTree(tree.rootNode)
	// fmt.Println("widthOfBinaryTree : ", res)

	// res := addOneRow(tree.rootNode, 1, 3)
	// fmt.Println("addOneRow : ", res)
	// tree.PrintBinaryTree()

}
